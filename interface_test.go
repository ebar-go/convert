package convert

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInterface_Int(t *testing.T) {
	var num int = 1
	assert.Equal(t, NewInterface("1").Int(), num)
	assert.Equal(t, NewInterface(1).Int(), num)
	assert.Equal(t, NewInterface(1.0).Int(), num)
	assert.Equal(t, NewInterface(uint(1)).Int(), num)
	assert.Equal(t, NewInterface(true).Int(), num)
}

func BenchmarkInterface_Int(b *testing.B) {
	var a int64 = 1
	inter := NewInterface(a)
	for i := 0; i < b.N; i++ {
		inter.Int()
	}
}

func TestInterface_String(t *testing.T) {
	var str = "1"
	assert.Equal(t, str, NewInterface(1).String())
	assert.Equal(t, "1.12", NewInterface(1.12).String())
	assert.Equal(t, str, NewInterface(uint64(1)).String())
	assert.Equal(t, "true", NewInterface(true).String())
	assert.Equal(t, "false", NewInterface(false).String())
}

func BenchmarkInterface_String(b *testing.B) {
	inter := NewInterface(1)
	for i := 0; i < b.N; i++ {
		_ = inter.String()
	}
}

func TestInterface_Float64(t *testing.T) {
	assert.Equal(t, 1.1, NewInterface(1.1).Float64())
	assert.Equal(t, 1.2, NewInterface("1.2").Float64())
	assert.Equal(t, float64(1), NewInterface(int32(1)).Float64())
	assert.Equal(t, float64(1), NewInterface(uint(1)).Float64())
}

func BenchmarkInterface_Float64(b *testing.B) {
	inter := NewInterface(1)
	for i := 0; i < b.N; i++ {
		inter.Float64()
	}
}