package handler

import (
	"net/http"

	"github.com/HIG52/chatBot/service"
	"github.com/gin-gonic/gin"
)

func RoomHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "채팅방",
	})
}

func ChatHandler(c *gin.Context) {

	chat := c.PostForm("chat")

	result := service.ChatBot(chat)

	c.JSON(http.StatusOK, gin.H{
		"ans": result,
	})
}
