package convert

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringValue(t *testing.T) {
	s := String("1")
	assert.Equal(t, "1", s.String())
	assert.Equal(t, []byte("1"), s.ToByte())
	assert.Equal(t, 1, s.ToInt())
	assert.Equal(t, float64(1), s.ToFloat())
}

func BenchmarkStringValue_String(b *testing.B) {
	s := String("1")
	for i := 0; i < b.N; i++ {
		_ = s.String()
	}
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