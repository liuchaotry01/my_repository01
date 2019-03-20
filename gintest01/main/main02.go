package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK,"Hello %s", name)
	})
	r.Run("127.0.0.1:8080") // listen and serve on 0.0.0.0:8080
}

