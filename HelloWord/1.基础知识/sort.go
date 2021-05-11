package main

import (
	"fmt"
	"sort"
)

func main() {
	//排序
	nums := []int{1, 2, 45, 3, 5, 6}
	//对数组进行排序但不输出
	sort.Ints(nums)
	fmt.Println(nums)
	//二分查找，在有序的数组中查找元素,函数输出的是要插入位置的索引
	fmt.Println(sort.SearchInts(nums, 3))

	num := 5
	idx := sort.SearchInts(nums, 5)
	if nums[idx] == num {
		fmt.Println("存在")
	} else {
		fmt.Println("不存在")
	}
}
