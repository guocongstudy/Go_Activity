package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	driverName := "mysql"
	//传参格式："user:password@protocol(host:port)/dbname?charset=utf8mb4&loc=Local&parseTime=true"
	dsn := "root:password@tcp(127.0.0.1:3306)/user?charset=utf8mb4&loc=Local&parseTime=true"
	//data store name 数据库连接信息，使用协议，用户名&密码，数据库，连接数据库
	db, err := sql.Open(driverName, dsn)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		fmt.Println(err)
		return
	}

	//操作
	//不要輕易去用* 如果后面进行改表，那就会出现问题
	//rows, err := db.Query("select * from user ")
	rows, err := db.Query("select id,name,password,sex,birthday,addr,tel from user ")
	if err != nil {
		fmt.Println(err)
		return
	}
	//要把创建的资源rows延迟关闭
	defer rows.Close()

	for rows.Next() {
		var (
			id       int64
			name     string
			password string
			sex      bool
			birthday *time.Time
			addr     string
			tel      string
		)
		//进行扫描
		err := rows.Scan(&id, &name, &password, &sex, &birthday, &addr, &tel)
		if err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println(id, name, password, sex, birthday, addr, tel)
		}
	}
	//查询一行数据
	var id int64
	err2 := db.QueryRow("select id from user ").Scan(&id)
	fmt.Println(err2)
}
