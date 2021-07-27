package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
)

func main() {
	url := "www.baidu.com"

	request, _ := http.NewRequest("DELETE", url, nil)

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: transport}
	response, err := client.Do(request)

	if err == nil {
		fmt.Println(response.Proto, response.Status)
		fmt.Println(response.Header)

		//io.Copy(os.Stdout,response.Body)
		response.Write(os.Stdout)
	}

}
