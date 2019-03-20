package main

import (
	"github.com/go-xorm/xorm"
		"github.com/gin-gonic/gin"
	"fmt"
	_"github.com/go-sql-driver/mysql"
	"log"
)

type Group struct {
	Id int64
	Name string
}

type Test struct {
	Id int64
	Name string
	GroupId int64 `xorm:"index"`
}

type TestGroup struct {
	Test `xorm:"extends"`
	Name string
}

func (TestGroup) TableName() string {
	return "test"
}

var X  *xorm.Engine

func init() {
	var err error
	X, err = xorm.NewEngine("mysql", "root:liuchao@tcp(localhost:3306)/lcctest01?charset=utf8")
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
	if err := X.Sync(new(Test),new(Group)); err != nil {
		log.Fatal("数据表同步失败:", err)
	}
	//err1 := X.Ping()
	//fmt.Println(err1)
	//X.Close()
}

func main(){
	//fmt.Println("nihao")
	router := gin.Default()

	router.GET("/join/test",MyJoin)
	router.GET("/count/test",MyCount)
	router.GET("/rows/test",MyRows)
	router.GET("/sum/test",MySum)
	router.GET("/update/test",MyUpdate)
	router.Run(":8090")
}

//func (UserGroup) TableName() string {
//	return "user"
//}

func MyUpdate(c *gin.Context){
	test := new(Test)
	test.Name = "chaochao01"
	test.GroupId = 12
	//affected, err := X.Id(1).Update(test)
	affected, err := X.Id(2).Cols("group_id").Update(test)
	if err != nil{
		fmt.Println(err)
	}else{
		c.JSON(200, gin.H{
			"msg": affected,
		})
	}
}

func MySum(c *gin.Context){
	test := Test{}
	total, err := X.Where("id < ?",3).Sum(test,"group_id")
	if err != nil{
		fmt.Println(err)
	}else{
		c.JSON(200, gin.H{
			"msg": total,
		})
	}
}

func MyRows(c *gin.Context){
	test := new(Test)
	rows, err := X.Where("id < ?",3).Rows(test)
	if err != nil{
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next(){
		err = rows.Scan(test)
		c.JSON(200, gin.H{
			"msg": test,
		})
	}
}

func MyCount(c *gin.Context){
	test := new(Test)
	total, err := X.Where("id < ?",3).Count(test)
	if err != nil{
		fmt.Println(err)
	}else{
		c.JSON(200, gin.H{
			"total": total,
			"test":  test,
		})
	}
}

func MyJoin(c *gin.Context){
	tests := make([]TestGroup, 0)
	err := X.Join("INNER", "group", "group.id = test.group_id").Find(&tests)
	//err := X.SQL("select test.*, group.name from test, group where test.group_id = group.id").Find(&tests)
	if err != nil{
		fmt.Println(err)
	}else{
		c.JSON(200, gin.H{
			"msg": tests,
		})
	}
}

