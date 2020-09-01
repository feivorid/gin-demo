package main

import "github.com/gin-gonic/gin"

//基本结构体
type StructA struct {
	FieldA string `form:"field_a"`
}

//嵌套结构体
type StructB struct {
	NestStruct StructA
	Field      string `form:"field_b"`
}

//嵌套结构体指针的结构体
type StructC struct {
	NestStructPointer *StructA
	FiledC            string `form:"field_c"`
}

//嵌套匿名结构体的结构体
type StructD struct {
	NestAnonyStruct struct {
		FieldX string `form:"field_x"`
	}
	FieldD string `form:"field_d"`
}

func GetDataB(c *gin.Context) {
	var b StructB

	c.Bind(&b)

	c.JSON(200, gin.H{
		"a": b.NestStruct,
		"b": b.Field,
	})
}

func GetDataC(c *gin.Context) {
	var b StructC

	c.Bind(&b)

	c.JSON(200, gin.H{
		"a": b.NestStructPointer,
		"b": b.FiledC,
	})
}

func GetDataD(c *gin.Context) {
	var d StructD

	c.Bind(&d)

	c.JSON(200, gin.H{
		"a": d.NestAnonyStruct,
		"b": d.FieldD,
	})
}

func main() {
	r := gin.Default()
	r.GET("/get/b", GetDataB)
	r.GET("/get/c", GetDataC)
	r.GET("/get/d", GetDataD)

	r.Run()
}
