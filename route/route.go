package route

import (
	"github.com/gin-gonic/gin"
	"github.com/jerryshell/my-flomo-server/handler"
)

func Setup(app *gin.Engine) {
	app.POST("/auth/login", handler.Login)
	app.POST("/auth/verifyToken/token/:token", handler.VerifyToken)
	app.POST("/auth/register/", handler.Register)
	app.GET("/memo/list", handler.MemoList)
	app.POST("/memo/create", handler.MemoCreate)
	app.POST("/memo/update", handler.MemoUpdate)
	app.POST("/memo/delete/id/:id", handler.MemoDelete)
	app.GET("/memo/sendRandomMemo", handler.SendRandomMemo)
	app.POST("/upload", handler.Upload)
	app.POST("/plugin/create/:secret", handler.CreateMemoByPluginToken)
	app.POST("/plugin/createToken", handler.CreatePluginToken)
}
