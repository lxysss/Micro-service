package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Student struct {
	Id   int
	Name string
	Age  int
}

type UserNew struct {
	ID            int           //用户编号
	Name          string        `gorm:"size:32;unique"`  //用户名
	Password_hash string        `gorm:"size:128" `       //用户密码加密的  hash
	Mobile        string        `gorm:"size:11" ` //手机号
	Real_name     string        `gorm:"size:32" `        //真实姓名  实名认证
	Id_card       string        `gorm:"size:20" `        //身份证号  实名认证
	Avatar_url    string        `gorm:"size:256" `       //用户头像路径       通过fastdfs进行图片存储
}

var GlobalConn *gorm.DB

func main() {
	// 连接数据库
	fmt.Println()
	// 链接数据库 --格式: 用户名:密码@协议(IP:port)/数据库名？xxx&yyy&
	conn, err := gorm.Open("mysql", "root:Sugon123456.@tcp(9.135.154.47:3306)/search_house?parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("gorm.Open err", err)
		return
	}
	GlobalConn = conn
	// 初始数
	GlobalConn.DB().SetMaxIdleConns(10)
	// 最大数
	GlobalConn.DB().SetMaxOpenConns(100)

	// 不要复数表名 创建表
	GlobalConn.SingularTable(true)
	fmt.Println(GlobalConn.AutoMigrate(new(UserNew)).Error)

	InsertData()
}

func InsertData() {
	// 插入数据
	var stu UserNew
	stu.Name = "po0tterdu"
	stu.Password_hash = "18"
	fmt.Println(GlobalConn.Create(&stu).Error)
}

func SearchData() {
	// 查询数据
	//var stu []serviceCounterpart
	//fmt.Println(GlobalConn.Find(&stu).Error)
	//fmt.Println(stu)
	//GlobalConn.Last(&stu)
	//fmt.Println(stu)

	var res []Student
	GlobalConn.Find(&res)
	//GlobalConn.Unscoped().Find(&res)
	fmt.Println(res)

}

func updateData() {
	// 更新数据
	//var stu Student
	//stu.Name = "李四d"
	//stu.Age = 22
	//fmt.Println(GlobalConn.Save(&stu).Error)
	//fmt.Println(GlobalConn.Model(new(Student)).Where("name = ?","potterliu").
	//	Update("name","李四").Error)
	fmt.Println(GlobalConn.Model(new(Student)).Where("name = ?", "李四").
		Updates(map[string]interface{}{"name": "李四c", "age": 12}).Error)
}

func deleteData() {
	fmt.Println(GlobalConn.Where("name = ?", "李四").Delete(new(Student)).Error)
	//fmt.Println(GlobalConn.Unscoped().Delete(new(Student)).Error)
}
