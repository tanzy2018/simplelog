package meta

import (
	"github.com/tanzy2018/simplelog/utils"
	"strconv"
)

var toBytes = utils.ToBytes

// Meta ...
type Meta interface {
	Key() []byte
	Value() []byte
	Wrap() bool
}

type imeta struct {
	key      []byte
	value    []byte
	needWrap bool
}

func (m imeta) Key() []byte {
	return m.key
}

func (m imeta) Value() []byte {
	return m.value
}

func (m imeta) Wrap() bool {
	return m.needWrap
}

// Int ...
func Int(key string, val int) Meta {
	return imeta{
		key:   toBytes(key),
		value: toBytes(strconv.FormatInt(int64(val), 10)),
		// value:    strconv.AppendInt([]byte{}, int64(val), 10),
		needWrap: false,
	}
}

// Int64 ...
func Int64(key string, val int64) Meta {
	return imeta{
		key:   toBytes(key),
		value: toBytes(strconv.FormatInt(val, 10)),
		// value:    strconv.AppendInt([]byte{}, val, 10),
		needWrap: false,
	}
}

// String ...
func String(key string, val string) Meta {
	return imeta{
		key:      toBytes(key),
		value:    toBytes(val),
		needWrap: true,
	}
}
