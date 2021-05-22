package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func addN(a, b int, args ...int) int {
	return 0
}

//callback 格式化 将传递的数据按照每行打印还是按照一行按|分隔打印
func print(callback func(...string),args ...string){
	fmt.Println("print函数输出：")
	callback(args...)
}

func list(args ...string){
	for i,v :=range args{
		fmt.Println(i,":",v)
	}
}
func main() {
	fmt.Println(add(1, 2))
	fmt.Printf("%T\n", addN)
	//函数类型的变量
	var f func(int, int) int = add
	fmt.Printf("%T\n", f)
	fmt.Println(f(1, 4))

	print(list,"A","C","E","g","h")

	//匿名函数
	sayHello :=func(name string){
		fmt.Println("Hello",name)
	}
	sayHello("KK")


}
