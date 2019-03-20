package main

import (
	"github.com/go-xorm/xorm"
	"log"
		"github.com/gin-gonic/gin"
	"strconv"
	"fmt"
	_"github.com/go-sql-driver/mysql"
	"time"
)

type User1 struct {
	Id int64
	Name string   `xorm:"'my_name'"`
	Age int64
	CreatedAt time.Time `xorm:"created"`
}
//type Test1 struct {
//
//}

var engine1 *xorm.Engine

func init(){
	var err error
	engine1, err = xorm.NewEngine("mysql", "root:liuchao@tcp(localhost:3306)/mysql?charset=utf8")
	if err != nil{
		log.Fatal("数据库连接失败", err)
	}
	if err := engine1.Sync2(new(User1));err != nil{
		log.Fatal("数据表同步失败",err)
	}
	//engine1.Table("liu_user")
	//engine1.CreateTables(Test1{})
	//engine1.DBMetas()
	//tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "liu_")
	//engine1.SetTableMapper(tbMapper)
}

func main(){
     router := gin.Default()
     router.GET("/insert",Insert)
     router.GET("/select",Select)
     router.GET("select/sql",SelectSql)
     router.GET("/select/in",SelectIn)
     router.GET("/select/col", SelectCol)
     router.GET("/exist",Exist)
     router.GET("/find",MyFind)
     router.Run(":8090")
}

func Insert(c *gin.Context){
	user := new(User1)
	user.Name = c.Query("name")
	Age := c.Query("age")
	age, err := strconv.ParseInt(Age,10,64)
	if err != nil{
		fmt.Println(err)
	}
	user.Age = age
	rel, err := engine1.Insert(user)
	if rel == 0 || err != nil {
		c.JSON(200, gin.H{"msg": "添加错误", "err": err, "rel": rel})
	} else {
		c.JSON(200, gin.H{"msg": "添加成功"})
	}
}

func Select(c *gin.Context)  {
	//user := new(User1)
	sql := "select * from user1"
	results, err := engine1.Query(sql)
	if err != nil{
		fmt.Println(err)
	}else{
		c.JSON(200, gin.H{
			"msg": results,
		})
	}
}

func SelectSql(c *gin.Context){
	//user := make([]User1,0)

	var user []User1
	err := engine1.SQL("select * from user1 ").Find(&user)
	if err != nil{
		fmt.Println(err)
	}else{
		c.JSON(200,gin.H{
			"msg": user,
		})
	}
}

func SelectIn(c *gin.Context){
	user := []User1{}
	err := engine1.In("id",1,2,3).Find(&user)
	if err != nil{
		fmt.Println(err)
	}else{
		c.JSON(200, gin.H{
			"msg": user,
		})
	}
}

func SelectCol(c *gin.Context){
	user := User1{}
	//user.Id = 1
	user.Age = 10
	user.Name = "liu"
	// ler,err := engine1.Cols("age","name").Get(&user)
	//if !ler || err != nil{
	//	fmt.Println(err)
	//}else{
	//	c.JSON(200, gin.H{
	//		"msg": user,
	//	})
	//}
	rel, err := engine1.Cols("age", "my_name").ID(1).Update(&user)
	if rel == 0 || err != nil {
		c.JSON(200, gin.H{"msg": "修改错误!", "rel": rel, "err": err, "user": user})
	} else {
		c.JSON(200, gin.H{"mag": "修改成功"})
	}
}

func MyFind(c *gin.Context){
	//users := make([]User1,0)
	//err := engine1.Where("age = ?",10).Find(&users)
	//
	//if err!= nil{
	//	fmt.Println(err)
	//}else{
	//	c.JSON(200, gin.H{
	//		"msg": users,
	//	})
	//}

	var ints []int64
	err := engine1.Table("user1").Cols("id").Find(&ints)
	if err !=nil{
		fmt.Println(err)
	}else{
		c.JSON(200, gin.H{
			"msg": ints,
		})
	}
}

func Exist(c *gin.Context){
	user := &User1{}
	//user.Id = 1
	(*user).Name = "liu"
	has, err := engine1.Exist(user)
	if !has || err != nil{
		fmt.Println(err)
	}else{
		c.JSON(200,gin.H{
			"msg": has,
		})
	}
}

//func Get(c *gin.Context){
//	user := new(User1)
//	has, err := engine1.Id(1).Get(user)
//	if !has || err != nil{
//
//	}else{
//
//	}
//}
