package main

import "fmt"

type EStruct struct {
}
type Empty interface {
}

func fargs(args ...interface{}) {
	fmt.Println("--------------------")
	for _, arg := range args {
		fmt.Println(arg)
		switch v := arg.(type) {
		case int:
			fmt.Printf("Int:%T,%v", v, v)
		case string:
			fmt.Printf("string", v, v)
		default:
			fmt.Printf("other:%T %v", v, v)
		}
	}
}
func main() {
	es := EStruct{}
	var e Empty = 1
	fmt.Println(es, e)
	e = "tet"
	fmt.Println(e)

	e = true
	fmt.Println(e)

	fmt.Println(e)

}
