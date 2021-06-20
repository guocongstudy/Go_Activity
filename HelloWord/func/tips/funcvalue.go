package main

import "fmt"

func changeInt(a int) {
	a = 100
}
func changeSlice(s []int) {
	s[0] = 100
}

func main() {
	num := 1
	changeInt(num)
	fmt.Println(num)

	nums := []int{1, 2, 3}
	changeSlice(nums)
	fmt.Println(nums)
}
