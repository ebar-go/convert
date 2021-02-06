package convert

import (
	"strconv"
)

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
type intArray []int

func (a intArray) ToString() []string {
	result := make([]string, len(a))

	for idx, item := range a {
		result[idx] = strconv.Itoa(item)
	}

	return result
}

func (a intArray) ToInt() []int {
	return a
}
func (a intArray) ToInterface() []interface{} {
	result := make([]interface{}, len(a))

	for idx, item := range a {
		result[idx] = item
	}

	return result
}

// 字符串数组
type stringArray []string

func (a stringArray) ToString() []string {
	return a
}

func (a stringArray) ToInt() []int {
	result := make([]int, len(a))

	for idx, item := range a {
		n, _ := strconv.Atoi(item)
		result[idx] = n
	}

	return result
}

func (a stringArray) ToInterface() []interface{} {
	result := make([]interface{}, len(a))

	for idx, item := range a {
		result[idx] = item
	}

	return result
}

// 整形数组
func IntArray(items []int) Array {
	return intArray(items)
}

// 字符串数组
func StringArray(items []string) Array {
	return stringArray(items)
}