package main

import (
	"fmt"
	"time"
)

func main() {
	names := make([]string, 10)
	dnames := make(map[string]string)

	go func(names []string, dnames map[string]string) {
		fmt.Println(names)
		fmt.Println(dnames)
	}(names, dnames)

	//names = append(names, "guocong")
	names[0] = "kk"
	dnames["kk"] = "test"

	time.Sleep(time.Second * 4)

}

//27:39
