package convert

import "strconv"

// 数组
type Array interface {
	// 转整形数组
	ToInt() []int
	// 转字符串数组
	ToString() []string
	// 转interface
	ToInterface() []interface{}
}

// 整形数组
type intArray struct {
	items []int
}

func (a *intArray) ToString() []string {
	result := make([]string, 0)

	for _, item := range a.items {
		result = append(result, strconv.Itoa(item))
	}

	return result
}

func (a *intArray) ToInt() []int {
	return a.items
}
func (a *intArray) ToInterface() []interface{} {
	result := make([]interface{}, len(a.items))

	for idx, item := range a.items {
		result[idx] = item
	}

	return result
}

// 字符串数组
type stringArray struct {
	items []string
}

func (a *stringArray) ToString() []string {

	return a.items
}

func (a *stringArray) ToInt() []int {
	result := make([]int, 0)

	for _, item := range a.items {
		n, _ := strconv.Atoi(item)
		result = append(result, n)
	}

	return result
}

func (a *stringArray) ToInterface() []interface{} {
	result := make([]interface{}, len(a.items))

	for idx, item := range a.items {
		result[idx] = item
	}

	return result
}

// 整形数组
func IntArray(items []int) Array {
	return &intArray{items}
}

// 字符串数组
func StringArray(items []string) Array {
	return &stringArray{items}
}
