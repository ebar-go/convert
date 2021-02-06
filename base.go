package convert

import (
	"strconv"
	"unsafe"
)

func intToString(i int) string {
	return strconv.Itoa(i)
}

func stringToInt(s string) int  {
	n, _ := strconv.Atoi(s)
	return n
}

func stringToByte(s string) []byte  {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}