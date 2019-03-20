package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"github.com/go-xorm/xorm"
	"log"
	_"github.com/go-sql-driver/mysql"
)

var X *xorm.Engine

type User struct {
	Id int `xorm:"pk autoincr"`
	Name string `xorm:"unique"`
	Age int64
}


func init() {
	var err error
	X, err = xorm.NewEngine("mysql", "root:liuchao11@tcp(localhost:3306)/mysql?charset=utf8")
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
	if err := X.Sync(new(User)); err != nil {
		log.Fatal("数据表同步失败:", err)
	}
	//X.Ping()
}

func main()  {
	router := gin.Default()
	router.GET("/insert",insert)
	router.GET("/updates",updates)
	router.GET("find",find)
	router.GET("/get",get)
	router.GET("/delete",delete)
	router.Run(":8090")
}

func insert(c *gin.Context){
	name := c.Query("name")
	if name == ""{
		c.JSON(200,gin.H{
			"msg": "name不得为空",
		})
		return
	}

	age := c.Query("age")
	if age == ""{
		c.JSON(200,gin.H{
			"msg": "age不能为空",
		})
		return
	}

	Age, _ := strconv.ParseInt(age,10,0)
	user := User{}
	user.Name = name
	user.Age = Age
	rel, err := X.Insert(user)

	if rel == 0 || err != nil {
		c.JSON(200, gin.H{"msg": "添加错误", "err": err, "rel": rel})
	} else {
		c.JSON(200, gin.H{"msg": "添加成功"})
	}
}

func get(c *gin.Context){
	id := c.Query("id")
	if id == ""{
		c.JSON(200,gin.H{
			"msg": "id不能为空",
		})
		return
	}
	ids, _ := strconv.ParseInt(id,10,64)
	user := User{}
	rel, err := X.Where("id = ?", ids).Get(user)

	if !rel || err != nil{
		c.JSON(200,gin.H{"msg": "查询错误"})
	}else{
		c.JSON(200, gin.H{"user": user})
	}
}

func find(c *gin.Context){
	users := make(map[int64]User)
	err := X.Find(&users)
	if err != nil{
		c.JSON(200,gin.H{"msg": err})
	}
	c.JSON(200, gin.H{"msg": users})
}

func updates(c *gin.Context){
	id := c.Query("id")
	if id == "" {
		c.JSON(200, gin.H{"msg": "id1不得为空!", "id": id})
		return
	}
	ids, _ := strconv.ParseInt(id, 10, 64)

	name := c.Query("name")
	if name == "" {
		c.JSON(200, gin.H{"msg": "name不得为空!"})
		return
	}
	age := c.Query("age")
	if age == ""{
		c.JSON(200, gin.H{"msg": "age不能为空"})
		return
	}
	Age, _ := strconv.ParseInt(age,10,64)
	user := User{}
	user.Name = name
	user.Age = Age
	rel, err := X.Id(ids).Update(user)

	if rel == 0 || err != nil {
		c.JSON(200, gin.H{"msg": "修改错误!", "rel": rel, "err": err, "user": user})
	} else {
		c.JSON(200, gin.H{"mag": "修改成功"})
	}
}


func delete(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(200, gin.H{"msg": "id1不得为空!", "id": id})
		return
	}
	//string转化int64
	ids, _ := strconv.ParseInt(id, 10, 64)
	//删除
	user := User{}
	rel, err := X.Id(ids).Delete(user)

	if rel == 0 || err != nil {
		c.JSON(200, gin.H{"msg": "删除错误!", "rel": rel, "err": err, "user": user})
	} else {
		c.JSON(200, gin.H{"mag": "删除成功"})
	}
}


