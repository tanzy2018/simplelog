package simplelog

import (
	"io"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/tanzy2018/simplelog/encode"
	"github.com/tanzy2018/simplelog/internal"
)

// Logger ...
type Logger interface {
	// Debug ...
	Debug(msg string, md ...encode.Meta)
	// Info ...
	Info(msg string, md ...encode.Meta)
	// Warn ...
	Warn(msg string, md ...encode.Meta)
	// Error ...
	Error(msg string, md ...encode.Meta)
	// Panic ...
	Panic(msg string, md ...encode.Meta)
	// Fatal ...
	Fatal(msg string, md ...encode.Meta)
}

// Log ...
type Log struct {
	op          *options
	wc          io.WriteCloser
	curFileSize int64
	autoReName  bool
	syncBufs    []*syncBuffer
	recordBufs  []*oneRecordBuffer
	alock       int32
	tk          *time.Ticker
	lo          *sync.Mutex
	nopClose    bool
}

// New ...
func New(ops ...Option) *Log {
	l := &Log{
		op: defaultOption,
	}
	for _, f := range ops {
		f(l.op)
	}
	l.syncBufs = newSyncBuffers(l, l.op.maxSyncBufSize, l.op.syncBufsLen)
	l.recordBufs = newOneRecordBuffers(l, l.op.maxRecordSize, l.op.recordBufsLen)
	l.lo = new(sync.Mutex)
	atomic.StoreInt32(&l.alock, 0)
	l.wc = os.Stdout
	l.nopClose = true
	l.backendSync()
	return l
}

// Debug ...
func (l *Log) Debug(msg string, md ...encode.Meta) {
	// 少一次函数调用
	if l.op.level > int32(DEBUG) {
		return
	}
	l.write(DEBUG, msg, md...)
}

// Info ...
func (l *Log) Info(msg string, md ...encode.Meta) {
	// reduce another function call
	if l.op.level > int32(INFO) {
		return
	}
	l.write(INFO, msg, md...)
}

// Warn ...
func (l *Log) Warn(msg string, md ...encode.Meta) {
	// reduce another function call
	if l.op.level > int32(WARN) {
		return
	}
	l.write(WARN, msg, md...)
}

// Error ...
func (l *Log) Error(msg string, md ...encode.Meta) {
	// reduce another function call
	if l.op.level > int32(ERROR) {
		return
	}
	l.write(ERROR, msg, md...)
}

// Panic ...
func (l *Log) Panic(msg string, md ...encode.Meta) {
	// reduce another function call
	if l.op.level > int32(PANIC) {
		return
	}
	l.write(PANIC, msg, md...)
}

// Fatal ...
func (l *Log) Fatal(msg string, md ...encode.Meta) {
	// reduce another function call
	if l.op.level > int32(FATAL) {
		return
	}
	l.write(FATAL, msg, md...)
	l.lock()
	l.syncAll()
	l.close()
	l.unlock()
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
	l.lock()
	l.unlock()
	l.syncAll()
	l.close()
}

// WithWriterCloser ...
func (l *Log) WithWriterCloser(wc io.WriteCloser, needAutoRename, nopClose bool) *Log {
	l.lock()
	defer l.unlock()
	l.syncAll()
	l.close()
	l.wc = wc
	l.autoReName = needAutoRename
	l.nopClose = true
	return l
}

// WithFileWriter ...
func (l *Log) WithFileWriter(root, topic, fname string) *Log {
	l.lock()
	defer l.unlock()
	l.syncAll()
	l.errHandle(l.close())
	l.updateFileOption(root, topic, fname)
	l.errHandle(l.makedir())
	l.errHandle(l.newWriterCloserFromFile())
	return l
}

func (l *Log) updateFileOption(root, topic, fname string) {
	l.op.updateFileOption(root, topic, fname)
}

func (l *Log) errHandle(errs ...error) {
	if errHandle := l.op.errHandler; errHandle != nil {
		for _, err := range errs {
			if err != nil {
				errHandle(err)
			}
		}
	}
}

func (l *Log) sync(idx int) {
	b := l.syncBufs[idx].flushAsBytes()
	if len(b) == 0 {
		return
	}
	l.curFileSize += int64(len(b))
	l.wc.Write(b)
	l.orChangeFileWriter()

}

func (l *Log) syncAll() {
	for i := 0; i < l.op.syncBufsLen; i++ {
		b := l.syncBufs[i].flushAsBytes()
		if len(b) == 0 {
			continue
		}
		l.curFileSize += int64(len(b))
		l.wc.Write(b)
		l.orChangeFileWriter()
	}
}

func (l *Log) orChangeFileWriter() {
	if !l.autoReName {
		return
	}
	if l.curFileSize < l.op.maxFileSize {
		return
	}

	l.errHandle(
		// close
		l.close(),
		// rename
		os.Rename(l.op.fullPath(), l.op.rename()),
		// open
		l.newWriterCloserFromFile(),
	)

}

func (l *Log) write(level LevelType, msg string, md ...encode.Meta) {
	idx := internal.RandInt(l.op.recordBufsLen)
	l.recordBufs[idx].write(level, msg, md)
	sync := l.syncBufs[idx].write(l.recordBufs[idx].flushAsBytes())
	if l.op.writeDirect || sync {
		l.lock()
		defer l.unlock()
		l.sync(idx)
		return
	}

}

func (l *Log) makedir() error {
	dir := l.op.dir()
	if len(dir) == 0 {
		return nil
	}
	fi, err := os.Stat(dir)
	if err == nil && fi.IsDir() {
		return nil
	}
	return os.MkdirAll(dir, 0755)
}

func (l *Log) newWriterCloserFromFile() error {
	f, err := os.OpenFile(l.op.fullPath(), os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	l.wc = f
	l.autoReName = false
	l.op.cTime = time.Now().Unix()
	if !l.op.isAutoRenameFile() {
		l.nopClose = true
		return nil
	}
	l.autoReName = true
	fi, err := f.Stat()
	if err == nil {
		l.curFileSize = fi.Size()
		l.op.cTime = fi.ModTime().Unix()
	}
	return nil
}

func (l *Log) close() error {
	if l.nopClose {
		return nil
	}
	return l.wc.Close()
}

func (l *Log) backendSync() {
	l.tk = time.NewTicker(l.op.syncInterval)
	go func() {
		for range l.tk.C {
			if l.op.writeDirect {
				continue
			}
			l.lock()
			l.syncAll()
			l.unlock()
		}
	}()
}

func (l *Log) lock() {
	l.lo.Lock()
}

func (l *Log) unlock() {
	l.lo.Unlock()
}

func wrapPath(paths ...string) string {
	path := strings.Join(paths, "/")
	// `//`is empty fname.
	if len(path) == 2 {
		return ""
	}
	return path
}
