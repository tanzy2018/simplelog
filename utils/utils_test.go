package utils

import (
	"fmt"
	"testing"
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
	fmt.Println(1596788792192 - 1596788790100)
}
