package main

import (
	"fmt"
	"sort"
)
//切片的排序
func main(){
	nums :=[]int{4,5,6,7}
	sort.Ints(nums)
	fmt.Println(nums)

	names := []string {"test","kk","123","zzz","xxx"}
	sort.Strings(names)
	fmt.Println(names)

	heights :=[]float64{1.1,-1.1,3,2}
	sort.Float64s(heights)
}
