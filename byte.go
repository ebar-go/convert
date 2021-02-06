package convert

import (
	"bytes"
	"encoding/binary"
	"io"
)
type Byte []byte


func (b Byte) ToString() string {
	return byteToString(b)
}

func (b Byte) ToReader() io.Reader {
	return bytes.NewReader(b)
}

func (b Byte) ToInt() int {
	var x int32
	_ = binary.Read(b.ToReader(), binary.BigEndian, &x)
	return int(x)
}