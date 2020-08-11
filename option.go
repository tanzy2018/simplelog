package simplelog

import (
	"fmt"
	"github.com/tanzy2018/simplelog/internal"
	"os"
	"path"
	"strings"
	"time"
)

var defaultOption = _defaultOPtion()

var notAutoRenameFileNames = [3]string{"/dev/stdout", "/dev/stdin", "/dev/stderr"}

func _defaultOPtion() *options {
	return &options{
		level:          int32(DEBUG),
		maxFileSize:    1024 * 1024 * 1024,
		maxSyncBufSize: 1024 * 1024,
		maxRecordSize:  1024 * 10,
		syncDirect:     true,
		syncInterval:   time.Second * 1,
		hook:           new(hook),
		errHandler: func(err error) {
			fmt.Fprintf(os.Stderr, "log err:%v\n", err)
		},
	}
}

// ErrorHandler ... handle the error from the log.
type ErrorHandler func(error)

// Option ...
type Option func(op *options)
type options struct {
	level          int32
	root           string
	topic          string
	fname          string
	errHandler     ErrorHandler
	maxSyncBufSize int
	maxFileSize    int64
	maxRecordSize  int
	syncDirect     bool
	cTime          int64
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
	subfix := genRenameSubfix(op.cTime, time.Now().Unix())

	newName := make([]byte, 0, 1+len(full)+len(subfix))
	newName = append(newName, full[:len(full)-len(base)]...)
	newName = append(newName, base[:len(base)-len(ext)]...)
	newName = append(newName, subfix...)
	if len(ext) > 0 {
		newName = append(newName, ext...)
	}
	return internal.ToString(newName)
}

func (op *options) dir() string {
	return wrapPath(
		strings.TrimRight(strings.ReplaceAll(op.root, "\\", "/"), "/"),
		strings.Trim(strings.ReplaceAll(op.topic, "\\", "/"), "/"))
}

func (op *options) isAutoRenameFile() bool {
	fname := op.fullPath()
	for _, fn := range notAutoRenameFileNames {
		if fname == fn {
			return false
		}
	}
	return true
}

func (op *options) updateFileOption(root, topic, fname string) {
	op.root = root
	op.topic = topic
	op.fname = fname
}

func (op *options) isValidFileName() bool {
	if len(op.fullPath()) == 0 {
		return false
	}
	return true
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
func WithLevel(level LevelType) Option {
	return func(op *options) {
		if level.isValid() {
			op.level = int32(level)
		}
	}
}

// WithHook ...
func WithHook(hook IHook) Option {
	return func(op *options) {
		op.hook = hook
	}
}

// WithSyncDirect ...
func WithSyncDirect(syncDirect bool) Option {
	return func(op *options) {
		op.syncDirect = syncDirect
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

// WithErrorHandler ...
func WithErrorHandler(f ErrorHandler) Option {
	return func(op *options) {
		op.errHandler = f
	}
}
