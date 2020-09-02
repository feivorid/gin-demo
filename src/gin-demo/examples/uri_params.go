package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Student struct {
	ID   string `uri:"id" binding:"required"`
	Name string `uri:"name" binding:"required"`
}

func main() {
	route := gin.Default()
	route.GET("/:name/:id", func(c *gin.Context) {
		var student Student

		if err := c.ShouldBindUri(&student); err != nil { //注意这里绑定方法是uri
			log.Println(err.Error())
			c.JSON(400, gin.H{"msg": err})
			return
		}

		c.JSON(200, gin.H{"name": student.Name, "uuid": student.ID})
	})

	route.Run(":8080")
}
