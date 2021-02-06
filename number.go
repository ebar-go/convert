package convert

import "strconv"

type number struct {
	v int64
}

// 整形
func Number(v int) *number {
	return &number{v: int64(v)}
}

// 转字符串
func (n *number) ToString() string {
	return strconv.FormatInt(n.v, 10)
}

// 转float
func (n *number) ToFloat() float64 {
	return float64(n.v)
}

