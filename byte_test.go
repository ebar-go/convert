package convert

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestByte_ToString(t *testing.T) {
	b := Byte([]byte{'h','e'})
	assert.Equal(t, "he", b.ToString())


}

func TestByte_ToInt(t *testing.T) {
	n := Byte([]byte{0, 0, 0, 1})
	assert.Equal(t, 1, n.ToInt())
}

func TestByte_ToReader(t *testing.T) {
	b := Byte([]byte("hello world"))
	assert.Equal(t, bytes.NewReader([]byte("hello world")), b.ToReader())
}

func BenchmarkByte_ToString(b *testing.B) {
	by := Byte([]byte{'h','e'})
	for i := 0; i < b.N; i++ {
		_ = by.ToString()
	}
}


func BenchmarkByte_ToInt(b *testing.B) {
	n := Byte([]byte{0, 0, 0, 1})
	for i := 0; i < b.N; i++ {
		_ = n.ToInt()
	}
}


func BenchmarkByte_ToReader(b *testing.B) {
	by := Byte([]byte{'h','e'})
	for i := 0; i < b.N; i++ {
		_ = by.ToReader()
	}
}