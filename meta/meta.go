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

type _imeta struct {
	key      []byte
	value    []byte
	needWrap bool
}

func (m *_imeta) Key() []byte {
	return m.key
}

func (m *_imeta) Value() []byte {
	return m.value
}

func (m *_imeta) Wrap() bool {
	return m.needWrap
}

// Int ...
func Int(key string, val int) Meta {
	return &_imeta{
		key:      toBytes(key),
		value:    toBytes(strconv.FormatInt(int64(val), 10)),
		needWrap: false,
	}
}

// Int64 ...
func Int64(key string, val int64) Meta {
	return &_imeta{
		key:      toBytes(key),
		value:    toBytes(strconv.FormatInt(val, 10)),
		needWrap: false,
	}
}

// String ...
func String(key string, val string) Meta {
	return &_imeta{
		key:      toBytes(key),
		value:    toBytes(val),
		needWrap: true,
	}
}
