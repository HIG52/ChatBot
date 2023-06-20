package router

import (
	"log"

	"github.com/HIG52/chatBot/api/handler"
	"github.com/gin-gonic/gin"
)

func RunServer() {
	r := gin.Default()

	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	//메인
	r.GET("/", handler.MainHandler)

	//메뉴
	r.GET("/menu", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"menu1": "menu1",
		})
	})

	//게시판
	board := r.Group("/board")
	//공지사항
	noticeController(board)

	//채팅
	chat := r.Group("/chat")
	chatController(chat)

	r.Run(":3000") // 서버가 실행 되고 0.0.0.0:8080 에서 요청을 기다립니다.

}

// 공지사항
func noticeController(board *gin.RouterGroup) {
	//List
	board.GET("/notice", handler.ListHandler)
	//Write
	board.POST("/notice", handler.WriteHandler)
}

// 채팅방
func chatController(chat *gin.RouterGroup) {
	//채팅방
	chat.GET("/room", handler.RoomHandler)
	//채팅
	chat.POST("/post", handler.ChatHandler)
}
