package main

import "fmt"

//冒泡排序
func main() {
	heights := []int{10, 6, 7, 9, 5}
	//先把最高的人移动到最后
	//从第一个人开始，两两开始比较，如果前面的人高，这两人就交换位置
	//6,10,7,9,5
	//6,7,10,9,5
	//6,7,9,10,5
	//6,7,9,5,10
	//n个人比较n-1次
	//交换a=1,b=2,
	for j:=0;j<len(heights)-1;j++{
		fmt.Println("第",j,"轮")
		for i := 0; i < len(heights)-1; i++ {
			if heights[i] > heights[i+1] {
				fmt.Println("交换：",heights[i],heights[i+1])
				tmp :=heights[i]
				heights[i]=heights[i+1]
				heights[i+1]=tmp
			}
			fmt.Println(i,"交换完毕",heights)
		}
		fmt.Println("第",j,"轮，结果：",heights)
	}
	}

