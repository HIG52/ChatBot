package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MainHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"main": "main",
	})
}
