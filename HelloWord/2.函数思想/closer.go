package main

import "fmt"

func main(){
	name :="KK"

	add2 :=func(n int)int{
		return n+2
		}

	fmt.Println(add2(4))
	func(){
		fmt.Println(name)
	}()

	addBase :=func(base int) func(int)int{
		//返回函数
		return func(n int)int{
			return base +n
		}
	}
	add8 :=addBase
	fmt.Printf("%T\n",add8)
	fmt.Println(add8(10))
}
