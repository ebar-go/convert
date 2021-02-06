package convert

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestString_ToInt(t *testing.T) {
	s := String("1")

	assert.Equal(t, 1, s.ToInt())

}

func TestString_ToByte(t *testing.T) {
	s := String("1")
	assert.Equal(t, []byte("1"), s.ToByte())
}


func TestString_ToFloat(t *testing.T) {
	s := String("1.2")
	assert.Equal(t, 1.2, s.ToFloat())
}


func TestString_ToReader(t *testing.T) {
	s := String("hello")
	assert.Equal(t, strings.NewReader("hello"), s.ToReader())

}


func BenchmarkStringValue_ToInt(b *testing.B) {
	s := String("1")
	for i := 0; i < b.N; i++ {
		_ = s.ToInt()
	}
}


func BenchmarkStringValue_ToFloat(b *testing.B) {
	s := String("1")
	for i := 0; i < b.N; i++ {
		_ = s.ToFloat()
	}
}


func BenchmarkStringValue_ToByte(b *testing.B) {
	s := String("1")
	for i := 0; i < b.N; i++ {
		_ = s.ToByte()
	}
}