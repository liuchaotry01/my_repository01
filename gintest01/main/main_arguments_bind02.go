package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"log"
	"net/http"
)

type User1 struct{
	Username string `form:"username" json:"username" binding:"required"`
	Passwd   string `form:"passwd" json:"passwd" bdinding:"required"`
	Age      int    `form:"age" json:"age"`
}


func main(){

	router := gin.Default()

	router.POST("/login", func(c *gin.Context) {
		var user User1

		err := c.Bind(&user)
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"username":   user.Username,
			"passwd":     user.Passwd,
			"age":        user.Age,
		})

	})

	router.Run(":8090")
}