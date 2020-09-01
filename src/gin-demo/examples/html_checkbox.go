package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type myForm struct {
	Colors []string `form:"colors[]"`
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/colors", func(c *gin.Context) {
		c.HTML(http.StatusOK, "checkbox/color.tmpl", gin.H{
			"title": "Select Colors",
		})
	})

	r.POST("/colors", func(c *gin.Context) {
		var fakeForm myForm
		c.ShouldBind(&fakeForm)
		c.JSON(200, gin.H{"color": fakeForm.Colors})
	})

	r.Run()
}
