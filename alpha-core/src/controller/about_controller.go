package controller

/**
关于栏
*/
import (
	"alpha-core/src/services"
	"alpha-core/src/views"
	"github.com/gin-gonic/gin"
	"net/http"
)

func QueryAllMessage(c *gin.Context) {
	var req views.Page
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	messages, _ := services.AllMessages(req.PageIndex, req.PageSize)
	c.JSON(200, messages)
}

func AddMessage(c *gin.Context) {
	var req views.AddMessageRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ret := services.AddMessage(req)
	c.JSON(200, gin.H{
		"result": ret,
	})
}
