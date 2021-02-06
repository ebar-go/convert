package convert

import (
	"io"
	"strconv"
	"strings"
	"unsafe"
)

type StringValue string

func String(val string) StringValue {
	return StringValue(val)
}

func (s StringValue) String() string {
	return string(s)
}

// 转int
func (s StringValue) ToInt() int {
	n, _ := strconv.Atoi(s.String())
	return n
}

// 转float
func (s StringValue) ToFloat() float64 {
	result, _ := strconv.ParseFloat(s.String(), 64)
	return result
}
// 转byte
func (s StringValue) ToByte() []byte {
	return stringToByte(string(s))
}

// 转reader
func (s StringValue) ToReader() io.Reader {
	return strings.NewReader(s.String())
}

func stringToByte(s string) []byte  {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}