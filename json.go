package convert

import "encoding/json"

type Json string

func (s Json) ToByte() []byte {
	return stringToByte(string(s))
}

func (s Json) ToStruct(v interface{}) error {
	return json.Unmarshal(s.ToByte(), v)
}

func (s Json) ToMap() (map[string]interface{}, error) {
	item := make(map[string]interface{})
	err := s.ToStruct(&item)
	return item, err
}
