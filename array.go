package convert

// 整形数组
type IntArray []int
// 转字符串数组
func (a IntArray) ToString() []string {
	result := make([]string, len(a))

	for idx, item := range a {
		result[idx] = intToString(item)
	}

	return result
}
// 转interface数组
func (a IntArray) ToInterface() []interface{} {
	result := make([]interface{}, len(a))

	for idx, item := range a {
		result[idx] = item
	}
	return result
}

// 字符串数组
type StringArray []string
// 转int数组
func (a StringArray) ToInt() []int {
	result := make([]int, len(a))

	for idx, item := range a {
		result[idx] = stringToInt(item)
	}
	return result
}
// 转interface数组
func (a StringArray) ToInterface() []interface{} {
	result := make([]interface{}, len(a))

	for idx, item := range a {
		result[idx] = item
	}
	return result
}
