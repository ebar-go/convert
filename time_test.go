package convert

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTimestamp(t *testing.T) {
	time := Timestamp(1612510505)
	assert.Equal(t, time.UnixNano(), time.ToTimestamp())
	assert.Equal(t, "2021-02-05 15:35:05", time.ToString())
	assert.Equal(t, "2021-02-05", time.ToDate())
}
