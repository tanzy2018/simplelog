package encode

import (
	"reflect"
	"testing"
)

func TestInt(t *testing.T) {
	type args struct {
		key string
		val int
	}

	tests := []struct {
		name string
		args args
		want Meta
	}{
		{"Int", args{"key0", 12}, imeta{[]byte("key0"), []byte("12"), false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int(tt.args.key, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8(t *testing.T) {
	type args struct {
		key string
		val int8
	}

	tests := []struct {
		name string
		args args
		want Meta
	}{
		{"Int8", args{"key0", 12}, imeta{[]byte("key0"), []byte("12"), false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int8(tt.args.key, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16(t *testing.T) {
	type args struct {
		key string
		val int16
	}

	tests := []struct {
		name string
		args args
		want Meta
	}{
		{"Int16", args{"key0", 12}, imeta{[]byte("key0"), []byte("12"), false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int16(tt.args.key, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32(t *testing.T) {
	type args struct {
		key string
		val int32
	}

	tests := []struct {
		name string
		args args
		want Meta
	}{
		{"Int32", args{"key0", 12}, imeta{[]byte("key0"), []byte("12"), false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int32(tt.args.key, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64(t *testing.T) {
	type args struct {
		key string
		val int64
	}

	tests := []struct {
		name string
		args args
		want Meta
	}{
		{"Int64", args{"key0", 12}, imeta{[]byte("key0"), []byte("12"), false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int64(tt.args.key, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUInt(t *testing.T) {
	type args struct {
		key string
		val uint
	}

	tests := []struct {
		name string
		args args
		want Meta
	}{
		{"Uint", args{"key0", 12}, imeta{[]byte("key0"), []byte("12"), false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint(tt.args.key, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint8(t *testing.T) {
	type args struct {
		key string
		val uint8
	}

	tests := []struct {
		name string
		args args
		want Meta
	}{
		{"Uint8", args{"key0", 12}, imeta{[]byte("key0"), []byte("12"), false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint8(tt.args.key, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16(t *testing.T) {
	type args struct {
		key string
		val uint16
	}

	tests := []struct {
		name string
		args args
		want Meta
	}{
		{"Uint16", args{"key0", 12}, imeta{[]byte("key0"), []byte("12"), false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint16(tt.args.key, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32(t *testing.T) {
	type args struct {
		key string
		val int32
	}

	tests := []struct {
		name string
		args args
		want Meta
	}{
		{"Uint32", args{"key0", 12}, imeta{[]byte("key0"), []byte("12"), false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int32(tt.args.key, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64(t *testing.T) {
	type args struct {
		key string
		val uint64
	}

	tests := []struct {
		name string
		args args
		want Meta
	}{
		{"Uint64", args{"key0", 12}, imeta{[]byte("key0"), []byte("12"), false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint64(tt.args.key, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64(t *testing.T) {
	type args struct {
		key string
		val float64
	}

	tests := []struct {
		name string
		args args
		want Meta
	}{
		{"Float64", args{"key0", 12.01}, imeta{[]byte("key0"), []byte("12.01"), false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float64(tt.args.key, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Float64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32(t *testing.T) {
	type args struct {
		key string
		val float32
	}

	tests := []struct {
		name string
		args args
		want Meta
	}{
		{"Float64", args{"key0", 12.01}, imeta{[]byte("key0"), []byte("12.01"), false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float32(tt.args.key, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Float32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString(t *testing.T) {
	type args struct {
		key string
		val string
	}

	tests := []struct {
		name string
		args args
		want Meta
	}{
		{"String", args{"key0", "str0"}, imeta{[]byte("key0"), []byte("str0"), true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := String(tt.args.key, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBool(t *testing.T) {
	type args struct {
		key string
		b   bool
	}
	tests := []struct {
		name string
		args args
		want Meta
	}{
		{"Bool-false", args{"key0", false}, imeta{[]byte("key0"), []byte("false"), false}},
		{"Bool-true", args{"key1", true}, imeta{[]byte("key1"), []byte("true"), false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bool(tt.args.key, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInts(t *testing.T) {
	type args struct {
		key  string
		ints []int
	}
	tests := []struct {
		name string
		args args
		want Meta
	}{
		{"Ints-nil", args{"nil", nil}, imeta{[]byte("nil"), []byte("[]"), false}},
		{"Ints-empty", args{"empty", []int{}}, imeta{[]byte("empty"), []byte("[]"), false}},
		{"Ints-one", args{"one", []int{1}}, imeta{[]byte("one"), []byte("[1]"), false}},
		{"Ints-three", args{"three", []int{1, 2, 3}}, imeta{[]byte("three"), []byte("[1,2,3]"), false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ints(tt.args.key, tt.args.ints); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64s(t *testing.T) {
	type args struct {
		key  string
		ints []int64
	}
	tests := []struct {
		name string
		args args
		want Meta
	}{
		{"Int64s-empty", args{"empty", []int64{}}, imeta{[]byte("empty"), []byte("[]"), false}},
		{"Int64s-one", args{"one", []int64{1}}, imeta{[]byte("one"), []byte("[1]"), false}},
		{"Int64s-three", args{"three", []int64{1, 2, 3}}, imeta{[]byte("three"), []byte("[1,2,3]"), false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int64s(tt.args.key, tt.args.ints); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int64s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8s(t *testing.T) {
	type args struct {
		key  string
		ints []int8
	}
	tests := []struct {
		name string
		args args
		want Meta
	}{
		{"Int8s-nil", args{"nil", nil}, imeta{[]byte("nil"), []byte("[]"), false}},
		{"Int8s-empty", args{"empty", []int8{}}, imeta{[]byte("empty"), []byte("[]"), false}},
		{"Int8s-one", args{"one", []int8{1}}, imeta{[]byte("one"), []byte("[1]"), false}},
		{"Int8s-three", args{"three", []int8{1, 2, 3}}, imeta{[]byte("three"), []byte("[1,2,3]"), false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int8s(tt.args.key, tt.args.ints); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int8s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16s(t *testing.T) {
	type args struct {
		key  string
		ints []int16
	}
	tests := []struct {
		name string
		args args
		want Meta
	}{
		{"Int16s-nil", args{"nil", nil}, imeta{[]byte("nil"), []byte("[]"), false}},
		{"Int16s-empty", args{"empty", []int16{}}, imeta{[]byte("empty"), []byte("[]"), false}},
		{"Int16s-one", args{"one", []int16{1}}, imeta{[]byte("one"), []byte("[1]"), false}},
		{"Int16s-three", args{"three", []int16{1, 2, 3}}, imeta{[]byte("three"), []byte("[1,2,3]"), false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int16s(tt.args.key, tt.args.ints); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int16s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32s(t *testing.T) {
	type args struct {
		key  string
		ints []int32
	}
	tests := []struct {
		name string
		args args
		want Meta
	}{
		{"Int32s-nil", args{"nil", nil}, imeta{[]byte("nil"), []byte("[]"), false}},
		{"Int32s-empty", args{"empty", []int32{}}, imeta{[]byte("empty"), []byte("[]"), false}},
		{"Int32s-one", args{"one", []int32{1}}, imeta{[]byte("one"), []byte("[1]"), false}},
		{"Int32s-three", args{"three", []int32{1, 2, 3}}, imeta{[]byte("three"), []byte("[1,2,3]"), false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int32s(tt.args.key, tt.args.ints); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int32s() = %v, want %v", got, tt.want)
			}
		})
	}
}
