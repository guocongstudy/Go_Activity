package main

import "fmt"

func main() {

	var nums []int
	fmt.Printf("%T\n", nums)
	fmt.Printf("%#v\n %d %d \n", nums, len(nums), cap(nums))
	fmt.Println(nums == nil)

	//字面量
	nums = []int{1, 2, 3}
	fmt.Printf("%#v\n %d %d \n", nums, len(nums), cap(nums))
	nums = []int{1, 2, 3, 4}
	fmt.Printf("%#v\n %d %d \n", nums, len(nums), cap(nums))

	//数组切片赋值
	var arrays [10]int = [10]int{1, 2, 3, 4}
	nums = arrays[1:10]
	fmt.Printf("%#v\n %d %d \n", nums, len(nums), cap(nums))

	//make函数
	nums = make([]int, 3)
	fmt.Printf("%#v %d %d", nums, len(nums), cap(nums))
	nums = make([]int, 3, 5) //容量为5，只用了3个容量
	fmt.Printf("%#v %d %d\n", nums, len(nums), cap(nums))

	//元素操作（增删改查）
	fmt.Println(nums[0])
	fmt.Println(nums[1])
	fmt.Println(nums[2])
	nums[2] = 10
	fmt.Println(nums)

	nums = append(nums, 1)
	fmt.Printf("%#v %d %d\n", nums, len(nums), cap(nums))

	for i := 0; i < len(nums); i++ {
		fmt.Println(i, nums[i])
	}

	for index, value := range nums {
		fmt.Println(index, value)
	}

	//切片操作
	fmt.Printf("%T\n", nums[1:5])
	//切片操作
	n := nums[1:3:5]
	fmt.Printf("%T %#v %d %d\n", n, n, len(n), cap(n))
	n = nums[2:3]
	//src-cap-start
	fmt.Printf("%T %#v %d %d\n", n, n, len(n), cap(n))

	//删除
	//go中用copy实现删除的机制
	nums04 := []int{1, 2, 3}
	nums05 := []int{10, 30, 20, 40}

	copy(nums05, nums04) //将04copy到05 05在前的
	fmt.Println(nums05)

	nums05 = []int{10, 20, 30, 40}
	copy(nums04, nums05)
	fmt.Println(nums04)
	//索引为0的，索引最后一个
	nums06 := []int{1, 2, 3, 4, 5}
	fmt.Println(nums06[1:])
	fmt.Println(nums06[:len(nums06)-1])
	//删除中间的第二个元素2
	nums07 := []int{1, 3, 4, 5}
	copy(nums06, nums07)
	fmt.Println(nums06)
	fmt.Println(nums06[:len(nums06)-1])

	//方法二
	//删除中间的元素2(3)
	copy(nums06[2:], nums06[3:])
	fmt.Println(nums06[:len(nums06)-1])

	//队列，堆栈
	//队列：每次添加在队尾，移除元素在对头（先进先出）
	//堆栈：每次添加在队尾，移除元素在对头（先进后出）
	queue := []int{}
	queue = append(queue, 1)
	queue = append(queue, 2)
	queue = append(queue, 4)
	queue = append(queue, 5)
	//1,2,3,5
	fmt.Println(queue[0])
	queue = queue[1:]
	//2,3,5
	fmt.Println(queue)
	fmt.Println(queue[0])
	queue = queue[1:]
	//3,5
	fmt.Println(queue)

	stack := []int{}
	stack = append(stack, 1)
	stack = append(stack, 2)
	stack = append(stack, 3)
	fmt.Println(stack)

	fmt.Println(stack[len(stack)-1])
	stack = stack[:len(stack)-1]
	fmt.Println(stack[len(stack)-1])
	stack = stack[:len(stack)-1]
	fmt.Println(stack[len(stack)-1])
	stack = stack[:len(stack)-1]

	//多维切片
	points := [][]int{}
	points02 := make([][]int, 0)
	fmt.Println(points02)
	points = append(points, []int{1, 2, 3})
	points = append(points, []int{3, 4, 0})
	points = append(points, []int{3, 4, 0, 2, 4, 5})
	fmt.Println(points)
	fmt.Println(points[0])
	fmt.Println(points[0][1]) //第0组中的第2个 {1,2,3}中的2

	//数组是值类型
	slice01 := []int{1, 2, 3}
	slice02 := slice01

	slice02[0] = 10
	fmt.Println(slice01, slice02)
}

//len（）长度是指当前切片的长度，cap（）容量是指切片底层数组的长度
