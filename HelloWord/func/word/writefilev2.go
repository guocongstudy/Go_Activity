package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	err := ioutil.WriteFile("user.txt", []byte("xxxxxxxxxxxxxxxx"), os.ModePerm)
	fmt.Println(err)
}
