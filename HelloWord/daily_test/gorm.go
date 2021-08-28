package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	dbUser     string = "root"
	dbPassword string = "password"
	dbHost     string = "127.0.0.1"
	dbPort     int    = 3306
	dbName     string = "test"
)

type User struct {
	Id       int    `gorm:"primary_key，auto_increment"` //设置主键自增
	Name     string `gorm:"size(255);unique;not null;default: ''"`
	Password string
	Sex      bool   `gorm:"index:index_desc"`
	Desc     string `gorm:"type:text;not null"`
}

func (u *User) TableName() string {
	return "user"
}

var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=local&parseTime=true",
	dbUser, dbPassword, dbHost, dbPort, dbName)

//func main(){
//	db,err :=gorm.Open("mysql",dsn)
//	if err !=nil{
//		fmt.Println(err)
//		//退出一下
//		os.Exit(-1)
//	}
//	//删除一列
//	db.Model(&User{}).RemoveIndex("idx_sex")
//	//迁移一下数据库
//	db.AutoMigrate(&User{})
//	for i:=0;i<10;i++{
//		user :=User{
//			Name :fmt.Sprintf("kk_%s",i),
//			Password:fmt.Sprintf("password_%s",i),
//			Sex: true,
//			Desc: "dwdwdwdwd",
//		}
//		db.Create(user)
//		fmt.Println(user)
//	}
//	db.Close()
//}
//func main() {
//	db, err := gorm.Open("mysql", dsn)
//	if err != nil {
//		fmt.Println(err)
//		//退出一下
//		os.Exit(-1)
//	}
//	var user User
//	db.First(&user, "name = ?", "kk")
//	fmt.Println(user)
//
//	var user2 User
//	db.Last(&user2, "name = ?", "kk_2")
//	fmt.Println(user)
//
//	var users []User
//	//db.Where("name = ? and password = ?","kk_5","password_6").Find(&users)
//	//db.Where("name in (?)",[]string{"kk_6","KK_8"}).Find(&users)
//
//	//select
//	db.Select("name,password").Find(&user)
//	//形式二
//	db.Select([]string{"name", "password"}).Find(&users)
//	fmt.Println(users)
//
//	//db.Model(&User{}).Where("name=?","kk_2").Count(&count)
//	//fmt.Println(users)
//
//	var count int
//	db.Model(&User{}).Where("name = ?", "kk_2").Count(&count)
//	fmt.Println(count)
//
//	db.Model(&User{}).Rows()
//
//	rows, _ := db.Model(&User{}).Select("name.password").Rows()
//	for rows.Next() {
//		var name, password string
//		rows.Scan(&name, &password)
//	}
//
//	db.Model(&User{}).Select("name,count(*) as cnt").Group("name").Rows()
//	for rows.Next() {
//		var name string
//		var count int
//		rows.Scan(&name, &count)
//		fmt.Println(name, count)
//	}
//
//	db.Close()
//
//	db = db.Model(&User{})
//	db = db.Select("name,count(*) as cnt")
//	db = db.Group("name")
//	db = db.Having("count(*) > ?", 1)
//	//rows := db.Rows()
//}

//func (u *User) First() *User{
//	var uu *User
//	return uu
//}
//
//func main(){
//	db,err :=gorm.Open("mysql",dsn)
//	if err !=nil{
//		fmt.Println(err)
//		os.Exit(-1)
//	}
//	db.LogMode(true)
//	db.AutoMigrate(&User{})
//
//	//查找对象进行更新
//	var user User
//	db.First(&user,"name = ?","kk_20")
//
//	//查找对象进行更新
//	if db.First(&user,"name=?","kk_2").Error ==nil{
//		user.Name="郭聪"
//		db.Save(user)
//	}
//	db.Model(&User{}).Where("id>?",10).UpdateColumn("sex",true)
//	db.Model(&User{}).Where("id > ?",15).UpdateColumns(map[string]interface{}{"desc":"abc","name":"中国"})
//	db.Model(&User{}).Where("id > ?",18).Updates(User{Desc:"xxxx",Sex:true})
//
//	db.Close()
//}


//func main(){
//	db,err :=gorm.Open("mysql",dsn)
//	if err !=nil{
//		fmt.Println(err)
//		os.Exit(-1)
//	}
//	db.LogMode(true)
//	db.AutoMigrate(&User{})
//
//	//查找对象进行更新
//	var user User
//	//if db.First(&user,"name","kk_2").Error ==nil{
//	//	db.Delete(&user)
//	//}
//
//	db.First(&user)
//
//	fmt.Println(user)
//	user.Sex =false
//	db.Save(&user)
//	db.Model(&User{}).Where("id=2").UpdateColumn("sex",true)
//	db.Model(&User{}).Where("id=3").UpdateColumns(map[string]interface{}{"sex":true})
//
//	db.Model(&User{}).Where("id=4").Updates("sex",true)
//
//	db.Unscoped().Delete(&user)
//
//	db.Exec("insert into user3(name) where id =?","kk4")
//	db.Where("id > ?",16).Delete(&User{})
//	db.Close()
//
//}

var db *gorm.DB

func init(){
	db, err :=gorm.Open("mysql",dsn)

	if err !=nil || db.DB().Ping() !=nil{

		panic("Error cannot DB")
	}
}