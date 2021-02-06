package convert

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumber(t *testing.T) {
	n := Number(1)
	assert.Equal(t, "1", n.ToString())
	assert.Equal(t, float64(1), n.ToFloat())
	fmt.Println(n.ToByte())
}