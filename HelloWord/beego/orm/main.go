package main

/*
orm是设计模式
*/

//import (
//	"fmt"
//	"github.com/jinzhu/gorm"
//	_ "github.com/jinzhu/gorm/dialects/mysql"
//	"os"
//	"time"
//)
//
//type User struct {
//	gorm.Model
//	Name     string
//	Password string
//	Birthday time.Time
//	Sex      bool
//	Tel      string
//	Addr     string
//	Desc     string
//}
//
//const (
//	dbUser     string = "root"
//	dbPassword string = "password"
//	dbHost     string = "127.0.0.1"
//	dbPort     int    = 3306
//	dbName     string = "user"
//)
//
//var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset-utf8mb4&loc=local&parseTime=true",
//	dbUser, dbPassword, dbHost, dbPort, dbName)
//
//func main() {
//	db, err := gorm.Open("mysql", dsn)
//	if err != nil {
//		fmt.Println(err)
//		os.Exit(-1)
//	}
//	db.AutoMigrate(&User{})
//}

//gorm初始
//00:24:54
