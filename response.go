package convert

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type httpResponse struct {
	body []byte
}

// http响应
func Response(resp *http.Response) (*httpResponse, error)  {
	if resp.Body == nil {
		return nil, fmt.Errorf("no response")
	}
	defer resp.Body.Close()
	b, err := bodyReader.read(resp.Body)
	if err != nil {
		return nil, err
	}

	return &httpResponse{b}, nil
}

// 转string
func (response *httpResponse) ToString() string {
	return byteToString(response.body)
}

// 转byte
func (response *httpResponse) ToByte() []byte {
	return response.body
}

// http响应内容reader
type responseReader struct {
	pool sync.Pool
}

var bodyReader = getResponseReader(4096)

// getResponseReader
func getResponseReader(bufferLen int) *responseReader {
	return &responseReader{pool: sync.Pool{New: func() interface{} {
		return bytes.NewBuffer(make([]byte, bufferLen))
	}}}
}

// Read 读取流数据
func (adapter *responseReader) read(reader io.Reader) ([]byte, error) {
	buffer := adapter.pool.Get().(*bytes.Buffer)
	buffer.Reset()
	defer func() {
		if buffer != nil {
			adapter.pool.Put(buffer)
			buffer = nil
		}
	}()
	_, err := io.Copy(buffer, reader)

	if err != nil {
		return nil, fmt.Errorf("failed to read respone:%s", err.Error())
	}

	return buffer.Bytes(), nil
}
