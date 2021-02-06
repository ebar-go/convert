package convert

import (
	"io"
	"strconv"
	"strings"
	"unsafe"
)

type StringValue struct {
	val string
}

func String(val string) StringValue {
	return StringValue{val}
}

func (s StringValue) String() string {
	return s.val
}

// 转int
func (s StringValue) ToInt() int {
	n, _ := strconv.Atoi(s.val)
	return n
}

// 转float
func (s StringValue) ToFloat() float64 {
	result, _ := strconv.ParseFloat(s.val, 64)
	return result
}
// 转byte
func (s StringValue) ToByte() []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s.val))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// 转reader
func (s StringValue) ToReader() io.Reader {
	return strings.NewReader(s.val)
}
