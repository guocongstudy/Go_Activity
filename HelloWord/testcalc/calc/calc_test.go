package calc

import (
	"fmt"
	"testing"
)

//这个函数要放一起就没事，不放一起放另一个文件中会产生undefined
func Add1(x, y int) int {
	return x + y
}

//单元测试，单元测试是用来测功能的
func TestAdd(t *testing.T) {
	if 5 != Add1(1, 4) {
		fmt.Println("计算出错")
	}
}
