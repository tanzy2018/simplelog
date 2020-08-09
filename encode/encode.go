package encode

import (
	"encoding/json"
	"github.com/tanzy2018/simplelog/internal"
	"reflect"
	"strconv"
)

var toBytes = internal.ToBytes
var toString = internal.ToString

const null = "null"

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
	return Int64(key, int64(val))
}

// Ints ...
func Ints(key string, ints []int) Meta {
	if len(ints) == 0 {
		return emptyArrayImeta(key)
	}
	vals := make([]byte, 0, 2)
	vals = append(vals, '[')
	for i := range ints {
		if i > 0 {
			vals = append(vals, ',')
		}
		vals = append(vals, toBytes(strconv.FormatInt(int64(ints[i]), 10))...)
	}
	vals = append(vals, ']')
	return imeta{
		key:      toBytes(key),
		value:    vals[:len(vals)],
		needWrap: false,
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
	vals := make([]byte, 0, 2)
	vals = append(vals, '[')
	for i := range ints {
		if i > 0 {
			vals = append(vals, ',')
		}
		vals = append(vals, toBytes(strconv.FormatInt(int64(ints[i]), 10))...)
	}
	vals = append(vals, ']')
	return imeta{
		key:      toBytes(key),
		value:    vals[:len(vals)],
		needWrap: false,
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
	vals := make([]byte, 0, 2)
	vals = append(vals, '[')
	for i := range ints {
		if i > 0 {
			vals = append(vals, ',')
		}
		vals = append(vals, toBytes(strconv.FormatInt(int64(ints[i]), 10))...)
	}
	vals = append(vals, ']')
	return imeta{
		key:      toBytes(key),
		value:    vals[:len(vals)],
		needWrap: false,
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
	vals := make([]byte, 0, 2)
	vals = append(vals, '[')
	for i := range ints {
		if i > 0 {
			vals = append(vals, ',')
		}
		vals = append(vals, toBytes(strconv.FormatInt(int64(ints[i]), 10))...)
	}
	vals = append(vals, ']')
	return imeta{
		key:      toBytes(key),
		value:    vals[:len(vals)],
		needWrap: false,
	}
}

// Int64 ...
func Int64(key string, val int64) Meta {
	return imeta{
		key:      toBytes(key),
		value:    toBytes(strconv.FormatInt(val, 10)),
		needWrap: false,
	}
}

// Int64s ...
func Int64s(key string, ints []int64) Meta {
	if len(ints) == 0 {
		return emptyArrayImeta(key)
	}
	vals := make([]byte, 0, 2)
	vals = append(vals, '[')
	for i := range ints {
		if i > 0 {
			vals = append(vals, ',')
		}
		vals = append(vals, toBytes(strconv.FormatInt(ints[i], 10))...)
	}
	vals = append(vals, ']')
	return imeta{
		key:      toBytes(key),
		value:    vals[:len(vals)],
		needWrap: false,
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
	vals := make([]byte, 0, 2)
	vals = append(vals, '[')
	for i := range ints {
		if i > 0 {
			vals = append(vals, ',')
		}
		vals = append(vals, toBytes(strconv.FormatUint(uint64(ints[i]), 10))...)
	}
	vals = append(vals, ']')
	return imeta{
		key:      toBytes(key),
		value:    vals[:len(vals)],
		needWrap: false,
	}
}

// Uint8 ...
func Uint8(key string, val uint8) Meta {
	return Uint64(key, uint64(val))
}

// Uint8s ...
func Uint8s(key string, ints []uint8) Meta {
	if len(ints) == 0 {
		return emptyArrayImeta(key)
	}
	vals := make([]byte, 0, 2)
	vals = append(vals, '[')
	for i := range ints {
		if i > 0 {
			vals = append(vals, ',')
		}
		vals = append(vals, toBytes(strconv.FormatUint(uint64(ints[i]), 10))...)
	}
	vals = append(vals, ']')
	return imeta{
		key:      toBytes(key),
		value:    vals[:len(vals)],
		needWrap: false,
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
	vals := make([]byte, 0, 2)
	vals = append(vals, '[')
	for i := range ints {
		if i > 0 {
			vals = append(vals, ',')
		}
		vals = append(vals, toBytes(strconv.FormatUint(uint64(ints[i]), 10))...)
	}
	vals = append(vals, ']')
	return imeta{
		key:      toBytes(key),
		value:    vals[:len(vals)],
		needWrap: false,
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
	vals := make([]byte, 0, 2)
	vals = append(vals, '[')
	for i := range ints {
		if i > 0 {
			vals = append(vals, ',')
		}
		vals = append(vals, toBytes(strconv.FormatUint(uint64(ints[i]), 10))...)
	}
	vals = append(vals, ']')
	return imeta{
		key:      toBytes(key),
		value:    vals[:len(vals)],
		needWrap: false,
	}
}

// Uint64 ...
func Uint64(key string, val uint64) Meta {
	return imeta{
		key:      toBytes(key),
		value:    toBytes(strconv.FormatUint(val, 10)),
		needWrap: false,
	}
}

// Uint64s ...
func Uint64s(key string, ints []uint64) Meta {
	if len(ints) == 0 {
		return emptyArrayImeta(key)
	}
	vals := make([]byte, 0, 2)
	vals = append(vals, '[')
	for i := range ints {
		if i > 0 {
			vals = append(vals, ',')
		}
		vals = append(vals, toBytes(strconv.FormatUint(ints[i], 10))...)
	}
	vals = append(vals, ']')
	return imeta{
		key:      toBytes(key),
		value:    vals[:len(vals)],
		needWrap: false,
	}
}

// Float32 ...
func Float32(key string, val float32) Meta {
	return Float64(key, float64(val))
}

// Float32s ...
func Float32s(key string, ints []float32) Meta {
	if len(ints) == 0 {
		return emptyArrayImeta(key)
	}
	vals := make([]byte, 0, 2)
	vals = append(vals, '[')
	for i := range ints {
		if i > 0 {
			vals = append(vals, ',')
		}
		vals = append(vals, toBytes(strconv.FormatFloat(float64(ints[i]), 'g', 5, 32))...)
	}
	vals = append(vals, ']')
	return imeta{
		key:      toBytes(key),
		value:    vals[:len(vals)],
		needWrap: false,
	}
}

// Float64 ...
func Float64(key string, val float64) Meta {
	return imeta{
		key:      toBytes(key),
		value:    toBytes(strconv.FormatFloat(float64(val), 'g', 5, 64)),
		needWrap: false,
	}
}

// Float64s ...
func Float64s(key string, ints []float64) Meta {
	if len(ints) == 0 {
		return emptyArrayImeta(key)
	}
	vals := make([]byte, 0, 2)
	vals = append(vals, '[')
	for i := range ints {
		if i > 0 {
			vals = append(vals, ',')
		}
		vals = append(vals, toBytes(strconv.FormatFloat(float64(ints[i]), 'g', 5, 64))...)
	}
	vals = append(vals, ']')
	return imeta{
		key:      toBytes(key),
		value:    vals[:len(vals)],
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

// Strings ...
func Strings(key string, strs []string) Meta {
	if len(strs) == 0 {
		return emptyArrayImeta(key)
	}
	vals := make([]byte, 0, 2)
	vals = append(vals, '[')
	for i := range strs {
		if i > 0 {
			vals = append(vals, ',')
		}
		vals = append(vals, toBytes(strs[i])...)
	}
	vals = append(vals, ']')
	return imeta{
		key:      toBytes(key),
		value:    vals[:len(vals)],
		needWrap: false,
	}
}

// Bool ...
func Bool(key string, b bool) Meta {
	return imeta{
		key:      toBytes(key),
		value:    toBytes(strconv.FormatBool(b)),
		needWrap: false,
	}
}

// Bools ...
func Bools(key string, bs []bool) Meta {
	if len(bs) == 0 {
		return emptyArrayImeta(key)
	}
	vals := make([]byte, 0, 2+5*len(bs))
	vals = append(vals, '[')
	for i := range bs {
		if i > 0 {
			vals = append(vals, ',')
		}
		vals = append(vals, toBytes(strconv.FormatBool(bs[i]))...)
	}
	vals = append(vals, ']')
	return imeta{
		key:      toBytes(key),
		value:    vals[:len(vals)],
		needWrap: false,
	}
}

// Object ... map(*map) or struct(*struct)
func Object(key string, val interface{}) Meta {
	if val == nil {
		return nullImeta(key)
	}

	v := reflect.ValueOf(val)
	if v.Kind() == reflect.Ptr {
		v = reflect.ValueOf(v.Elem())
	}

	if v.Kind() != reflect.Map || v.Kind() != reflect.Struct {
		return nullImeta(key)
	}

	b, err := json.Marshal(v.Interface())
	if err != nil {
		return nullImeta(key)
	}

	return imeta{
		key:      toBytes(key),
		value:    b,
		needWrap: false,
	}
}

// Objects ...[map(*map),] or [struct(*struct),]
func Objects(key string, vals []interface{}) Meta {
	if len(vals) == 0 {
		return emptyArrayImeta(key)
	}

	bs := make([]byte, 0, 2)
	bs = append(bs, '[')
	for i := range vals {
		if i > 0 {
			bs = append(bs, ',')
		}
		bs = append(bs, Object("", vals[i]).Value()...)
	}
	bs = append(bs, ']')
	return imeta{
		key:      toBytes(key),
		value:    bs,
		needWrap: false,
	}
}

// Any ...
func Any(key string, any interface{}) Meta {

	if _imeta := Object(key, any); null != toString(_imeta.Value()) {
		return _imeta
	}

	if v, ok := any.([]interface{}); ok {
		_imeta := Objects(key, v)
		if null != toString(_imeta.Value()) {
			return _imeta
		}
	}

	switch any.(type) {

	case int:
		return Int(key, any.(int))
	case int8:
		return Int8(key, any.(int8))
	case int16:
		return Int16(key, any.(int16))
	case int32:
		return Int32(key, any.(int32))
	case int64:
		return Int64(key, any.(int64))
	case uint:
		return Uint(key, any.(uint))
	case uint8:
		return Uint8(key, any.(uint8))
	case uint16:
		return Uint16(key, any.(uint16))
	case uint32:
		return Uint32(key, any.(uint32))
	case uint64:
		return Uint64(key, any.(uint64))
	case float32:
		return Float32(key, any.(float32))
	case float64:
		return Float64(key, any.(float64))
	case string:
		return String(key, any.(string))
	case bool:
		return Bool(key, any.(bool))
	case []int:
		return Ints(key, any.([]int))
	case []int8:
		return Int8s(key, any.([]int8))
	case []int16:
		return Int16s(key, any.([]int16))
	case []int32:
		return Int32s(key, any.([]int32))
	case []int64:
		return Int64s(key, any.([]int64))
	case []uint:
		return Uints(key, any.([]uint))
	case []uint8:
		return Uint8s(key, any.([]uint8))
	case []uint16:
		return Uint16s(key, any.([]uint16))
	case []uint32:
		return Uint32s(key, any.([]uint32))
	case []uint64:
		return Uint64s(key, any.([]uint64))
	case []float32:
		return Float32s(key, any.([]float32))
	case []float64:
		return Float64s(key, any.([]float64))
	case []string:
		return Strings(key, any.([]string))
	case []bool:
		return Bools(key, any.([]bool))

	case *int:
		return Int(key, *(any.(*int)))
	case *int8:
		return Int8(key, *(any.(*int8)))
	case *int16:
		return Int16(key, *(any.(*int16)))
	case *int32:
		return Int32(key, *(any.(*int32)))
	case *int64:
		return Int64(key, *(any.(*int64)))
	case *uint:
		return Uint(key, *(any.(*uint)))
	case *uint8:
		return Uint8(key, *(any.(*uint8)))
	case *uint16:
		return Uint16(key, *(any.(*uint16)))
	case *uint32:
		return Uint32(key, *(any.(*uint32)))
	case *uint64:
		return Uint64(key, *(any.(*uint64)))
	case *float32:
		return Float32(key, *(any.(*float32)))
	case *float64:
		return Float64(key, *(any.(*float64)))
	case *string:
		return String(key, *(any.(*string)))
	case *bool:
		return Bool(key, *(any.(*bool)))

	case []*int:
		return Ints(key, intPtrs2ints(any.([]*int)))
	case []*int8:
		return Int8s(key, int8Ptrs2int8s(any.([]*int8)))
	case []*int16:
		return Int16s(key, int16Ptrs2int16s(any.([]*int16)))
	case []*int32:
		return Int32s(key, int32Ptrs2int32s(any.([]*int32)))
	case []*int64:
		return Int64s(key, int64Ptrs2int64s(any.([]*int64)))
	case []*uint:
		return Uints(key, uintPtrs2uints(any.([]*uint)))
	case []*uint8:
		return Uint8s(key, uint8Ptrs2uint8s(any.([]*uint8)))
	case []*uint16:
		return Uint16s(key, uint16Ptrs2uint16s(any.([]*uint16)))
	case []*uint32:
		return Uint32s(key, uint32Ptrs2uint32s(any.([]*uint32)))
	case []*uint64:
		return Uint64s(key, uint64Ptrs2uint64s(any.([]*uint64)))
	case []*float32:
		return Float32s(key, float32Ptrs2Float32s(any.([]*float32)))
	case []*float64:
		return Float64s(key, float64Ptrs2Float64s(any.([]*float64)))
	case []*string:
		return Strings(key, stringPtrs2Strings(any.([]*string)))
	case []*bool:
		return Bools(key, boolPtrs2Bools(any.([]*bool)))

	case *[]int:
		return Ints(key, *(any.(*[]int)))
	case *[]int8:
		return Int8s(key, *(any.(*[]int8)))
	case *[]int16:
		return Int16s(key, *(any.(*[]int16)))
	case *[]int32:
		return Int32s(key, *(any.(*[]int32)))
	case *[]int64:
		return Int64s(key, *(any.(*[]int64)))
	case *[]uint:
		return Uints(key, *(any.(*[]uint)))
	case *[]uint8:
		return Uint8s(key, *(any.(*[]uint8)))
	case *[]uint16:
		return Uint16s(key, *(any.(*[]uint16)))
	case *[]uint32:
		return Uint32s(key, *(any.(*[]uint32)))
	case *[]uint64:
		return Uint64s(key, *(any.(*[]uint64)))
	case *[]float32:
		return Float32s(key, *(any.(*[]float32)))
	case *[]float64:
		return Float64s(key, *(any.(*[]float64)))
	case *[]string:
		return Strings(key, *(any.(*[]string)))
	case *[]bool:
		return Bools(key, *(any.(*[]bool)))
	default:
		return nullImeta(key)
	}
}

func intPtrs2ints(in []*int) []int {
	out := make([]int, 0, len(in))
	for i := range in {
		out = append(out, *in[i])
	}
	return out
}

func int8Ptrs2int8s(in []*int8) []int8 {
	out := make([]int8, 0, len(in))
	for i := range in {
		out = append(out, *in[i])
	}
	return out
}

func int16Ptrs2int16s(in []*int16) []int16 {
	out := make([]int16, 0, len(in))
	for i := range in {
		out = append(out, *(in[i]))
	}
	return out
}

func int32Ptrs2int32s(in []*int32) []int32 {
	out := make([]int32, 0, len(in))
	for i := range in {
		out = append(out, *(in[i]))
	}
	return out
}

func int64Ptrs2int64s(in []*int64) []int64 {
	out := make([]int64, 0, len(in))
	for i := range in {
		out = append(out, *(in[i]))
	}
	return out
}

func uintPtrs2uints(in []*uint) []uint {
	out := make([]uint, 0, len(in))
	for i := range in {
		out = append(out, *in[i])
	}
	return out
}

func uint8Ptrs2uint8s(in []*uint8) []uint8 {
	out := make([]uint8, 0, len(in))
	for i := range in {
		out = append(out, *in[i])
	}
	return out
}

func uint16Ptrs2uint16s(in []*uint16) []uint16 {
	out := make([]uint16, 0, len(in))
	for i := range in {
		out = append(out, *(in[i]))
	}
	return out
}

func uint32Ptrs2uint32s(in []*uint32) []uint32 {
	out := make([]uint32, 0, len(in))
	for i := range in {
		out = append(out, *(in[i]))
	}
	return out
}

func uint64Ptrs2uint64s(in []*uint64) []uint64 {
	out := make([]uint64, 0, len(in))
	for i := range in {
		out = append(out, *(in[i]))
	}
	return out
}

func float64Ptrs2Float64s(in []*float64) []float64 {
	out := make([]float64, 0, len(in))
	for i := range in {
		out = append(out, *(in[i]))
	}
	return out
}

func float32Ptrs2Float32s(in []*float32) []float32 {
	out := make([]float32, 0, len(in))
	for i := range in {
		out = append(out, *(in[i]))
	}
	return out
}

func stringPtrs2Strings(in []*string) []string {
	out := make([]string, 0, len(in))
	for i := range in {
		out = append(out, *(in[i]))
	}
	return out
}

func boolPtrs2Bools(in []*bool) []bool {
	out := make([]bool, 0, len(in))
	for i := range in {
		out = append(out, *(in[i]))
	}
	return out
}

func nullImeta(key string) Meta {
	return imeta{
		key:      toBytes(key),
		value:    []byte(null),
		needWrap: false,
	}
}

func emptyArrayImeta(key string) Meta {
	return imeta{
		key:      toBytes(key),
		value:    []byte{'[', ']'},
		needWrap: false,
	}
}
