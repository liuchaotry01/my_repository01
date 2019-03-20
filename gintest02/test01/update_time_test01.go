package main

import (
	"github.com/go-xorm/xorm"
	"log"
	"time"
	_"github.com/go-sql-driver/mysql"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"io"
	"strings"
)

type TimeTest struct{
	Id int64    `xorm:"notnull pk id" json:"id"`
	Name string `xorm:"name" json:"name"`
	UpdateTime  time.Time  `xorm:"updated update_time" json:"update_time"`
}

var Engine2 *xorm.Engine

func init(){
	var err error
	Engine2, err = xorm.NewEngine("mysql", "root:liuchao@tcp(localhost:3306)/lcctest01?charset=utf8")
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
	if err := Engine2.Sync(new(TimeTest)); err != nil {
		log.Fatal("数据表同步失败:", err)
	}
}


func main(){
   router := gin.Default()
   router.POST("/upload",UploadPicture)
   router.Run(":8090")
}

func UploadPicture(c *gin.Context) {
	name := c.PostForm("name")
	fmt.Println(name)
	file,header,err := c.Request.FormFile("upload")
	if err !=nil{
		c.String(400,"Bad requeat")
		return
	}
	filename := header.Filename
	fmt.Println(file,err,filename)
	out,err := os.Create(filename)
	if err != nil{
		log.Fatal(err)
	}
	defer out.Close()
	_,err = io.Copy(out,file)
	if err != nil {
		log.Fatal(err)
	}
	PictureName := strings.Split(filename,".")
	symbol := strings.ToLower(PictureName[len(PictureName)-1])
	if symbol != "jpg" && symbol != "png" && symbol != "gif" {
		c.String(400,"please upload fine file!")
		return
	}
	c.String(200,"upload picture successful")

}