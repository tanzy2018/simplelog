package internal_test

import (
	"fmt"
	"strings"
	"testing"

	. "github.com/tanzy2018/simplelog/internal"
)

func TestToString(t *testing.T) {
	testData := []struct {
		name     string
		expected string
		data     []byte
	}{
		{"accii", "012abc", []byte("012abc")},
		{"中文", "中文abc", []byte("中文abc")},

		{"日本語", "こんにちは", []byte("こんにちは")},
	}

	for _, td := range testData {
		if actual := ToString(td.data); td.expected != actual {
			t.Logf("\nname:%s,\nexpected:%s,\nactual:%s\n", td.name, td.expected, actual)
		}
	}
}

func TestDemo(t *testing.T) {
	s := strings.Join([]string{"", "", ""}, "/")
	fmt.Println(s)
}

func TestToBytes(t *testing.T) {
	rs := RandomString(10240)
	testData := []struct {
		name     string
		data     string
		expected []byte
	}{
		{"accii", "012abc", []byte("012abc")},
		{"中文", "中文abc", []byte("中文abc")},

		{"日本語", "こんにちは", []byte("こんにちは")},
		{"bigstr", rs, []byte(rs)},
	}

	for _, td := range testData {
		actual := ToBytes(td.data)
		if string(td.expected) != string(actual) {
			t.Logf("\nname:%s,\nexpected:%s,\nactual:%s\n", td.name, td.expected, actual)
		}
	}
}

var testStr = RandomString(10240)

func BenchmarkToBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ToBytes(testStr)
	}
}

func TestCallStack(t *testing.T) {
	fmt.Println("stacks:\n", call())
}

func call() (s string) {
	defer func() {
		if err := recover(); err != nil {
			s = fmt.Sprintf("%v:\n%s", err, CallStack(2))
		}

	}()
	a := make([]byte, 0)
	_ = a[1]
	return ""
}
