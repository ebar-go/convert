package convert

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestResponse(t *testing.T) {
	resp := new(http.Response)
	content := "hello,world"
	resp.Body = ioutil.NopCloser(strings.NewReader(content))

	_, err := Response(resp)
	assert.Nil(t, err)
}

func TestResponse_ToString(t *testing.T) {
	resp := new(http.Response)
	content := "hello,world"
	resp.Body = ioutil.NopCloser(strings.NewReader(content))

	response, err := Response(resp)
	assert.Nil(t, err)
	assert.Equal(t, content, response.ToString())
}


func TestResponse_ToByte(t *testing.T) {
	resp := new(http.Response)
	content := "hello,world"
	resp.Body = ioutil.NopCloser(strings.NewReader(content))

	response, err := Response(resp)
	assert.Nil(t, err)
	assert.Equal(t, []byte(content), response.ToByte())
}

func BenchmarkResponse_ToByte(b *testing.B) {
	resp := new(http.Response)
	content := "hello,world"
	resp.Body = ioutil.NopCloser(strings.NewReader(content))

	response, err := Response(resp)
	assert.Nil(b, err)
	for i := 0; i < b.N; i++ {
		response.ToByte()
	}
}

func BenchmarkResponse_ToString(b *testing.B) {
	resp := new(http.Response)
	content := "hello,world"
	resp.Body = ioutil.NopCloser(strings.NewReader(content))

	response, err := Response(resp)
	assert.Nil(b, err)
	for i := 0; i < b.N; i++ {
		_ = response.ToString()
	}
}


func BenchmarkResponse(b *testing.B) {
	resp := new(http.Response)
	content := "hello,world"
	resp.Body = ioutil.NopCloser(strings.NewReader(content))

	for i := 0; i < b.N; i++ {
		_, _ = Response(resp)
	}
}