package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "공지사항 게시판",
	})
}

func WriteHandler(c *gin.Context) {
	id := c.Query("id")
	page := c.DefaultQuery("page", "0")
	name := c.PostForm("name")
	contents := c.PostForm("message")

	fmt.Printf("id: %s; page: %s; name: %s; contents: %s", id, page, name, contents)
	c.JSON(http.StatusOK, gin.H{
		"id":       id,
		"page":     page,
		"name":     name,
		"contents": contents,
	})
}
