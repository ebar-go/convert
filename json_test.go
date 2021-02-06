package convert

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJson_ToByte(t *testing.T) {
	j := Json(`{"name":"convert"}`)
	assert.Equal(t, []byte(`{"name":"convert"}`), j.ToByte())
}

func TestJson_ToMap(t *testing.T) {
	j := Json(`{"name":"convert"}`)
	result, err := j.ToMap()
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{"name":"convert"}, result)
}

func TestJson_ToStruct(t *testing.T) {
	j := Json(`{"name":"convert"}`)

	p := new(struct{
		Name string
	})
	assert.Nil(t, j.ToStruct(p))
	assert.Equal(t, "convert", p.Name)
}

func BenchmarkJson_ToByte(b *testing.B) {
	j := Json(`{"name":"convert"}`)
	for i := 0; i < b.N; i++ {
		j.ToByte()
	}
}

func BenchmarkJson_ToStruct(b *testing.B) {
	p := new(struct{
		Name string
	})

	j := Json(`{"name":"convert"}`)
	for i := 0; i < b.N; i++ {
		_ = j.ToStruct(p)
	}
}


func BenchmarkJson_ToMap(b *testing.B) {

	j := Json(`{"name":"convert"}`)
	for i := 0; i < b.N; i++ {
		_, _ = j.ToMap()
	}
}
