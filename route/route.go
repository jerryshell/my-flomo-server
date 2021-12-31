package route

import (
	"github.com/gin-gonic/gin"
	"github.com/jerryshell/my-flomo-server/handler"
	"github.com/jerryshell/my-flomo-server/middleware"
)

func Setup(app *gin.Engine) {
	app.POST("/auth/login", handler.LoginOrRegister)
	app.POST("/auth/verifyToken/token/:token", handler.VerifyToken)
	app.POST("/auth/register/", handler.Register)
	app.GET("/memo/list", middleware.JwtMiddleware(), handler.MemoList)
	app.POST("/memo/create", middleware.JwtMiddleware(), handler.MemoCreate)
	app.POST("/memo/update", middleware.JwtMiddleware(), handler.MemoUpdate)
	app.POST("/memo/delete/id/:id", middleware.JwtMiddleware(), handler.MemoDelete)
	app.GET("/memo/sendRandomMemo", handler.SendRandomMemo)
	app.POST("/upload", middleware.JwtMiddleware(), handler.Upload)
	app.POST("/plugin/createMemo/:pluginToken", middleware.JwtMiddleware(), handler.CreateMemoByPluginToken)
	app.POST("/plugin/createToken", middleware.JwtMiddleware(), handler.CreatePluginToken)
	app.GET("/plugin/getToken", middleware.JwtMiddleware(), handler.GetPluginToken)
}
