package constant_test

import (
	"fmt"
	"testing"
)

// 自定义类型
type testInt int64

func TestConstant(t *testing.T) {
	var a int64 = 11
	var b int64
	b = a
	fmt.Println(b)
}
func TestConstant1(t *testing.T) {
	var a int64 = 11
	var b testInt
	b = testInt(a)
	fmt.Println(b)
}
