package convert

import (
	"io"
	"strconv"
	"strings"
)

type String string

// 转int
func (s String) ToInt() int {
	return stringToInt(string(s))
}

// 转float
func (s String) ToFloat() float64 {
	result, _ := strconv.ParseFloat(string(s), 64)
	return result
}
// 转byte
func (s String) ToByte() []byte {
	return stringToByte(string(s))
}

// 转reader
func (s String) ToReader() io.Reader {
	return strings.NewReader(string(s))
}
