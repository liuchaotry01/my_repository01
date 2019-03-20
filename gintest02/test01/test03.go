package main

import (
	"github.com/go-xorm/xorm"
	"log"
		_"github.com/go-sql-driver/mysql"
	"fmt"
	)

type Skhdfialsdddd struct {
	Id int64
	Name string
	CreatedAt string
}

type Utest struct {
	Id int64
	Nmae string
}

var Engine1 *xorm.Engine

func init(){
	var err error
	Engine1, err = xorm.NewEngine("mysql", "root:liuchao@tcp(localhost:3306)/lcctest01?charset=utf8")
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
	//if err := Engine1.Sync(new(User)); err != nil {
	//	log.Fatal("数据表同步失败:", err)
	//}
}
//const base_format = "2006-01-02 15:04:05.000000"

func main(){

	//ut := make([]Utest,0)
	//Engine1.Table()
	//
	//user:=User{}

	//user.Id = 8
	//user.Name = "liuchao07"
	//nt := time.Now()
	//fmt.Println(nt.Unix())
	//str_time := nt.Format(base_format)
	//fmt.Printf("now time string:%v\n", str_time)
	//user.CreatedAt=	str_time
	//_,err := Engine1.Insert(&user)
	//fmt.Println(err)
	//time.Now().Unix()
	//1544425962

		user1 := make([]Skhdfialsdddd,0)
		Engine1.Table("user").Where("id=8").Find(&user1)
		fmt.Println(user1)
	//loc, _ := time.LoadLocation("Asia/Shanghai")
	//parse_str_time,_:= time.ParseInLocation(base_format, user1[0].CreatedAt, loc)
	//	fmt.Printf("string to datetime :%v\n", parse_str_time.Unix())




}