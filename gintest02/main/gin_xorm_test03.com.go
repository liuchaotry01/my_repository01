
package main

import (
	"github.com/go-xorm/xorm"
	"log"
	_"github.com/go-sql-driver/mysql"
		"github.com/gin-gonic/gin"
	"fmt"
)

type User2 struct {
	Id int64
	Name string
	GroupId int64 `xorm:"index"`
}

type Group1 struct {
	Id int64
	Name string
}

type User2Group1 struct {
	User2 `xorm:"extends"`
	Name string
}


var engine2 *xorm.Engine

func init(){
	var err error
	engine2, err = xorm.NewEngine("mysql", "root:liuchao@tcp(localhost:3306)/mysql?charset=utf8")
	if err != nil{
		log.Fatal("数据库连接失败", err)
	}
	if err := engine2.Sync2(new(User2),new(Group1),new(User2Group1));err != nil{
		log.Fatal("数据表同步失败",err)
	}
}

func main(){
	router := gin.Default()
	router.GET("/join/test",MyJoin)
	router.Run(":8090")
}

func (User2Group1) TableName() string{
	return "user"
}

func MyJoin(c *gin.Context){
	users := make([]User2Group1,0)
	err := engine2.Join("INNER","group1","group1.id = user.group_id").Find(&users)
	if err != nil{
		fmt.Println(err)
	}else{
		c.JSON(200,gin.H{
			"msg": users,
		})
	}
}


