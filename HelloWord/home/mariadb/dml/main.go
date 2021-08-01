package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	driverName := "mysql"
	dsn := "root:password@tcp(127.0.0.1:3306)/user?parseTime=true&loc=Local&charset=utf8mb4"
	db, err := sql.Open(driverName, dsn)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	sql := `
    INSERT  INTO user(id,password,birthday) VALUE (5,"1223","1997-08-23");

`
	result, err := db.Exec(sql)
	if err != nil {
		fmt.Println(err)
	} else {
		//最后一次插入的id
		fmt.Println(result.LastInsertId())
		//影响的行数
		fmt.Println(result.RowsAffected())
	}

	//删除和更新
	sql = `
     UPDATE user 
SET addr='北京'
`
	result, err = db.Exec(sql)
	fmt.Println(err)
	fmt.Println(result.LastInsertId())
	//影响行数
	fmt.Println(result.RowsAffected())
}
