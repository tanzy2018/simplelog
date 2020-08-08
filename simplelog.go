package simplelog

import (
	"github.com/tanzy2018/simplelog/meta"
	"os"
	"sync"
	"sync/atomic"
)

var defaultLog = _defaultLog()

func _defaultLog() *Log {
	l := &Log{
		op: defaultOption,
	}
	l.syncBufs = newSyncBuffers(l, l.op.maxSyncBufSize, l.op.syncBufsLen)
	l.recordBufs = newOneRecordBuffers(l, l.op.maxRecordSize, l.op.recordBufsLen)
	l.file = os.Stdout
	l.isStdio = true
	atomic.StoreInt32(&l.alock, 0)
	l.lo = new(sync.Mutex)
	l.backendSync()
	return l
}

// Debug ...
func Debug(msg string, md ...meta.Meta) {
	defaultLog.Debug(msg, md...)
}

// Info ...
func Info(msg string, md ...meta.Meta) {
	defaultLog.Info(msg, md...)
}

// Warn ...
func Warn(msg string, md ...meta.Meta) {
	defaultLog.Warn(msg, md...)
}

// Error ...
func Error(msg string, md ...meta.Meta) {
	defaultLog.Error(msg, md...)
}

// Panic ...
func Panic(msg string, md ...meta.Meta) {
	defaultLog.Panic(msg, md...)
}

// Fatal ...
func Fatal(msg string, md ...meta.Meta) {
	defaultLog.Fatal(msg, md...)
}

// Hook ...
func Hook(hfs ...HookFunc) {
	defaultLog.Hook(hfs...)
}

// Sync ...
func Sync() {
	defaultLog.Sync()
}

// InitFileWriter ... Warn
func InitFileWriter(root, topic, fname string) error {
	defaultLog.op.root = root
	defaultLog.op.topic = topic
	defaultLog.op.fname = fname
	if err := defaultLog.makedir(); err != nil {
		return err
	}
	if err := defaultLog.closeFile(); err != nil {
		return err
	}
	if err := defaultLog.openFile(); err != nil {
		return err
	}
	return nil
}

// SetDirectWrite ...
func SetDirectWrite(directWrite bool) {
	defaultLog.lock()
	defer defaultLog.unlock()
	defaultLog.op.writeDirect = directWrite
}
