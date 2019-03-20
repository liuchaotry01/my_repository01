package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
)

func main(){
	router := gin.Default()

	router.PUT("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")
		fmt.Printf("id: %s; page: %s; name: %s; message: %s \n", id, page, name, message)
		c.JSON(http.StatusOK, gin.H{
			"status_code": http.StatusOK,

			"id": id,
			"page": page,
			"name": name,
			"message": message,
		})
	})

	router.Run(":8090")
}
