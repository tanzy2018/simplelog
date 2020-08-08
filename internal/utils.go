package internal

import (
	"math/rand"
	"time"
	"unsafe"
)

const deFaultTimeFormat = "2006-01-02 15:04:05"

const utilsRandStr = "23456789abcdefghijkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"

var utilsRand = rand.New(rand.NewSource(time.Now().UnixNano()))

// TimeFormat ...
func TimeFormat(format string) string {
	if len(format) == 0 {
		format = deFaultTimeFormat
	}
	return time.Now().Format(format)
}

// ToString ...
func ToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// ToBytes ...
func ToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

// RandInt ...
func RandInt(n int) int {
	if n <= 1 {
		return 0
	}
	return utilsRand.Intn(n)
}

// RandomString ...
func RandomString(n int) string {
	out := make([]byte, 0, n)
	for l := len(utilsRandStr); n > l; {
		out = append(out, randomString(l)...)
		n -= len(utilsRandStr)
	}
	out = append(out, randomString(n)...)
	return ToString(out)
}

func randomString(n int) []byte {
	if n <= 0 {
		return nil
	}
	tpl := []byte(utilsRandStr)

	for i := 0; i < len(tpl)/2; i++ {
		idx := utilsRand.Intn(len(tpl) - i)
		tpl[idx], tpl[len(tpl)-i-1] = tpl[len(tpl)-i-1], tpl[idx]
	}
	if n > len(tpl) {
		return tpl[:]
	}

	return tpl[:n]
}
