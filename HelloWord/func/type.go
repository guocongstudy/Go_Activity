package main

import "fmt"

type Counter int
type User map[string]string
type Callback func(...string)

func main() {
	var counter Counter
	counter += 10
	fmt.Println(counter)

	me := make(User)
	me["name"] = "郭聪"
	me["城市"] = "西安"
	fmt.Println(me)
	fmt.Printf("%T,%T\n", me, counter)
	var list Callback = func(args ...string) {
		for i, v := range args {
			fmt.Println(i, ":", v)
		}
	}
	list("a", "b", "c")
}
