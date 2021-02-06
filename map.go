package convert

import "encoding/json"

type Map map[string]interface{}

// 转json
func (m Map) ToJson() string {
	b, _ := json.Marshal(m)
	return Byte(b).ToString()
}

func (m Map) ToStruct(obj interface{}) error {
	return nil
}