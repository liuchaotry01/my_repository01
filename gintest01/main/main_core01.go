package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"time"
	"log"
)

type User struct {
	Username string `form:"username" json:"username" binding:"required"`
	Passwd   string `form:"passwd" json:"passwd" bdinding:"required"`
	Age      int    `form:"age" json:"age"`
}

func main(){
	router := gin.Default()

	router.GET("/render",MultiForm)
	router.GET("redict/baidu",MyRedirect)

	v1 := router.Group("/v1")
	v1.GET("/login",RoutGroup1)
	v2 := router.Group("/v2")
	v2.GET("/login",RoutGroup2)

	router.Use(MiddleWare())
	{
		router.GET("/middleware",MyMiddleWare)
	}
	router.GET("/before",MiddleWare(),Mymw1)
	//authorized := router.Group("/" , MyMiddleWare())
	//authorized := router.Group("/")
	//authorized.Use(MyMiddleWare())
	//{
	//	authorized.POST("login",LoginEndpoint)
	//}

	router.GET("/auth/signin",MiddleWareTest01)
	router.GET("/home", AuthMiddleWare(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "home"})
	})

    MyTest(router)

	router.Run(":8090")
}

func MyTest(router *gin.Engine){
	router.GET("/sync", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		log.Println("Done! in path" + c.Request.URL.Path)
	})

	router.GET("/async", func(c *gin.Context) {
		cCp := c.Copy()
		go func() {
			time.Sleep(5 * time.Second)
			log.Println("Done! in path" + cCp.Request.URL.Path)
		}()
	})
}

func MultiForm(c *gin.Context){
	contentType := c.DefaultQuery("content_type","json")
	if contentType == "json"{
		c.JSON(http.StatusOK,gin.H{
			"user": "liuchao",
			"passwd": "123",
		})
	}else if contentType == "xml"{
		c.XML(http.StatusOK,gin.H{
			"user": "ping",
			"passwd": "456",
		})
	}
}

func MyRedirect(c *gin.Context){
	c.Redirect(http.StatusMovedPermanently,"http://baidu.com")
}

func RoutGroup1(c *gin.Context){
	c.String(http.StatusOK,"v1 login")
}

func RoutGroup2(c *gin.Context){
	c.String(http.StatusOK,"v2 login")
}

func MiddleWare() gin.HandlerFunc{
	return func(c *gin.Context){
		fmt.Println("before biddleware")
		c.Set("request","client_request")
		c.Next()
		fmt.Println("before middleware")
	}
}

func MyMiddleWare(c *gin.Context){
	request := c.MustGet("request").(string)
	req,_ := c.Get("request")
	c.JSON(http.StatusOK,gin.H{
		"middle_request": request,
		"request": req,
	})
}

func Mymw1(c *gin.Context){
	request := c.MustGet("request").(string)
	c.JSON(http.StatusOK,gin.H{
		"middle_request": request,
	})
}

//func LoginEndpoint(c *gin.Context){
//	c.String(http.StatusOK, "loginEndpoint")
//}

func MiddleWareTest01(c *gin.Context){
	cookie := &http.Cookie{
		Name: "session_id",
		Value: "123",
		Path: "/",
		HttpOnly: true,
	}
	http.SetCookie(c.Writer,cookie)
	c.String(http.StatusOK,"Login successful")
}

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		if cookie, err := c.Request.Cookie("session_id"); err == nil {
			value := cookie.Value
			fmt.Println(value)
			if value == "123" {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		c.Abort()
		return
	}
}