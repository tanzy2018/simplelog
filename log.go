package simplelog

import (
	"github.com/tanzy2018/simplelog/meta"
	"github.com/tanzy2018/simplelog/utils"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// Logger ...
type Logger interface {
	// Debug ...
	Debug(msg string, md ...meta.Meta)
	// Info ...
	Info(msg string, md ...meta.Meta)
	// Warn ...
	Warn(msg string, md ...meta.Meta)
	// Error ...
	Error(msg string, md ...meta.Meta)
	// Panic ...
	Panic(msg string, md ...meta.Meta)
	// Fatal ...
	Fatal(msg string, md ...meta.Meta)
}

// Log ...
type Log struct {
	op          *options
	file        *os.File
	curFileSize int64
	isStdio     bool // stdout,stderr
	syncBufs    []*syncBuffer
	recordBufs  []*oneRecordBuffer
	alock       int32
	tk          *time.Ticker
	lo          *sync.Mutex
}

// New ...
func New(ops ...Option) (*Log, error) {
	l := &Log{
		op: defaultOption,
	}
	for _, f := range ops {
		f(l.op)
	}
	if err := l.makedir(); err != nil {
		return nil, err
	}
	if err := l.openFile(); err != nil {
		return nil, err
	}
	l.syncBufs = newSyncBuffers(l, l.op.maxSyncBufSize, l.op.syncBufsLen)
	l.recordBufs = newOneRecordBuffers(l, l.op.maxRecordSize, l.op.recordBufsLen)
	l.lo = new(sync.Mutex)
	atomic.StoreInt32(&l.alock, 0)
	l.backendSync()
	return l, nil
}

// Debug ...
func (l *Log) Debug(msg string, md ...meta.Meta) {
	l.write(DEBUG, msg, md...)
}

// Info ...
func (l *Log) Info(msg string, md ...meta.Meta) {
	l.write(INFO, msg, md...)
}

// Warn ...
func (l *Log) Warn(msg string, md ...meta.Meta) {
	l.write(WARN, msg, md...)
}

// Error ...
func (l *Log) Error(msg string, md ...meta.Meta) {
	l.write(ERROR, msg, md...)
}

// Panic ...
func (l *Log) Panic(msg string, md ...meta.Meta) {
	l.write(PANIC, msg, md...)
}

// Fatal ...
func (l *Log) Fatal(msg string, md ...meta.Meta) {
	l.write(FATAL, msg, md...)
	l.syncAll()
	l.closeFile()
	os.Exit(-1)
}

// Hook ...
func (l *Log) Hook(hfs ...HookFunc) {
	if l.op.hook == nil {
		l.op.hook = new(hook)
	}
	l.op.hook.Add(hfs...)
}

// Sync ...
func (l *Log) Sync() {
	l.syncAll()
	l.closeFile()
}

func (l *Log) sync(idx int) {
	b := l.syncBufs[idx].flushAsBytes()
	if len(b) == 0 {
		return
	}
	l.lock()
	defer l.unlock()
	l.curFileSize += int64(len(b))

	l.file.WriteString(utils.ToString(b))

	l.orChangeFileWriter()

}

func (l *Log) syncAll() {
	l.lock()
	defer l.unlock()
	for i := 0; i < l.op.syncBufsLen; i++ {
		b := l.syncBufs[i].flushAsBytes()
		if len(b) == 0 {
			continue
		}
		l.curFileSize += int64(len(b))
		l.file.WriteString(utils.ToString(b))
		// l.file.WriteString("11")
		l.orChangeFileWriter()
	}
}

func (l *Log) orChangeFileWriter() {
	if l.isStdio {
		return
	}
	if l.curFileSize < l.op.maxFileSize {
		return
	}

	// close
	l.file.Close()
	// rename
	os.Rename(l.op.fullPath(), l.op.rename())
	// open
	l.openFile()
}

func (l *Log) write(level LevelType, msg string, md ...meta.Meta) {
	if l.op.level > int32(level) {
		return
	}
	idx := utils.RandInt(l.op.recordBufsLen)
	l.recordBufs[idx].write(level, msg, md)
	sync := l.syncBufs[idx].write(l.recordBufs[idx].flushAsBytes())
	if l.op.writeDirect || sync {
		l.sync(idx)
		return
	}

}

func (l *Log) makedir() error {
	dir := l.op.dir()
	fi, err := os.Stat(dir)
	if err == nil && fi.IsDir() {
		return nil
	}
	return os.MkdirAll(dir, 0755)
}

func (l *Log) openFile() error {
	f, err := os.OpenFile(l.op.fullPath(), os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	l.file = f
	l.isStdio = false
	l.op.fileCreateTime = time.Now().Unix()
	fi, err := f.Stat()
	if err != nil {
		return err
	}

	l.curFileSize = fi.Size()

	return nil
}

func (l *Log) closeFile() error {
	l.lock()
	defer l.unlock()
	return l.file.Close()
}

func (l *Log) backendSync() {
	l.tk = time.NewTicker(l.op.syncInterval)
	go func() {
		for range l.tk.C {
			if l.op.writeDirect {
				continue
			}
			l.syncAll()
		}
	}()
}

func (l *Log) lock() {
	// for {
	// 	if atomic.CompareAndSwapInt32(&l.alock, 0, 1) {
	// 		return
	// 	}
	// }
	// l.lo.Lock()
	l.lo.Lock()
}

func (l *Log) unlock() {
	// for {
	// 	if atomic.CompareAndSwapInt32(&l.alock, 1, 0) {
	// 		return
	// 	}
	// }
	// l.lo.UnLock()
	l.lo.Unlock()
}

func wrapPath(paths ...string) string {
	return strings.Join(paths, "/")
}
