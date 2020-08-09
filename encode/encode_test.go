package encode

import (
	"reflect"
	"testing"
)

func TestAny(t *testing.T) {
	type anyMap map[string]interface{}
	type anyStruct struct {
		Name string
		Age  int
		Flag bool
	}

	type args struct {
		key string
		any interface{}
	}

	var (
		i0  int   = 0
		i8  int8  = 8
		i16 int16 = 16
		i32 int32 = 32
		i64 int64 = 64

		u0  uint   = 0
		u8  uint8  = 8
		u16 uint16 = 16
		u32 uint32 = 32
		u64 uint64 = 64

		f32 float32 = 32.0
		f64 float64 = 64.0

		str  string = "str"
		str2 string = "str2"
		bo   bool   = false
		b1   bool   = true

		// *i8 = 8
		// *i16 = 16
		// *i32 = 32
		// *i64 = 64

		// *u0 = 0
		// *u8 = 8
		// *u16 = 16
		// *u32 = 32
		// *u64 = 64
	)

	tests := []struct {
		name string
		args args
		want Meta
	}{
		// int*
		{"Int_0", args{"int_0", int(1)}, imeta{[]byte("int_0"), []byte("1"), false}},
		{"Int8_0", args{"int8_0", int8(1)}, imeta{[]byte("int8_0"), []byte("1"), false}},
		{"Int16_0", args{"int16_0", int16(1)}, imeta{[]byte("int16_0"), []byte("1"), false}},
		{"Int32_0", args{"int32_0", int32(1)}, imeta{[]byte("int32_0"), []byte("1"), false}},
		{"Int64_0", args{"int64_0", int64(1)}, imeta{[]byte("int64_0"), []byte("1"), false}},
		{"Int_1", args{"int_1", int(-1)}, imeta{[]byte("int_1"), []byte("-1"), false}},
		{"Int8_1", args{"int8_1", int8(-1)}, imeta{[]byte("int8_1"), []byte("-1"), false}},
		{"Int16_1", args{"int16_1", int16(-1)}, imeta{[]byte("int16_1"), []byte("-1"), false}},
		{"Int32_1", args{"int32_1", int32(-1)}, imeta{[]byte("int32_1"), []byte("-1"), false}},
		{"Int64_1", args{"int64_1", int64(-1)}, imeta{[]byte("int64_1"), []byte("-1"), false}},

		// *int*
		{"*Int_0", args{"int_0", &i0}, imeta{[]byte("int_0"), []byte("0"), false}},
		{"*Int8_0", args{"int8_0", &i8}, imeta{[]byte("int8_0"), []byte("8"), false}},
		{"*Int16_0", args{"int16_0", &i16}, imeta{[]byte("int16_0"), []byte("16"), false}},
		{"*Int32_0", args{"int32_0", &i32}, imeta{[]byte("int32_0"), []byte("32"), false}},
		{"*Int64_0", args{"int64_0", &i64}, imeta{[]byte("int64_0"), []byte("64"), false}},

		// uint*
		{"Uint_0", args{"uint_0", u0}, imeta{[]byte("uint_0"), []byte("0"), false}},
		{"Uint8_0", args{"uint8_0", u8}, imeta{[]byte("uint8_0"), []byte("8"), false}},
		{"Uint16_0", args{"uint16_0", u16}, imeta{[]byte("uint16_0"), []byte("16"), false}},
		{"Uint32_0", args{"uint32_0", u32}, imeta{[]byte("uint32_0"), []byte("32"), false}},
		{"Uint64_0", args{"uint64_0", u64}, imeta{[]byte("uint64_0"), []byte("64"), false}},

		// *uint*
		{"*Uint_0", args{"uint_0", &u0}, imeta{[]byte("uint_0"), []byte("0"), false}},
		{"*Uint8_0", args{"uint8_0", &u8}, imeta{[]byte("uint8_0"), []byte("8"), false}},
		{"*Uint16_0", args{"uint16_0", &u16}, imeta{[]byte("uint16_0"), []byte("16"), false}},
		{"*Uint32_0", args{"uint32_0", &u32}, imeta{[]byte("uint32_0"), []byte("32"), false}},
		{"*Uint64_0", args{"uint64_0", &u64}, imeta{[]byte("uint64_0"), []byte("64"), false}},

		// float*
		{"Float32_0", args{"float32_0", float32(1.0)}, imeta{[]byte("float32_0"), []byte("1"), false}},
		{"Float64_0", args{"float64_0", float64(1.0)}, imeta{[]byte("float64_0"), []byte("1"), false}},
		{"Float32_1", args{"float32_1", float32(1.01)}, imeta{[]byte("float32_1"), []byte("1.01"), false}},
		{"Float64_1", args{"float64_1", float64(1.01)}, imeta{[]byte("float64_1"), []byte("1.01"), false}},

		// *float*
		{"*Float32_1", args{"float32_1", &f32}, imeta{[]byte("float32_1"), []byte("32"), false}},
		{"*Float64_1", args{"float64_1", &f64}, imeta{[]byte("float64_1"), []byte("64"), false}},

		// bool
		{"Bool_true", args{"bool_true", true}, imeta{[]byte("bool_true"), []byte("true"), false}},
		{"Bool_false", args{"bool_false", false}, imeta{[]byte("bool_false"), []byte("false"), false}},

		// *bool
		{"*Bool_false", args{"bool_false", &bo}, imeta{[]byte("bool_false"), []byte("false"), false}},
		// string
		{"String_0", args{"string_0", " "}, imeta{[]byte("string_0"), []byte(" "), true}},
		{"String_1", args{"string_1", "abc"}, imeta{[]byte("string_1"), []byte("abc"), true}},
		{"String_2", args{"string_2", "中文"}, imeta{[]byte("string_2"), []byte("中文"), true}},
		{"String_3", args{"string_3", ""}, imeta{[]byte("string_3"), nil, true}},

		// *string
		{"*String_0", args{"string_0", &str}, imeta{[]byte("string_0"), []byte("str"), true}},

		// object map
		{"Object_map_0", args{"object_map_0", map[string]interface{}{"name": "tanzy"}}, imeta{[]byte("object_map_0"), []byte(`{"name":"tanzy"}`), false}},
		{"Object_map_1", args{"object_map_1", map[int]interface{}{1: 18}}, imeta{[]byte("object_map_1"), []byte(`{"1":18}`), false}},
		{"Object_map_2", args{"object_map_2", map[int]interface{}{}}, imeta{[]byte("object_map_2"), []byte("{}"), false}},
		{"Object_map_3", args{"object_map_3", (map[string]interface{})(nil)}, imeta{[]byte("object_map_3"), []byte("null"), false}},
		{"Object_map_4", args{"object_map_4", map[string]int{"age": 18}}, imeta{[]byte("object_map_4"), []byte(`{"age":18}`), false}},
		{"Object_map_5", args{"object_map_5", map[string]map[int]string{"first": {1: "A"}}}, imeta{[]byte("object_map_5"), []byte(`{"first":{"1":"A"}}`), false}},
		// object struct
		{"Object_struct_0", args{"object_struct_0", anyStruct{}}, imeta{[]byte("object_struct_0"), []byte(`{"Name":"","Age":0,"Flag":false}`), false}},
		{"Object_struct_1", args{"object_struct_1", anyStruct{"tanzy", 18, true}}, imeta{[]byte("object_struct_1"), []byte(`{"Name":"tanzy","Age":18,"Flag":true}`), false}},

		// object *map
		{"Object_*map_0", args{"object_*map_0", &map[string]interface{}{"name": "tanzy"}}, imeta{[]byte("object_*map_0"), []byte(`{"name":"tanzy"}`), false}},
		{"Object_*map_1", args{"object_*map_1", &map[int]interface{}{1: 18}}, imeta{[]byte("object_*map_1"), []byte(`{"1":18}`), false}},
		{"Object_*map_2", args{"object_*map_2", &map[int]interface{}{}}, imeta{[]byte("object_*map_2"), []byte("{}"), false}},
		{"Object_*map_3", args{"object_*map_3", (*map[string]interface{})(nil)}, imeta{[]byte("object_*map_3"), []byte("null"), false}},

		// object *struct
		{"Object_*struct_0", args{"object_*struct_0", &anyStruct{}}, imeta{[]byte("object_*struct_0"), []byte(`{"Name":"","Age":0,"Flag":false}`), false}},
		{"Object_*struct_1", args{"object_*struct_1", &anyStruct{"tanzy", 18, true}}, imeta{[]byte("object_*struct_1"), []byte(`{"Name":"tanzy","Age":18,"Flag":true}`), false}},
		{"Object_*struct_2", args{"object_*struct_2", (*anyStruct)(nil)}, imeta{[]byte("object_*struct_2"), []byte("null"), false}},

		// ints*
		{"Ints_0", args{"ints_0", []int{}}, imeta{[]byte("ints_0"), []byte("[]"), false}},
		{"Ints_1", args{"ints_1", ([]int)(nil)}, imeta{[]byte("ints_1"), []byte("[]"), false}},
		{"Ints_2", args{"ints_2", []int{1}}, imeta{[]byte("ints_2"), []byte("[1]"), false}},
		{"Ints_3", args{"ints_3", []int{1, -2, 3}}, imeta{[]byte("ints_3"), []byte("[1,-2,3]"), false}},
		{"Ints_4", args{"ints_4", [3]int{1, -2, 3}}, imeta{[]byte("ints_4"), []byte("[1,-2,3]"), false}},

		{"Int8s_0", args{"int8s_0", []int8{}}, imeta{[]byte("int8s_0"), []byte("[]"), false}},
		{"Int8s_1", args{"int8s_1", ([]int8)(nil)}, imeta{[]byte("int8s_1"), []byte("[]"), false}},
		{"Int8s_2", args{"int8s_2", []int8{1}}, imeta{[]byte("int8s_2"), []byte("[1]"), false}},
		{"Int8s_3", args{"int8s_3", []int8{1, -2, 3}}, imeta{[]byte("int8s_3"), []byte("[1,-2,3]"), false}},

		{"Int16s_0", args{"int16s_0", []int16{}}, imeta{[]byte("int16s_0"), []byte("[]"), false}},
		{"Int16s_1", args{"int16s_1", ([]int16)(nil)}, imeta{[]byte("int16s_1"), []byte("[]"), false}},
		{"Int16s_2", args{"int16s_2", []int16{1}}, imeta{[]byte("int16s_2"), []byte("[1]"), false}},
		{"Int16s_3", args{"int16s_3", []int16{1, -2, 3}}, imeta{[]byte("int16s_3"), []byte("[1,-2,3]"), false}},

		{"Int32s_0", args{"int32s_0", []int32{}}, imeta{[]byte("int32s_0"), []byte("[]"), false}},
		{"Int32s_1", args{"int32s_1", ([]int32)(nil)}, imeta{[]byte("int32s_1"), []byte("[]"), false}},
		{"Int32s_2", args{"int32s_2", []int32{1}}, imeta{[]byte("int32s_2"), []byte("[1]"), false}},
		{"Int32s_3", args{"int32s_3", []int32{1, -2, 3}}, imeta{[]byte("int32s_3"), []byte("[1,-2,3]"), false}},

		{"Int64s_0", args{"int64s_0", []int64{}}, imeta{[]byte("int64s_0"), []byte("[]"), false}},
		{"Int64s_1", args{"int64s_1", ([]int64)(nil)}, imeta{[]byte("int64s_1"), []byte("[]"), false}},
		{"Int64s_2", args{"int64s_2", []int64{1}}, imeta{[]byte("int64s_2"), []byte("[1]"), false}},
		{"Int64s_3", args{"int64s_3", []int64{1, -2, 3}}, imeta{[]byte("int64s_3"), []byte("[1,-2,3]"), false}},

		// uints*
		{"Uints_0", args{"uints_0", []uint{}}, imeta{[]byte("uints_0"), []byte("[]"), false}},
		{"Uints_1", args{"uints_1", ([]uint)(nil)}, imeta{[]byte("uints_1"), []byte("[]"), false}},
		{"Uints_2", args{"uints_2", []uint{1}}, imeta{[]byte("uints_2"), []byte("[1]"), false}},
		{"Uints_3", args{"uints_3", []uint{1, 2, 3}}, imeta{[]byte("uints_3"), []byte("[1,2,3]"), false}},

		{"Uint8s_0", args{"uint8s_0", []uint8{}}, imeta{[]byte("uint8s_0"), []byte("[]"), false}},
		{"Uint8s_1", args{"uint8s_1", ([]uint8)(nil)}, imeta{[]byte("uint8s_1"), []byte("[]"), false}},
		{"Uint8s_2", args{"uint8s_2", []uint8{1}}, imeta{[]byte("uint8s_2"), []byte("[1]"), false}},
		{"Uint8s_3", args{"uint8s_3", []uint8{1, 2, 3}}, imeta{[]byte("uint8s_3"), []byte("[1,2,3]"), false}},

		{"Uint16s_0", args{"uint16s_0", []uint16{}}, imeta{[]byte("uint16s_0"), []byte("[]"), false}},
		{"Uint16s_1", args{"uint16s_1", ([]uint16)(nil)}, imeta{[]byte("uint16s_1"), []byte("[]"), false}},
		{"Uint16s_2", args{"uint16s_2", []uint16{1}}, imeta{[]byte("uint16s_2"), []byte("[1]"), false}},
		{"Uint16s_3", args{"uint16s_3", []uint16{1, 2, 3}}, imeta{[]byte("uint16s_3"), []byte("[1,2,3]"), false}},

		{"Uint32s_0", args{"uint32s_0", []uint32{}}, imeta{[]byte("uint32s_0"), []byte("[]"), false}},
		{"Uint32s_1", args{"uint32s_1", ([]uint32)(nil)}, imeta{[]byte("uint32s_1"), []byte("[]"), false}},
		{"Uint32s_2", args{"uint32s_2", []uint32{1}}, imeta{[]byte("uint32s_2"), []byte("[1]"), false}},
		{"Uint32s_3", args{"uint32s_3", []uint32{1, 2, 3}}, imeta{[]byte("uint32s_3"), []byte("[1,2,3]"), false}},

		{"Uint64s_0", args{"uint64s_0", []uint64{}}, imeta{[]byte("uint64s_0"), []byte("[]"), false}},
		{"Uint64s_1", args{"uint64s_1", ([]uint64)(nil)}, imeta{[]byte("uint64s_1"), []byte("[]"), false}},
		{"Uint64s_2", args{"uint64s_2", []uint64{1}}, imeta{[]byte("uint64s_2"), []byte("[1]"), false}},
		{"Uint64s_3", args{"uint64s_3", []uint64{1, 2, 3}}, imeta{[]byte("uint64s_3"), []byte("[1,2,3]"), false}},

		// []string
		{"Strings_0", args{"string_0", ([]string)(nil)}, imeta{[]byte("string_0"), []byte("[]"), false}},
		{"Strings_1", args{"string_1", []string{}}, imeta{[]byte("string_1"), []byte("[]"), false}},
		{"Strings_2", args{"string_2", []string{"abc"}}, imeta{[]byte("string_2"), []byte(`["abc"]`), false}},
		{"Strings_3", args{"string_3", []string{"中文", "abc"}}, imeta{[]byte("string_3"), []byte(`["中文","abc"]`), false}},

		// *[]string
		{"*Strings_0", args{"string_0", (*[]string)(nil)}, imeta{[]byte("string_0"), []byte("null"), false}},
		{"*Strings_1", args{"string_1", &[]string{}}, imeta{[]byte("string_1"), []byte("[]"), false}},
		{"*Strings_2", args{"string_2", &[]string{"abc"}}, imeta{[]byte("string_2"), []byte(`["abc"]`), false}},
		{"*Strings_3", args{"string_3", &[]string{"中文", "abc"}}, imeta{[]byte("string_3"), []byte(`["中文","abc"]`), false}},

		// []*string
		{"Strings_*0", args{"string_0", ([]*string)(nil)}, imeta{[]byte("string_0"), []byte("[]"), false}},
		{"Strings_*1", args{"string_1", []*string{}}, imeta{[]byte("string_1"), []byte("[]"), false}},
		{"Strings_*2", args{"string_2", []*string{&str}}, imeta{[]byte("string_2"), []byte(`["str"]`), false}},
		{"Strings_*3", args{"string_3", []*string{&str, &str2}}, imeta{[]byte("string_3"), []byte(`["str","str2"]`), false}},

		// []bool
		{"Bools_nil", args{"bool_nil", ([]bool)(nil)}, imeta{[]byte("bool_nil"), []byte("[]"), false}},
		{"Bools_empty", args{"bool_empty", []bool{}}, imeta{[]byte("bool_empty"), []byte("[]"), false}},
		{"Bools_true", args{"bool_true", []bool{true}}, imeta{[]byte("bool_true"), []byte("[true]"), false}},
		{"Bools_false", args{"bool_false", []bool{false}}, imeta{[]byte("bool_false"), []byte("[false]"), false}},
		{"Bools_mixed", args{"bool_mixed", []bool{true, false, true, true}}, imeta{[]byte("bool_mixed"), []byte("[true,false,true,true]"), false}},

		// *[]bool
		{"*Bools_nil", args{"bool_nil", (*[]bool)(nil)}, imeta{[]byte("bool_nil"), []byte("null"), false}},
		{"*Bools_empty", args{"bool_empty", &[]bool{}}, imeta{[]byte("bool_empty"), []byte("[]"), false}},
		{"*Bools_true", args{"bool_true", &[]bool{true}}, imeta{[]byte("bool_true"), []byte("[true]"), false}},
		{"*Bools_false", args{"bool_false", &[]bool{false}}, imeta{[]byte("bool_false"), []byte("[false]"), false}},
		{"*Bools_mixed", args{"bool_mixed", &[]bool{true, false, true, true}}, imeta{[]byte("bool_mixed"), []byte("[true,false,true,true]"), false}},

		// []*bool
		{"Bools_*nil", args{"bool_nil", ([]*bool)(nil)}, imeta{[]byte("bool_nil"), []byte("[]"), false}},
		{"Bools_*empty", args{"bool_empty", []*bool{}}, imeta{[]byte("bool_empty"), []byte("[]"), false}},
		{"Bools_*true", args{"bool_true", []*bool{&b1}}, imeta{[]byte("bool_true"), []byte("[true]"), false}},
		{"Bools_*false", args{"bool_false", []*bool{&bo}}, imeta{[]byte("bool_false"), []byte("[false]"), false}},
		{"Bools_*mixed", args{"bool_mixed", []*bool{&b1, &bo, &b1, &b1}}, imeta{[]byte("bool_mixed"), []byte("[true,false,true,true]"), false}},

		// floats*
		{"Float32s_0", args{"float32s_0", []float32{}}, imeta{[]byte("float32s_0"), []byte("[]"), false}},
		{"Float32s_1", args{"float32s_1", ([]float32)(nil)}, imeta{[]byte("float32s_1"), []byte("[]"), false}},
		{"Float32s_2", args{"float32s_2", []float32{1}}, imeta{[]byte("float32s_2"), []byte("[1]"), false}},
		{"Float32s_3", args{"float32s_3", []float32{1, -2, 3, 4.0, 5.01, 0}}, imeta{[]byte("float32s_3"), []byte("[1,-2,3,4,5.01,0]"), false}},

		{"Float64s_0", args{"float64s_0", []float64{}}, imeta{[]byte("float64s_0"), []byte("[]"), false}},
		{"Float64s_1", args{"float64s_1", ([]float64)(nil)}, imeta{[]byte("float64s_1"), []byte("[]"), false}},
		{"Float64s_2", args{"float64s_2", []float64{1}}, imeta{[]byte("float64s_2"), []byte("[1]"), false}},
		{"Float64s_3", args{"float64s_3", []float64{1, -2, 3, 4.0, 5.01, 0}}, imeta{[]byte("float64s_3"), []byte("[1,-2,3,4,5.01,0]"), false}},

		// any []interface{}
		{"Any_nil", args{"any_nil", ([]interface{})(nil)}, imeta{[]byte("any_nil"), []byte("[]"), false}},
		{"Any_empty", args{"any_empty", []interface{}{}}, imeta{[]byte("any_empty"), []byte("[]"), false}},
		{"Any_map", args{"any_map", []interface{}{anyMap{"name": "tanzy"}}},
			imeta{[]byte("any_map"), []byte(`[{"name":"tanzy"}]`), false}},
		{"Any_*map", args{"any_*map", []interface{}{&anyMap{"name": "tanzy"}}},
			imeta{[]byte("any_*map"), []byte(`[{"name":"tanzy"}]`), false}},
		{"Any_struct", args{"any_struct", []interface{}{anyStruct{"tanzy", 18, true}}},
			imeta{[]byte("any_struct"), []byte(`[{"Name":"tanzy","Age":18,"Flag":true}]`), false}},
		{"Any_*struct", args{"any_*struct", []interface{}{&anyStruct{"tanzy", 18, true}}},
			imeta{[]byte("any_*struct"), []byte(`[{"Name":"tanzy","Age":18,"Flag":true}]`), false}},
		{"Any_mixed", args{"any_mixed",
			[]interface{}{nil, anyMap(nil), (*anyMap)(nil), (*anyStruct)(nil), anyMap{}, anyMap{"name": "map"},
				&anyMap{"name": "*map"}, anyStruct{}, &anyStruct{}, &anyStruct{"tanzy", 18, true}}},
			imeta{[]byte("any_mixed"),
				[]byte(`[null,null,null,null,{},{"name":"map"},{"name":"*map"},{"Name":"","Age":0,"Flag":false},{"Name":"","Age":0,"Flag":false},{"Name":"tanzy","Age":18,"Flag":true}]`), false}},
		{"Any_mixed2", args{"any_mixed2",
			[]interface{}{nil, 1, "str", false, anyMap{"name": "map"}, []int{1, 2}, []interface{}{2, "str2", anyStruct{"tanzy", 19, false}},
				anyStruct{}, &anyStruct{}, &anyStruct{"tanzy", 18, true}}},
			imeta{[]byte("any_mixed2"),
				[]byte(`[null,1,"str",false,{"name":"map"},[1,2],[2,"str2",{"Name":"tanzy","Age":19,"Flag":false}],{"Name":"","Age":0,"Flag":false},{"Name":"","Age":0,"Flag":false},{"Name":"tanzy","Age":18,"Flag":true}]`), false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Any(tt.args.key, tt.args.any); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Actual =%s,%s\n, want %s,%s", got.Key(), got.Value(), tt.want.Key(), tt.want.Value())
				t.Errorf("Actual = %v, want %v", got, tt.want)
			}
		})
	}
}
