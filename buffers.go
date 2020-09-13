package simplelog

import (
	"bytes"
	"sync"

	"github.com/tanzy2018/simplelog/encode"
	"github.com/tanzy2018/simplelog/internal"
)

type syncBuffer struct {
	buf     *bytes.Buffer
	maxSize int
	l       *Log
	lo      *sync.Mutex
}

func newSyncBuffers(l *Log, maxSize int) *syncBuffer {
	sb := &syncBuffer{
		maxSize: maxSize,
		buf:     &bytes.Buffer{},
		l:       l,
		lo:      new(sync.Mutex),
	}
	return sb
}

func (sb *syncBuffer) lock() {
	sb.lo.Lock()
}

func (sb *syncBuffer) unlock() {
	sb.lo.Unlock()
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
	maxSize int
	l       *Log
	lo      *sync.Mutex
}

func newRecordBuffers(l *Log, maxSize int) *recordBuffer {
	rb := &recordBuffer{
		maxSize: maxSize,
		buf:     &bytes.Buffer{},
		l:       l,
		lo:      new(sync.Mutex),
	}
	return rb
}

func (rb *recordBuffer) lock() {
	rb.lo.Lock()
}

func (rb *recordBuffer) unlock() {
	rb.buf.Reset()
	rb.lo.Unlock()
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
