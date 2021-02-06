package convert

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntArray(t *testing.T) {
	arr := IntArray([]int{1,2,3})
	assert.Equal(t, []int{1,2,3}, arr.ToInt())
	assert.Equal(t, []string{"1","2", "3"}, arr.ToString())
	assert.Equal(t, []interface{}{1,2,3}, arr.ToInterface())
}


func TestStringArray(t *testing.T) {
	arr := StringArray([]string{"1","2", "3"})
	assert.Equal(t, []int{1,2,3}, arr.ToInt())
	assert.Equal(t, []string{"1","2", "3"}, arr.ToString())
	assert.Equal(t, []interface{}{"1","2", "3"}, arr.ToInterface())
}

func BenchmarkIntArray_ToInt(b *testing.B) {
	arr := IntArray([]int{1,2,3})
	for i := 0; i < b.N; i++ {
		arr.ToInt()
	}
}

func BenchmarkIntArray_ToString(b *testing.B) {
	arr := IntArray([]int{1,2,3})
	for i := 0; i < b.N; i++ {
		arr.ToString()
	}
}

func BenchmarkIntArray_ToInterface(b *testing.B) {
	arr := IntArray([]int{1,2,3})
	for i := 0; i < b.N; i++ {
		arr.ToInterface()
	}
}

func BenchmarkStringArray_ToInt(b *testing.B) {
	arr := StringArray([]string{"1","2", "3"})
	for i := 0; i < b.N; i++ {
		arr.ToInt()
	}
}


func BenchmarkStringArray_ToString(b *testing.B) {
	arr := StringArray([]string{"1","2", "3"})
	for i := 0; i < b.N; i++ {
		arr.ToString()
	}
}


func BenchmarkStringArray_ToInterface(b *testing.B) {
	arr := StringArray([]string{"1","2", "3"})
	for i := 0; i < b.N; i++ {
		arr.ToInterface()
	}
}