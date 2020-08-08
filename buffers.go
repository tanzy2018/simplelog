package simplelog

import (
	"bytes"
	"sync/atomic"

	"github.com/tanzy2018/simplelog/encode"
	"github.com/tanzy2018/simplelog/internal"
)

type syncBuffer struct {
	buf    *bytes.Buffer
	c      int32
	maxLen int
	l      *Log
}

func newSyncBuffers(l *Log, maxLen int, num int) []*syncBuffer {
	if num < 0 || num > 1024 {
		num = 10
	}
	bufs := make([]*syncBuffer, num)
	for i := range bufs {
		sb := &syncBuffer{
			maxLen: maxLen,
			buf:    &bytes.Buffer{},
			l:      l,
		}
		atomic.StoreInt32(&sb.c, 0)
		bufs[i] = sb
	}
	return bufs
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
	return sb.buf.Len() >= sb.maxLen
}

func (sb *syncBuffer) flushAsBytes() []byte {
	sb.lock()
	defer sb.unlock()
	b := sb.buf.Bytes()
	sb.buf.Reset()
	return b
}

type oneRecordBuffer struct {
	buf    *bytes.Buffer
	c      int32
	maxLen int
	l      *Log
}

func newOneRecordBuffers(l *Log, maxLen int, num int) []*oneRecordBuffer {
	if num < 0 || num > 1024 {
		num = 10
	}

	bufs := make([]*oneRecordBuffer, num)
	for i := range bufs {
		ob := &oneRecordBuffer{
			maxLen: maxLen,
			buf:    &bytes.Buffer{},
			l:      l,
		}
		atomic.StoreInt32(&ob.c, 0)
		bufs[i] = ob
	}

	return bufs
}

func (ob *oneRecordBuffer) lock() {
	for {
		if atomic.CompareAndSwapInt32(&ob.c, 0, 1) {
			return
		}
	}
}

func (ob *oneRecordBuffer) unlock() {
	for {
		if atomic.CompareAndSwapInt32(&ob.c, 1, 0) {
			return
		}
	}
}

func (ob *oneRecordBuffer) write(level LevelType, msg string, md []encode.Meta) {
	ob.lock()
	defer ob.unlock()
	ob.buf.Reset()
	//ob.buf.Grow(ob.maxLen)
	md0 := make([]encode.Meta, 0, 3)
	if UseTimeField {
		md0 = append(md0, timeMeta())
	}
	md0 = append(md0,
		levelMeta(level),
		msgMeta(msg))
	md0 = append(md0, globalHooks.Hooks()...)
	md0 = append(md0, ob.l.op.hook.Hooks()...)

	ob.writeLeftDelimiter()
	ob.writeCommonMeta(md0)
	ob.writeCustomMeta(md)
	ob.writeRightDelimiter()
	ob.writeEndDelimiter()

}

func (ob *oneRecordBuffer) writeCommonMeta(md []encode.Meta) {
	for i, msg := range md {
		if i != 0 {
			ob.writeFieldDelimiter()
		}
		ob.writeWrapper()
		ob.buf.Write(msg.Key())
		ob.writeWrapper()
		ob.writeKVDelimiter()
		if msg.Wrap() {
			ob.writeWrapper()
		}
		ob.buf.Write(msg.Value())
		if msg.Wrap() {
			ob.writeWrapper()
		}
	}
}

func (ob *oneRecordBuffer) writeCustomMeta(md []encode.Meta) {
	for _, msg := range md {
		lg := ob.buf.Len() + len(msg.Key()) + len(msg.Value())
		if lg >= ob.maxLen {
			return
		}
		ob.writeFieldDelimiter()
		ob.writeWrapper()
		ob.buf.Write(msg.Key())
		ob.writeWrapper()
		ob.writeKVDelimiter()
		if msg.Wrap() {
			ob.writeWrapper()
		}
		ob.buf.Write(msg.Value())
		if msg.Wrap() {
			ob.writeWrapper()
		}
	}
}

func (ob *oneRecordBuffer) writeWrapper() {
	ob.buf.WriteByte(valueWrapper)
}

func (ob *oneRecordBuffer) writeLeftDelimiter() {
	ob.buf.WriteByte(leftDelimiter)
}

func (ob *oneRecordBuffer) writeRightDelimiter() {
	ob.buf.WriteByte(rightDelimiter)
}

func (ob *oneRecordBuffer) writeEndDelimiter() {
	ob.buf.WriteByte(endDelimiter)
}

func (ob *oneRecordBuffer) writeFieldDelimiter() {
	ob.buf.WriteByte(fieldDelimiter)
}

func (ob *oneRecordBuffer) writeKVDelimiter() {
	ob.buf.WriteByte(kvDelimiter)
}

func (ob *oneRecordBuffer) flushAsBytes() []byte {
	ob.lock()
	defer ob.unlock()
	return ob.buf.Bytes()
}
