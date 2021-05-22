package main

import "fmt"

func main() {
	//冒泡，循环一次最后一个索引的元素
	//猴子掰包谷，记录最大元素，每次比较元素的最大值
	//初始化[0]切片/数组 元素数量0

	//找到最大值
	nums := []int{1, 4, 5, 6, 43, 32, 5, 7, 9}
	maxNum := nums[0]

	for i, v := range nums {
		if v > maxNum {
			maxNum = v
		}
		fmt.Println(i, ":", maxNum)
	}
	fmt.Println(maxNum)
}
