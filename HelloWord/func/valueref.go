package main

import "fmt"

func main() {
	//值类型，引用类型
	//将变量赋值给新的一个变量，并修改新变量的值，如果对旧变量有影响引用类型，无影响值类型
	array := [3]string{"A", "B", "C"} //数组
	slice := []string{"A", "B", "C"}  //切片
	/*数组和切片的区别：
	  1.数组值类型，切片指针类型
	  2.数组的长度是固定的，切片不是
	  3.切片比数组多一个容量属性，
	  4.切片的底层是数组
	*/
	arrayA := array
	sliceA := slice

	arrayA[0] = "z"
	sliceA[0] = "z"
	fmt.Println(arrayA, array)
	fmt.Println(sliceA, slice)
	fmt.Printf("%p %p\n", &arrayA, &array) //%p取地址符
	fmt.Printf("%p,%p\n", &sliceA, &slice) ///指针类型和引用数据类型在赋值后变量的地址并不相同，只是引用类型在底层共享数据结构
	//int bool float32 float64 array slice map(映射) 指针
	/*
	 值类型：int,bool,float, 指针
	 引用类型：slice,map,
	*/
	//验证map是什么类型？
	m := map[string]string{}
	mA := m
	mA["KK"] = "西安"
	fmt.Println(mA, m)
}
