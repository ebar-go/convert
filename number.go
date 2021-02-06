package convert

import (
	"bytes"
	"encoding/binary"
	"strconv"
)

type Number int64

// 转字符串
func (n Number) ToString() string {
	return strconv.FormatInt(int64(n), 10)
}

// 转float
func (n Number) ToFloat() float64 {
	return float64(n)
}

func (n Number) ToByte() []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	_ = binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

