package route

import (
	"github.com/gin-gonic/gin"
	"github.com/jerryshell/my-flomo-server/handler"
)

func Setup(app *gin.Engine) {
	app.GET("/memo/list", handler.MemoList)
	app.POST("/memo/create", handler.MemoCreate)
	app.POST("/memo/update", handler.MemoUpdate)
	app.POST("/memo/delete/id/:id", handler.MemoDelete)
}
