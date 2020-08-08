package simplelog

import (
	"path"
	"strings"
	"time"

	"github.com/tanzy2018/simplelog/utils"
)

// LevelType ...
type LevelType int32

const (
	_ LevelType = iota
	// DEBUG ...
	DEBUG
	// INFO ...
	INFO
	// WARN ...
	WARN
	// ERROR ...
	ERROR
	// PANIC ...
	PANIC
	// FATAL ...
	FATAL
)

var defaultOption = _defaultOPtion()

func _defaultOPtion() *options {
	return &options{
		level:          int32(DEBUG),
		maxFileSize:    1024 * 1024 * 1024,
		maxSyncBufSize: 1024 * 1024,
		maxRecordSize:  1024 * 1024,
		writeDirect:    false,
		syncBufsLen:    10,
		recordBufsLen:  10,
		syncInterval:   time.Millisecond * 100,
		hook:           new(hook),
	}
}

func (level LevelType) String() string {
	ltype := "debug"
	switch level {
	case DEBUG:
		ltype = "debug"
	case INFO:
		ltype = "info"
	case WARN:
		ltype = "warn"
	case ERROR:
		ltype = "error"
	case PANIC:
		ltype = "panic"
	case FATAL:
		ltype = "fatal"
	}
	return ltype
}

// Option ...
type Option func(op *options)
type options struct {
	level          int32
	root           string
	topic          string
	fname          string
	maxSyncBufSize int
	maxFileSize    int64
	maxRecordSize  int
	syncBufsLen    int
	recordBufsLen  int
	writeDirect    bool
	fileCreateTime int64
	syncInterval   time.Duration
	hook           IHook
}

func (op *options) fullPath() string {
	return wrapPath(
		strings.TrimRight(strings.ReplaceAll(op.root, "\\", "/"), "/"),
		strings.Trim(strings.ReplaceAll(op.topic, "\\", "/"), "/"),
		strings.Trim(strings.ReplaceAll(op.fname, "\\", "/"), "/"))
}

func (op *options) basePath() string {
	return path.Base(op.fname)
}

func (op *options) rename() string {
	full := op.fullPath()
	base := path.Base(op.fname)
	ext := path.Ext(base)
	subfix := genRenameSubfix(op.fileCreateTime, time.Now().Unix())
	newName := make([]byte, 0, len(full))
	newName = append(newName, []byte(full)[:len(full)-len(base)]...)
	newName = append(newName, []byte(base)[:len(base)-len(ext)]...)
	newName = append(newName, []byte(subfix)...)
	if len(ext) > 0 {
		newName = append(newName, []byte(ext)...)
	}
	return utils.ToString(newName)
}

func (op *options) dir() string {
	return wrapPath(
		strings.TrimRight(strings.ReplaceAll(op.root, "\\", "/"), "/"),
		strings.Trim(strings.ReplaceAll(op.topic, "\\", "/"), "/"))
}

// WithFileWriter ...
func WithFileWriter(root, topic, fname string) Option {
	return func(op *options) {
		op.fname = fname
		op.root = root
		op.topic = topic
	}
}

// WithMaxFileSize ...
func WithMaxFileSize(size int64) Option {
	return func(op *options) {
		op.maxFileSize = size
	}
}

// WithMaxSyncSize ...
func WithMaxSyncSize(size int) Option {
	return func(op *options) {
		op.maxSyncBufSize = size
	}
}

// WithMaxRecordSize ...
func WithMaxRecordSize(size int) Option {
	return func(op *options) {
		op.maxRecordSize = size
	}
}

// WithLevel ...
func WithLevel(level int32) Option {
	return func(op *options) {
		op.level = level
	}
}

// WithHook ...
func WithHook(hook IHook) Option {
	return func(op *options) {
		op.hook = hook
	}
}

// WithSyncBufsLen ...
func WithSyncBufsLen(syncBufsLen int) Option {
	return func(op *options) {
		op.syncBufsLen = syncBufsLen
	}
}

// WithRecordBufsLen ...
func WithRecordBufsLen(recordBufsLen int) Option {
	return func(op *options) {
		op.recordBufsLen = recordBufsLen
	}
}

// WithWriteDirect ...
func WithWriteDirect(writeDirect bool) Option {
	return func(op *options) {
		op.writeDirect = writeDirect
	}
}

// WithSyncInterval ...
func WithSyncInterval(dur time.Duration) Option {
	return func(op *options) {
		if dur > 0 {
			op.syncInterval = dur
		}
	}
}
