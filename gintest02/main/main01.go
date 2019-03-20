package main

import (
	"database/sql"
	"log"
	"github.com/gin-gonic/gin"
	"net/http"
	_"github.com/go-sql-driver/mysql"
)

func main(){
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/mysql?parseTime=true")
	if err != nil{
		log.Fatalln(err)
	}
	defer db.Close()

	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)

	if err := db.Ping();err != nil{
		log.Fatalln(err)
	}

	router := gin.Default()
	router.GET("/",func(c *gin.Context){
		c.String(http.StatusOK,"it works")
	})

	router.Run(":8090")
}
