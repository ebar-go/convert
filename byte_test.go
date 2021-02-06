package convert

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestByte(t *testing.T) {
	b := Byte([]byte{'h','e'})
	assert.Equal(t, "he", b.ToString())

	n := Byte([]byte{0, 0, 0, 1})
	assert.Equal(t, 1, n.ToInt())
}
