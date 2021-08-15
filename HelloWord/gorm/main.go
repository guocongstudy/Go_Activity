package main

import (
	"fmt"
	"os"
)

type User struct {
	Name     string
	Password string
	Birthday string
	Sex      bool
	Tel      string
	Addr     string
	Desc     string
}

var (
	dbName     string = ""
	dbUser     string
	dbPassword string
	dbHost     string
	dbPort     int
)
var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&local&parseRime=trie",
	dbUser, dbPassword, dbHost, dbPort, dbName)

func main() {
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	//在数据库中添加索引
	db.Model(&User{}).AddUniqueIndex("idx_name", "name")
	//删除数据库的索引
	db.Model(&User{}).RemoveIndex("idx_name")
	db.Close

}

//37：57
