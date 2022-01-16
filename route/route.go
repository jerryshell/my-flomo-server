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
	app.POST("/user/updateEmail", middleware.JWTMiddleware(), handler.UpdateUserEmail)
	app.GET("/memo/list", middleware.JWTMiddleware(), handler.MemoList)
	app.POST("/memo/create", middleware.JWTMiddleware(), handler.MemoCreate)
	app.POST("/memo/update", middleware.JWTMiddleware(), handler.MemoUpdate)
	app.POST("/memo/delete/id/:id", middleware.JWTMiddleware(), handler.MemoDeleteByID)
	app.GET("/memo/dailyReview", handler.MemoDailyReview)
	app.POST("/upload", middleware.JWTMiddleware(), handler.Upload)
	app.POST("/plugin/createMemo/:pluginToken", handler.PluginTokenCreateMemo)
	app.POST("/plugin/createToken", middleware.JWTMiddleware(), handler.PluginTokenCreate)
	app.GET("/plugin/getToken", middleware.JWTMiddleware(), handler.PluginTokenGet)
}
