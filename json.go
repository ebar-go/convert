package convert

import "encoding/json"

// json类型
type Json string

func (s Json) ToByte() []byte {
	return stringToByte(string(s))
}

// 转struct, v需要为指针
func (s Json) ToStruct(v interface{}) error {
	return json.Unmarshal(s.ToByte(), v)
}

// 转map
func (s Json) ToMap() (map[string]interface{}, error) {
	item := make(map[string]interface{})
	err := s.ToStruct(&item)
	return item, err
}
