package controller

import (
	"alpha-core/src/views"
	"github.com/gin-gonic/gin"
)

func GetDataB(c *gin.Context) {
	var b views.StructB
	c.Bind(&b)
	c.JSON(200, gin.H{
		"a": b.NestedStruct,
		"b": b.FieldB,
	})
}
