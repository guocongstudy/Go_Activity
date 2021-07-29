package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func main() {
	response, err := http.Get("http://localhost:8080")
	fmt.Println(err)
	io.Copy(os.Stdout, response.Body)

	response, err = http.Head("http://localhost:8080?x=1&y=1")
	fmt.Println(err)
	io.Copy(os.Stdout, response.Body)

	//x-www-form-urlencoded
	values := url.Values{}
	values.Add("x", "1")
	values.Add("X", "2")
	values.Set("y", "1")

	http.PostForm("http://localhost:8080", values)
}
