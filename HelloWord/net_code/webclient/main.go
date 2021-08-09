package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func main() {
	response, err := http.Get("http://localhost:8080")
	fmt.Println(err)
	//io.Copy(os.Stdout, response.Body)
	fmt.Println(response)
	//response, err = http.Head("http://localhost:8080?x=1&y=1")
	//fmt.Println(err)
	//io.Copy(os.Stdout, response.Body)

	//x-www-form-urlencoded
	values := url.Values{}
	values.Add("x", "1")
	values.Add("X", "2")
	values.Set("y", "1")

	http.PostForm("http://localhost:8080", values)
	fmt.Println(err)
	//io.Copy(os.Stdout, response.Body)
	//application/json
	f, _ := os.Open("test.json")
	response, err = http.Post("http://localhost:8080", "application/json", f)
	fmt.Println(err)
	//io.Copy(os.Stdout, response.Body)
}
