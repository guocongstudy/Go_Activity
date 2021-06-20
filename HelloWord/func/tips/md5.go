package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	bytes := md5.Sum([]byte("i'am kk"))
	x := fmt.Sprintf("%x", bytes)
	fmt.Println(hex.EncodeToString(bytes[:]))
	fmt.Println(x)

	m := md5.New()
	m.Write([]byte("i'am"))
	m.Write([]byte("kk"))

	fmt.Println("%x\n", m.Sum(nil))

}
