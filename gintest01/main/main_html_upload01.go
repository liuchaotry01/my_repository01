package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
				)

func main(){
	router := gin.Default()

	router.LoadHTMLGlob("E:/gin_html_test/templates/*")
	router.GET("/upload", func(c *gin.Context) {
		//name := c.PostForm("name")
		//fmt.Println(name)
		//file, header, err := c.Request.FormFile("upload")
		//if err != nil {
		//	c.String(http.StatusBadRequest, "Bad request")
		//	return
		//}
		//filename := header.Filename
		//
		//fmt.Println(file, err, filename)
		//
		//out, err := os.Create(filename)
		//if err != nil {
		//	log.Fatal(err)
		//}
		//defer out.Close()
		//_, err = io.Copy(out, file)
		//if err != nil {
		//	log.Fatal(err)
		//}
		//c.String(http.StatusCreated, "upload successful")
		c.HTML(http.StatusOK, "upload.html", gin.H{})
	})

	router.Run(":8090")
}
