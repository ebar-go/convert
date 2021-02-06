package convert

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMap(t *testing.T) {

	m1 := Map{"age": 10}
	assert.Equal(t, `{"age":10}`, m1.ToJson())

	e := map[string]interface{}{"name":"test"}
	m2 := Map(e)
	assert.Equal(t, `{"name":"test"}`, m2.ToJson())
}
