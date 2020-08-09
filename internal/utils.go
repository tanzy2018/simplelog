package internal

import (
	"math/rand"
	"runtime"
	"strconv"
	"strings"
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
	if len(b) == 0 {
		return ""
	}
	return *(*string)(unsafe.Pointer(&b))
}

// ToBytes ...
func ToBytes(s string) []byte {
	if len(s) == 0 {
		return nil
	}
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

// CallStack ...
func CallStack(skip int) string {
	if skip < 1 {
		skip = 1
	}
	firstFrame, depth := skip, 20+skip
	var (
		pc   uintptr
		file string
		line int
		ok   bool
	)
	ok = true
	callers := make([]byte, 0, 1024)
	for {
		pc, file, line, ok = runtime.Caller(skip)
		if !ok {
			return ToString(callers[:len(callers)])
		}
		if strings.Contains(file, "/src/runtime/") {
			skip++
			firstFrame++
			continue
		}
		f := runtime.FuncForPC(pc)
		if skip > firstFrame {
			callers = append(callers, []byte("->")...)
		}
		callers = append(callers, ToBytes(file)...)
		callers = append(callers, ':')
		callers = append(callers, ToBytes(strconv.FormatInt(int64(line), 10))...)
		callers = append(callers, '.')
		callers = append(callers, '(')
		callers = append(callers, ToBytes(f.Name())...)
		callers = append(callers, ')')
		skip++
		if skip > depth {
			return ToString(callers[:len(callers)])
		}
	}
}
