package simplelog

import (
	"github.com/tanzy2018/simplelog/encode"
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
	l.wc = os.Stdout
	l.autoReName = false
	atomic.StoreInt32(&l.alock, 0)
	l.lo = new(sync.Mutex)
	l.backendSync()
	return l
}

// Debug ...
func Debug(msg string, md ...encode.Meta) {
	defaultLog.Debug(msg, md...)
}

// Info ...
func Info(msg string, md ...encode.Meta) {
	defaultLog.Info(msg, md...)
}

// Warn ...
func Warn(msg string, md ...encode.Meta) {
	defaultLog.Warn(msg, md...)
}

// Error ...
func Error(msg string, md ...encode.Meta) {
	defaultLog.Error(msg, md...)
}

// Panic ...
func Panic(msg string, md ...encode.Meta) {
	defaultLog.Panic(msg, md...)
}

// Fatal ...
func Fatal(msg string, md ...encode.Meta) {
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

// DeFault ...
func DeFault() *Log {
	return defaultLog
}
