//package user
//
//import "fmt"
//
//type User struct {
//}
//
//func getUsers() []*User {
//	fmt.Println()
//	return nil
//}
//
//func addrUser(name string) {
//
//}
//func deleteUser(id int64) {
//
//}
//
//func main() {
//	//dns := "root:password@tcp"
//}
package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(nums, k)
	fmt.Println(nums)
}

func rotate(nums []int, k int) {
	k = k % len(nums)
	nums = append(nums[len(nums)-k:], nums[0:len(nums)-k]...)
}
