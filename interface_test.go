package convert

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInterface_Int(t *testing.T) {
	var num int = 1
	assert.Equal(t, Interface("1").ToInt(), num)
	assert.Equal(t, Interface(1).ToInt(), num)
	assert.Equal(t, Interface(1.0).ToInt(), num)
	assert.Equal(t, Interface(uint(1)).ToInt(), num)
	assert.Equal(t, Interface(true).ToInt(), num)
}

func BenchmarkInterface_Int(b *testing.B) {
	var a int64 = 1
	inter := Interface(a)
	for i := 0; i < b.N; i++ {
		inter.ToInt()
	}
}

func TestInterface_String(t *testing.T) {
	var str = "1"
	assert.Equal(t, str, Interface(1).ToString())
	assert.Equal(t, "1.12", Interface(1.12).ToString())
	assert.Equal(t, str, Interface(uint64(1)).ToString())
	assert.Equal(t, "true", Interface(true).ToString())
	assert.Equal(t, "false", Interface(false).ToString())
}

func BenchmarkInterface_String(b *testing.B) {
	inter := Interface(1)
	for i := 0; i < b.N; i++ {
		_ = inter.ToString()
	}
}

func TestInterface_Float64(t *testing.T) {
	assert.Equal(t, 1.1, Interface(1.1).ToFloat())
	assert.Equal(t, 1.2, Interface("1.2").ToFloat())
	assert.Equal(t, float64(1), Interface(int32(1)).ToFloat())
	assert.Equal(t, float64(1), Interface(uint(1)).ToFloat())
}

func BenchmarkInterface_Float64(b *testing.B) {
	inter := Interface(1)
	for i := 0; i < b.N; i++ {
		inter.ToFloat()
	}
}
