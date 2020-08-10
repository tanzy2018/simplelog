package encode

import (
	"encoding/json"
	"github.com/tanzy2018/simplelog/internal"
	"reflect"
	"strconv"
)

var toBytes = internal.ToBytes
var toString = internal.ToString

type imeta struct {
	key   []byte
	value []byte
	wrap  bool
}

func (m imeta) Key() []byte {
	return m.key
}

func (m imeta) Value() []byte {
	return m.value
}

func (m imeta) Wrap() bool {
	return m.wrap
}

func (m imeta) IsNil() bool {
	return false
}

type nullImeta string

func (n nullImeta) Key() []byte {
	return []byte(n)
}

func (n nullImeta) Value() []byte {
	return []byte("null")
}

func (n nullImeta) Wrap() bool {
	return false
}

func (n nullImeta) IsNil() bool {
	return true
}

type emptyArrayImeta string

func (e emptyArrayImeta) Key() []byte {
	return []byte(e)
}

func (e emptyArrayImeta) Value() []byte {
	return []byte("[]")
}

func (e emptyArrayImeta) Wrap() bool {
	return false
}

func (e emptyArrayImeta) IsNil() bool {
	return false
}

type emptyStrImeta string

func (e emptyStrImeta) Key() []byte {
	return []byte(e)
}

func (e emptyStrImeta) Value() []byte {
	return []byte("")
}

func (e emptyStrImeta) Wrap() bool {
	return true
}

func (e emptyStrImeta) IsNil() bool {
	return false
}

// Int ...
func Int(key string, val int) Meta {
	return Int64(key, int64(val))
}

// Ints ...
func Ints(key string, ints []int) Meta {
	if len(ints) == 0 {
		return emptyArrayImeta(key)
	}
	vals := make([]byte, 0, 2+len(ints)*11)
	vals = append(vals, '[')
	vals = strconv.AppendInt(vals, int64(ints[0]), 10)
	if len(ints) > 1 {
		for _, v := range ints {
			vals = append(vals, ',')
			vals = strconv.AppendInt(vals, int64(v), 10)
		}
	}
	vals = append(vals, ']')
	return imeta{
		key:   toBytes(key),
		value: vals[:len(vals)],
	}
}

// Int8 ...
func Int8(key string, val int8) Meta {
	return Int64(key, int64(val))
}

// Int8s ...
func Int8s(key string, ints []int8) Meta {
	if len(ints) == 0 {
		return emptyArrayImeta(key)
	}
	vals := make([]byte, 0, 2+len(ints)*4)
	vals = append(vals, '[')
	vals = strconv.AppendInt(vals, int64(ints[0]), 10)
	if len(ints) > 1 {
		for _, v := range ints[1:] {
			vals = append(vals, ',')
			vals = strconv.AppendInt(vals, int64(v), 10)
		}
	}
	vals = append(vals, ']')
	return imeta{
		key:   toBytes(key),
		value: vals[:len(vals)],
	}
}

// Int16 ...
func Int16(key string, val int16) Meta {
	return Int64(key, int64(val))
}

// Int16s ...
func Int16s(key string, ints []int16) Meta {
	if len(ints) == 0 {
		return emptyArrayImeta(key)
	}
	vals := make([]byte, 0, 2+len(ints)*6)
	vals = append(vals, '[')
	vals = strconv.AppendInt(vals, int64(ints[0]), 10)
	if len(ints) > 1 {
		for _, v := range ints[1:] {
			vals = append(vals, ',')
			vals = strconv.AppendInt(vals, int64(v), 10)
		}
	}

	vals = append(vals, ']')
	return imeta{
		key:   toBytes(key),
		value: vals[:len(vals)],
	}
}

// Int32 ...
func Int32(key string, val int32) Meta {
	return Int64(key, int64(val))
}

// Int32s ...
func Int32s(key string, ints []int32) Meta {
	if len(ints) == 0 {
		return emptyArrayImeta(key)
	}
	vals := make([]byte, 0, 2+len(ints)*11)
	vals = append(vals, '[')
	vals = strconv.AppendInt(vals, int64(ints[0]), 10)
	if len(ints) > 1 {
		for _, v := range ints[1:] {
			vals = append(vals, ',')
			vals = strconv.AppendInt(vals, int64(v), 10)
		}
	}
	vals = append(vals, ']')
	return imeta{
		key:   toBytes(key),
		value: vals[:len(vals)],
	}
}

// Int64 ...
func Int64(key string, val int64) Meta {
	return imeta{
		key:   toBytes(key),
		value: strconv.AppendInt(make([]byte, 0, 10), val, 10),
	}
}

// Int64s ...
func Int64s(key string, ints []int64) Meta {
	if len(ints) == 0 {
		return emptyArrayImeta(key)
	}
	vals := make([]byte, 0, 2+len(ints)*11)
	vals = append(vals, '[')
	vals = strconv.AppendInt(vals, ints[0], 10)
	if len(ints) > 1 {
		for _, v := range ints[1:] {
			vals = append(vals, ',')
			vals = strconv.AppendInt(vals, v, 10)
		}
	}
	vals = append(vals, ']')
	return imeta{
		key:   toBytes(key),
		value: vals[:len(vals)],
	}
}

// Uint ...
func Uint(key string, val uint) Meta {
	return Uint64(key, uint64(val))
}

// Uints ...
func Uints(key string, ints []uint) Meta {
	if len(ints) == 0 {
		return emptyArrayImeta(key)
	}
	vals := make([]byte, 0, 2*len(ints)*11)
	vals = append(vals, '[')
	vals = strconv.AppendUint(vals, uint64(ints[0]), 10)
	if len(ints) > 1 {
		for _, v := range ints[1:] {
			vals = append(vals, ',')
			vals = strconv.AppendUint(vals, uint64(v), 10)
		}
	}
	vals = append(vals, ']')
	return imeta{
		key:   toBytes(key),
		value: vals[:len(vals)],
	}
}

// Uint8 ...
func Uint8(key string, val uint8) Meta {
	return imeta{
		key:   toBytes(key),
		value: strconv.AppendUint(make([]byte, 0, 3), uint64(val), 10),
	}
}

// Uint8s ...
func Uint8s(key string, ints []uint8) Meta {
	if len(ints) == 0 {
		return emptyArrayImeta(key)
	}
	vals := make([]byte, 0, 2+len(ints)*4)
	vals = append(vals, '[')
	vals = strconv.AppendUint(vals, uint64(ints[0]), 10)
	if len(vals) > 1 {
		for _, v := range ints[1:] {
			vals = append(vals, ',')
			vals = strconv.AppendUint(vals, uint64(v), 10)
		}
	}
	vals = append(vals, ']')
	return imeta{
		key:   toBytes(key),
		value: vals[:len(vals)],
	}
}

// Bytes ... Will output as string
func Bytes(key string, bs []byte) Meta {
	if len(bs) == 0 {
		return emptyStrImeta(key)
	}
	return imeta{
		key:   toBytes(key),
		value: bs,
		wrap:  true,
	}
}

// Uint16 ...
func Uint16(key string, val uint16) Meta {
	return Uint64(key, uint64(val))
}

// Uint16s ...
func Uint16s(key string, ints []uint16) Meta {
	if len(ints) == 0 {
		return emptyArrayImeta(key)
	}
	vals := make([]byte, 0, 2+len(ints)*6)
	vals = append(vals, '[')
	vals = strconv.AppendUint(vals, uint64(ints[0]), 10)
	if len(ints) > 1 {
		for _, v := range ints[1:] {
			vals = append(vals, ',')
			vals = strconv.AppendUint(vals, uint64(v), 10)
		}
	}
	vals = append(vals, ']')
	return imeta{
		key:   toBytes(key),
		value: vals[:len(vals)],
	}
}

// Uint32 ...
func Uint32(key string, val uint32) Meta {
	return Uint64(key, uint64(val))
}

// Uint32s ...
func Uint32s(key string, ints []uint32) Meta {
	if len(ints) == 0 {
		return emptyArrayImeta(key)
	}
	vals := make([]byte, 0, 2+len(ints)*11)
	vals = append(vals, '[')
	vals = strconv.AppendUint(vals, uint64(ints[0]), 10)
	if len(ints) > 1 {
		for _, v := range ints[1:] {
			vals = append(vals, ',')
			vals = strconv.AppendUint(vals, uint64(v), 10)
		}
	}
	vals = append(vals, ']')
	return imeta{
		key:   toBytes(key),
		value: vals[:len(vals)],
	}
}

// Uint64 ...
func Uint64(key string, val uint64) Meta {
	return imeta{
		key:   toBytes(key),
		value: strconv.AppendUint(make([]byte, 0, 10), val, 10),
	}
}

// Uint64s ...
func Uint64s(key string, ints []uint64) Meta {
	if len(ints) == 0 {
		return emptyArrayImeta(key)
	}
	vals := make([]byte, 0, 2+len(ints)*11)
	vals = append(vals, '[')
	vals = strconv.AppendUint(vals, uint64(ints[0]), 10)
	if len(ints) > 1 {
		for _, v := range ints[1:] {
			vals = append(vals, ',')
			vals = strconv.AppendUint(vals, v, 10)
		}
	}
	vals = append(vals, ']')
	return imeta{
		key:   toBytes(key),
		value: vals[:len(vals)],
	}
}

// Float32 ...
func Float32(key string, val float32) Meta {
	return imeta{
		key:   toBytes(key),
		value: strconv.AppendFloat(make([]byte, 0, 8), float64(val), 'g', 5, 32),
	}
}

// Float32s ...
func Float32s(key string, ints []float32) Meta {
	if len(ints) == 0 {
		return emptyArrayImeta(key)
	}
	vals := make([]byte, 0, 2+8*len(ints))
	vals = append(vals, '[')
	vals = strconv.AppendFloat(vals, float64(ints[0]), 'g', 5, 32)
	if len(ints) > 1 {
		for _, f := range ints {
			vals = append(vals, ',')
			vals = strconv.AppendFloat(vals, float64(f), 'g', 5, 32)
		}
	}
	vals = append(vals, ']')
	return imeta{
		key:   toBytes(key),
		value: vals[:len(vals)],
	}
}

// Float64 ...
func Float64(key string, val float64) Meta {
	return imeta{
		key:   toBytes(key),
		value: strconv.AppendFloat(make([]byte, 0, 8), val, 'g', 5, 64),
	}
}

// Float64s ...
func Float64s(key string, ints []float64) Meta {
	if len(ints) == 0 {
		return emptyArrayImeta(key)
	}
	vals := make([]byte, 0, 2+8*len(ints))
	vals = append(vals, '[')
	vals = strconv.AppendFloat(vals, ints[0], 'g', 5, 64)
	if len(ints) > 1 {
		for _, v := range ints[1:] {
			vals = append(vals, ',')
			vals = strconv.AppendFloat(vals, v, 'g', 5, 64)
		}
	}

	vals = append(vals, ']')
	return imeta{
		key:   toBytes(key),
		value: vals[:len(vals)],
	}
}

// String ...
func String(key string, val string) Meta {
	return imeta{
		key:   toBytes(key),
		value: toBytes(val),
		wrap:  true,
	}
}

// Strings ...
func Strings(key string, strs []string) Meta {
	if len(strs) == 0 {
		return emptyArrayImeta(key)
	}
	vals := make([]byte, 0, 2)
	vals = append(vals, '[')
	vals = append(vals, '"')
	vals = append(vals, toBytes(strs[0])...)
	vals = append(vals, '"')
	if len(strs) > 1 {
		for i := 1; i < len(strs); i++ {
			vals = append(vals, ',')
			vals = append(vals, '"')
			vals = append(vals, toBytes(strs[i])...)
			vals = append(vals, '"')
		}
	}
	vals = append(vals, ']')
	return imeta{
		key:   toBytes(key),
		value: vals[:len(vals)],
	}
}

// Bool ...
func Bool(key string, b bool) Meta {
	return imeta{
		key:   toBytes(key),
		value: toBytes(strconv.FormatBool(b)),
	}
}

// Bools ...
func Bools(key string, bs []bool) Meta {
	if len(bs) == 0 {
		return emptyArrayImeta(key)
	}
	vals := make([]byte, 0, 2+6*len(bs))
	vals = append(vals, '[')
	vals = strconv.AppendBool(vals, bs[0])
	if len(bs) > 1 {
		for _, b := range bs {
			vals = append(vals, ',')
			vals = strconv.AppendBool(vals, b)
		}
	}
	vals = append(vals, ']')
	return imeta{
		key:   toBytes(key),
		value: vals[:len(vals)],
	}
}

// Object ... map(*map) or struct(*struct)
// func any_bak(key string, val interface{}) Meta {
// 	if val == nil {
// 		return nullImeta(key)
// 	}

// 	v := reflect.ValueOf(val)
// 	if v.Kind() == reflect.Ptr {
// 		if !v.Elem().IsValid() {
// 			return nullImeta(key)
// 		}
// 		v = reflect.ValueOf(v.Elem().Interface())
// 	}

// 	kind := v.Kind()
// 	if kind == reflect.Struct {
// 		buf := make([]byte, 0, 2)
// 		buf = append(buf, '{')
// 		for i := 0; i < v.NumField(); i++ {
// 			if i > 0 {
// 				buf = append(buf, ',')
// 			}
// 			field := v.Type().Field(i)
// 			name := field.Tag.Get("json")
// 			if len(name) == 0 {
// 				name = field.Name
// 			}
// 			md := Any(name, v.Field(i).Interface())
// 			buf = append(buf, '"')
// 			buf = append(buf, md.Key()...)
// 			buf = append(buf, '"')
// 			buf = append(buf, ':')
// 			if md.Wrap() {
// 				buf = append(buf, '"')
// 			}
// 			buf = append(buf, md.Value()...)
// 			if md.Wrap() {
// 				buf = append(buf, '"')
// 			}
// 		}
// 		buf = append(buf, '}')

// 		return imeta{
// 			key:   toBytes(key),
// 			value: buf,
// 		}
// 	}

// 	if kind == reflect.Map {
// 		if v.IsNil() {
// 			return nullImeta(key)
// 		}
// 		mIter := v.MapRange()
// 		buf := make([]byte, 0, 2)
// 		i := 0
// 		buf = append(buf, '{')
// 		for mIter.Next() {
// 			if i > 0 {
// 				buf = append(buf, ',')
// 			}
// 			md := Any(toString(append([]byte{}, Any("", mIter.Key().Interface()).Value()...)),
// 				mIter.Value().Interface())
// 			buf = append(buf, '"')
// 			buf = append(buf, md.Key()...)
// 			buf = append(buf, '"')
// 			buf = append(buf, ':')
// 			if md.Wrap() {
// 				buf = append(buf, '"')
// 			}
// 			buf = append(buf, md.Value()...)
// 			if md.Wrap() {
// 				buf = append(buf, '"')
// 			}
// 			i++
// 		}
// 		buf = append(buf, '}')

// 		return imeta{
// 			key:   toBytes(key),
// 			value: buf,
// 		}
// 	}

// 	if kind == reflect.Slice {
// 		if v.IsNil() || v.Len() == 0 {
// 			return emptyArrayImeta(key)
// 		}

// 		if u8s, ok := v.Interface().([]uint8); ok {
// 			return Uint8s(key, u8s)
// 		}

// 		buf := make([]byte, 0, 2)
// 		buf = append(buf, '[')
// 		for i := 0; i < v.Len(); i++ {
// 			if i > 0 {
// 				buf = append(buf, ',')
// 			}
// 			md := Any("", v.Index(i).Interface())
// 			if md.Wrap() {
// 				buf = append(buf, '"')
// 			}
// 			buf = append(buf, md.Value()...)
// 			if md.Wrap() {
// 				buf = append(buf, '"')
// 			}
// 		}
// 		buf = append(buf, ']')
// 		return imeta{
// 			key:   toBytes(key),
// 			value: buf,
// 		}
// 	}

// 	if kind == reflect.Array {
// 		buf := make([]byte, 0, 2)
// 		buf = append(buf, '[')
// 		for i := 0; i < v.Len(); i++ {
// 			if i > 0 {
// 				buf = append(buf, ',')
// 			}
// 			md := Any("", v.Index(i).Interface())
// 			if md.Wrap() {
// 				buf = append(buf, '"')
// 			}
// 			buf = append(buf, md.Value()...)
// 			if md.Wrap() {
// 				buf = append(buf, '"')
// 			}
// 		}
// 		buf = append(buf, ']')
// 		return imeta{
// 			key:   toBytes(key),
// 			value: buf,
// 		}
// 	}

// 	return nullImeta(key)

// }

func any(key string, val interface{}) Meta {
	if val == nil {
		return nullImeta(key)
	}

	v := reflect.ValueOf(val)
	if v.Kind() == reflect.Ptr {
		if !v.Elem().IsValid() {
			return nullImeta(key)
		}
		v = reflect.ValueOf(v.Elem().Interface())
	}

	kind := v.Kind()
	if kind == reflect.Struct {
		buf, err := json.Marshal(v.Interface())
		if err != nil {
			return nullImeta(key)
		}
		return imeta{
			key:   toBytes(key),
			value: buf,
		}
	}

	if kind == reflect.Map {
		if v.IsNil() {
			return nullImeta(key)
		}
		buf, err := json.Marshal(v.Interface())
		if err != nil {
			return nullImeta(key)
		}
		return imeta{
			key:   toBytes(key),
			value: buf,
		}
	}

	if kind == reflect.Slice {
		if v.IsNil() || v.Len() == 0 {
			return emptyArrayImeta(key)
		}

		buf := make([]byte, 0, 2)
		buf = append(buf, '[')

		md := Any("", v.Index(0).Interface())
		if md.Wrap() {
			buf = append(buf, '"')
			buf = append(buf, md.Value()...)
			buf = append(buf, '"')
		} else {
			buf = append(buf, md.Value()...)
		}

		for i := 1; i < v.Len(); i++ {
			buf = append(buf, ',')
			md := Any("", v.Index(i).Interface())
			if md.Wrap() {
				buf = append(buf, '"')
				buf = append(buf, md.Value()...)
				buf = append(buf, '"')
			} else {
				buf = append(buf, md.Value()...)
			}
		}

		buf = append(buf, ']')
		return imeta{
			key:   toBytes(key),
			value: buf,
		}
	}

	if kind == reflect.Array {
		if v.Len() == 0 {
			return emptyArrayImeta(key)
		}

		buf := make([]byte, 0, 2)
		buf = append(buf, '[')
		md := Any("", v.Index(0).Interface())
		if md.Wrap() {
			buf = append(buf, '"')
			buf = append(buf, md.Value()...)
			buf = append(buf, '"')
		} else {
			buf = append(buf, md.Value()...)
		}

		for i := 1; i < v.Len(); i++ {
			buf = append(buf, ',')
			md := Any("", v.Index(i).Interface())
			if md.Wrap() {
				buf = append(buf, '"')
				buf = append(buf, md.Value()...)
				buf = append(buf, '"')
			} else {
				buf = append(buf, md.Value()...)
			}
		}
		buf = append(buf, ']')
		return imeta{
			key:   toBytes(key),
			value: buf,
		}
	}
	return nullImeta(key)
}

// Any ...
func Any(key string, val interface{}) Meta {
	switch val.(type) {
	default:
		return any(key, val)
	case int:
		return Int(key, val.(int))
	case int32:
		return Int32(key, val.(int32))
	case int64:
		return Int64(key, val.(int64))
	case float32:
		return Float32(key, val.(float32))
	case float64:
		return Float64(key, val.(float64))
	case string:
		return String(key, val.(string))
	case bool:
		return Bool(key, val.(bool))
	case uint32:
		return Uint32(key, val.(uint32))
	case uint64:
		return Uint64(key, val.(uint64))

	case int8:
		return Int8(key, val.(int8))
	case int16:
		return Int16(key, val.(int16))
	case uint:
		return Uint(key, val.(uint))
	case uint8:
		return Uint8(key, val.(uint8))
	case uint16:
		return Uint16(key, val.(uint16))
	case *int:
		v := val.(*int)
		if v == nil {
			return nullImeta(key)
		}
		return Int(key, *v)
	case *int8:
		v := val.(*int8)
		if v == nil {
			return nullImeta(key)
		}
		return Int8(key, *v)
	case *int16:
		v := val.(*int16)
		if v == nil {
			return nullImeta(key)
		}
		return Int16(key, *v)
	case *int32:
		v := val.(*int32)
		if v == nil {
			return nullImeta(key)
		}
		return Int32(key, *v)
	case *int64:
		v := val.(*int64)
		if v == nil {
			return nullImeta(key)
		}
		return Int64(key, *v)
	case *uint:
		v := val.(*uint)
		if v == nil {
			return nullImeta(key)
		}
		return Uint(key, *v)
	case *uint8:
		v := val.(*uint8)
		if v == nil {
			return nullImeta(key)
		}
		return Uint8(key, *v)
	case *uint16:
		v := val.(*uint16)
		if v == nil {
			return nullImeta(key)
		}
		return Uint16(key, *v)
	case *uint32:
		v := val.(*uint32)
		if v == nil {
			return nullImeta(key)
		}
		return Uint32(key, *v)
	case *uint64:
		v := val.(*uint64)
		if v == nil {
			return nullImeta(key)
		}
		return Uint64(key, *v)
	case *float32:
		v := val.(*float32)
		if v == nil {
			return nullImeta(key)
		}
		return Float32(key, *v)
	case *float64:
		v := val.(*float64)
		if v == nil {
			return nullImeta(key)
		}
		return Float64(key, *v)
	case *string:
		v := val.(*string)
		if v == nil {
			return nullImeta(key)
		}
		return String(key, *v)
	case *bool:
		v := val.(*bool)
		if v == nil {
			return nullImeta(key)
		}
		return Bool(key, *v)
	}
}
