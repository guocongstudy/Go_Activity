package calc_value

import "testing"

func Face(x int) int {
	if x == 0 {
		return 1
	} else {
		return x * Face(x-1)
	}
}

//基准测试，测试性能的(函数写的好坏，复不复杂)
func BenchmarkFace(b *testing.B) {
	for i := 1; i <= 10000; i++ {
		Face(i)
	}
}
