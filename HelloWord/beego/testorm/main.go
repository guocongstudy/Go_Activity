package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"

	"log"
)

//標籤 orm
//列名 column
//類型 type
//decimal类型digits,decimals
//字符串長度 size
//主鍵 自動增長 pk quto
//默认值 default
type User struct {
	ID        int64      `orm:"column(uid);pk;auto"`
	Name      string     `orm:"size(64);default(aaa)"`
	Password  string     `orm:"size(1024)"`
	Sex       bool       `orm:"size(32)"`
	Height    float64    `orm:"digits(12),decimals(3)"`
	Tel       string     `orm:"index"`
	Addr      string     `orm:"type(text)"`
	Birthday  *time.Time `orm:"type(date)"`
	CreatedAt *time.Time `orm:"auto_now_add;description(創建時間)"`
	UpdatedAt *time.Time `orm:"auto_now"`
	DeleteAt  *time.Time `orm:"nul"`
}

func (u *User) TableName() string {
	return "a1"
}

func (u *User) TableIndex() [][]string {
	return [][]string{
		{"Name"},
		{"Password"},
		{"Tel"},
	}
}

func main() {
	dsn := "root:password@tcp(127.0.0.1:3306)/mysql?parseTime=true&loc=Local&charset=utf8mb4"
	//导入驱动(初始化)
	//导入orm包
	//在ORM包中注册驱动（mysql）
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//注册数据库
	if err := orm.RegisterDataBase("default", "mysql", dsn); err != nil {
		log.Fatal(err)
	}

	//注册模型
	orm.RegisterModel(&User{})

	//操作

	//DDL
	orm.RunSyncdb("default", true, true) //同步数据库
	//结构体对应表是否存在
	//表不存在 创建对应的表
	//若表存在，属性列死否在表中存在
	//属性不存在，添加列
	//索引是否存在，索引不存在添加索引
	//数据库别名
	//是否先删除所有表
	//
	//DML DQL

	//fmt.Println("klfdjklfds")

	ormer := orm.NewOrm()

	birthday, _ := time.Parse("2006-01-02", "1879-01-02")

	user := &User{
		Name:     "kk",
		Password: "123",
		Birthday: &birthday,
	}
	orm.Debug = true

	fmt.Println(user)
	fmt.Println(ormer.Insert(user))
	fmt.Println(user)
	fmt.Println(ormer.Insert(user))
	// 增，删，改，查

	//查询
	queryset := ormer.QueryTable(&User{})
	//满足条件
	fmt.Println(queryset.Count())
	var users []*User
	fmt.Println(queryset.All(&users))
	fmt.Println(users)
}
