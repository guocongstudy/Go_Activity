package main

import "fmt"

func test() (err error){
	defer func(){
		if e :=recover();e !=nil{  //发生错误做一些恢复处理
			err = fmt.Errorf("%v",e)
		}
	}()
	panic("error")
	return
}

func main(){
   err :=test()
   fmt.Println(err)
	//defer func(){
	//	if err :=recover();err !=nil{  //发生错误做一些恢复处理
	//		fmt.Println(err)
	//	}
	//}()

}
