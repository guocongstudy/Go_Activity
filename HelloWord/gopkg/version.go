package gopkg

import "fmt"

const VERSION = 1.1
const name = "local gopkg"

func PrintName() {
	fmt.Println(name)
}
func init() {
	fmt.Println("init")
}
