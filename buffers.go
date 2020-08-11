package simplelog

import (
	"bytes"
	"sync/atomic"

	"github.com/tanzy2018/simplelog/encode"
	"github.com/tanzy2018/simplelog/internal"
)

type syncBuffer struct {
	buf     *bytes.Buffer
	c       int32
	maxSize int
	l       *Log
}

func newSyncBuffers(l *Log, maxSize int) *syncBuffer {
	sb := &syncBuffer{
		maxSize: maxSize,
		buf:     &bytes.Buffer{},
		l:       l,
	}
	atomic.StoreInt32(&sb.c, 0)
	return sb
}

func (sb *syncBuffer) lock() {
	for {
		if atomic.CompareAndSwapInt32(&sb.c, 0, 1) {
			return
		}
	}
}

func (sb *syncBuffer) unlock() {
	for {
		if atomic.CompareAndSwapInt32(&sb.c, 1, 0) {
			return
		}
	}
}

func (sb *syncBuffer) write(b []byte) (sync bool) {
	sb.lock()
	defer sb.unlock()
	sb.buf.Grow(len(b))
	sb.buf.WriteString(internal.ToString(b))
	return sb.buf.Len() >= sb.maxSize
}

func (sb *syncBuffer) flushAsBytes() []byte {
	sb.lock()
	defer sb.unlock()
	b := sb.buf.Bytes()
	sb.buf.Reset()
	return b
}

type recordBuffer struct {
	buf     *bytes.Buffer
	c       int32
	maxSize int
	l       *Log
}

func newRecordBuffers(l *Log, maxSize int) *recordBuffer {
	rb := &recordBuffer{
		maxSize: maxSize,
		buf:     &bytes.Buffer{},
		l:       l,
	}
	atomic.StoreInt32(&rb.c, 0)
	return rb
}

func (rb *recordBuffer) lock() {
	for {
		if atomic.CompareAndSwapInt32(&rb.c, 0, 1) {
			return
		}
	}
}

func (rb *recordBuffer) unlock() {
	for {
		if atomic.CompareAndSwapInt32(&rb.c, 1, 0) {
			rb.buf.Reset()
			return
		}
	}
}

func (rb *recordBuffer) write(level LevelType, msg string, md []encode.Meta) []byte {
	rb.lock()
	defer rb.unlock()
	rb.buf.Reset()
	//rb.buf.Grow(rb.maxSize)
	md0 := make([]encode.Meta, 0, 3)
	if EnableTimeField {
		md0 = append(md0, timeMeta())
	}
	md0 = append(md0,
		levelMeta(level),
		msgMeta(msg))
	md0 = append(md0, rb.l.op.hook.Hooks()...)

	rb.writeLeftDelimiter()
	rb.writeCommonMeta(md0)
	rb.writeCustomMeta(md)
	if level == PANIC {
		rb.writeStackMeta()
	}
	rb.writeRightDelimiter()
	rb.writeEndDelimiter()
	return rb.buf.Bytes()

}

func (rb *recordBuffer) writeCommonMeta(md []encode.Meta) {
	for i, msg := range md {
		if i != 0 {
			rb.writeFieldDelimiter()
		}
		rb.writeWrapper()
		rb.buf.Write(msg.Key())
		rb.writeWrapper()
		rb.writeKVDelimiter()
		if msg.Wrap() {
			rb.writeWrapper()
		}
		rb.buf.Write(msg.Value())
		if msg.Wrap() {
			rb.writeWrapper()
		}
	}
}

func (rb *recordBuffer) writeCustomMeta(md []encode.Meta) {
	for _, msg := range md {
		size := rb.buf.Len() + len(msg.Key()) + len(msg.Value())
		if size >= rb.maxSize {
			return
		}
		rb.writeFieldDelimiter()
		rb.writeWrapper()
		rb.buf.Write(msg.Key())
		rb.writeWrapper()
		rb.writeKVDelimiter()
		if msg.Wrap() {
			rb.writeWrapper()
		}
		rb.buf.Write(msg.Value())
		if msg.Wrap() {
			rb.writeWrapper()
		}
	}
}

func (rb *recordBuffer) writeStackMeta() {
	md := stackMeta()
	rb.writeWrapper()
	rb.buf.Write(md.Key())
	rb.writeWrapper()
	rb.writeKVDelimiter()
	if md.Wrap() {
		rb.writeWrapper()
	}
	rb.buf.Write(md.Value())
	if md.Wrap() {
		rb.writeWrapper()
	}
}

func (rb *recordBuffer) writeWrapper() {
	rb.buf.WriteByte(valueWrapper)
}

func (rb *recordBuffer) writeLeftDelimiter() {
	rb.buf.WriteByte(leftDelimiter)
}

func (rb *recordBuffer) writeRightDelimiter() {
	rb.buf.WriteByte(rightDelimiter)
}

func (rb *recordBuffer) writeEndDelimiter() {
	rb.buf.WriteByte(endDelimiter)
}

func (rb *recordBuffer) writeFieldDelimiter() {
	rb.buf.WriteByte(fieldDelimiter)
}

func (rb *recordBuffer) writeKVDelimiter() {
	rb.buf.WriteByte(kvDelimiter)
}
