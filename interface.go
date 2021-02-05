package convert

import (
	"fmt"
	"reflect"
	"strconv"
)

type Interface struct {
	val interface{}
}

func NewInterface(val interface{}) Interface {
	return Interface{val: val}
}

// 转整形
func (inter Interface) Int() int {
	var result int
	if inter.val == nil {
		return 0
	}

	v := reflect.ValueOf(inter.val)
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64:
		result = int(v.Int())
	case reflect.String:
		result, _ = strconv.Atoi(inter.val.(string))
	case reflect.Float64, reflect.Float32:
		result, _ = strconv.Atoi(fmt.Sprintf("%1.0f", v.Float()))
	case reflect.Uint, reflect.Uint8, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		result = int(v.Uint())
	case reflect.Bool:
		if v.Bool() { result = 1}
	}

	return result
}

// 转字符串
func (inter Interface) String() string {
	var result string
	if inter.val == nil {
		return ""
	}
	v := reflect.ValueOf(inter.val)
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64:
		result = strconv.FormatInt(v.Int(), 10)
	case reflect.String:
		result = v.String()
	case reflect.Float64, reflect.Float32:
		result = strconv.FormatFloat(v.Float(), 'f', -1, 64)
	case reflect.Uint, reflect.Uint8, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		result = strconv.FormatUint(v.Uint(), 10)
	case reflect.Bool:
		if v.Bool() { result = "true"} else {result = "false"}
	}

	return result
}

// 转浮点数
func (inter Interface) Float64() float64 {
	var result float64
	v := reflect.ValueOf(inter.val)
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64:
		result = float64(v.Int())
	case reflect.String:
		result, _ = strconv.ParseFloat(inter.val.(string), 64)
	case reflect.Float64, reflect.Float32:
		result = v.Float()
	case reflect.Uint, reflect.Uint8, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		result = float64(v.Uint())
	}
	return result
}

