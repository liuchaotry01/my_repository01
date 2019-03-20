package main

import (
	"log"
	"os"
	"io"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()

	router.POST("/multi/upload", func(c *gin.Context) {
		err := c.Request.ParseMultipartForm(200000)
		if err != nil {
			log.Fatal(err)
		}

		formdata := c.Request.MultipartForm

		files := formdata.File["upload"]
		for i, _ := range files {
			file, err := files[i].Open()
			defer file.Close()
			if err != nil {
				log.Fatal(err)
			}

			out, err := os.Create(files[i].Filename)

			defer out.Close()

			if err != nil {
				log.Fatal(err)
			}

			_, err = io.Copy(out, file)

			if err != nil {
				log.Fatal(err)
			}

			c.String(http.StatusCreated, "upload successful")

		}

	})

	router.Run(":8090")
}
