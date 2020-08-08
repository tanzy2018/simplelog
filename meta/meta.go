package meta

import (
	"strconv"
)

// Meta ...
type Meta interface {
	Key() []byte
	Value() []byte
}

type intMeta struct {
	key   string
	value int
}

func (im intMeta) Key() []byte {
	return wrapValue(im.key)
}

func (im intMeta) Value() []byte {
	return []byte(strconv.FormatInt(int64(im.value), 10))
}

type int64Meta struct {
	key   string
	value int64
}

func (im int64Meta) Key() []byte {
	return wrapValue(im.key)
}

func (im int64Meta) Value() []byte {
	return []byte(strconv.FormatInt(int64(im.value), 10))
}

type stringMeta struct {
	key   string
	value string
}

func (sm stringMeta) Key() []byte {
	return wrapValue(sm.key)
}

func (sm stringMeta) Value() []byte {
	return wrapValue(sm.value)
}

// Int ...
func Int(key string, value int) Meta {
	return intMeta{
		key:   key,
		value: value,
	}
}

// Int64 ...
func Int64(key string, value int64) Meta {
	return int64Meta{
		key:   key,
		value: value,
	}
}

// String ...
func String(key string, value string) Meta {
	return stringMeta{
		key:   key,
		value: value,
	}
}

func wrapValue(val string) []byte {
	b := make([]byte, 0, len(val))
	b = append(b, '"')
	b = append(b, []byte(val)...)
	b = append(b, '"')
	return b
}
