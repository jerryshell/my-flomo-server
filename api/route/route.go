package route

import (
	"github.com/gin-gonic/gin"
	"github.com/jerryshell/my-flomo-server/handler"
	"github.com/jerryshell/my-flomo-server/middleware"
)

func Setup(app *gin.Engine) {
	app.GET("/health", handler.Health)

	authGroup := app.Group("/auth")
	{
		authGroup.POST("/login", handler.LoginOrRegister)
		authGroup.POST("/verifyToken/token/:token", handler.VerifyToken)
		authGroup.POST("/register/", handler.Register)
	}

	memoGroup := app.Group("/memo")
	{
		memoGroup.GET("/list", middleware.JWTMiddleware(), handler.MemoList)
		memoGroup.POST("/create", middleware.JWTMiddleware(), handler.MemoCreate)
		memoGroup.POST("/update", middleware.JWTMiddleware(), handler.MemoUpdate)
		memoGroup.POST("/delete/id/:id", middleware.JWTMiddleware(), handler.MemoDeleteByID)
		memoGroup.GET("/dailyReview", handler.MemoDailyReview)
	}

	pluginGroup := app.Group("/plugin")
	{
		pluginGroup.GET("/getToken", middleware.JWTMiddleware(), handler.PluginTokenGet)
		pluginGroup.POST("/createToken", middleware.JWTMiddleware(), handler.PluginTokenCreate)
		pluginGroup.POST("/createMemo/:pluginToken", handler.PluginTokenCreateMemo)
		pluginGroup.GET("/randomMemo/:pluginToken", handler.PluginTokenRandomMemo)
	}

	userGroup := app.Group("/user")
	{
		userGroup.POST("/updatePassword", middleware.JWTMiddleware(), handler.UpdateUserPassword)
	}

	app.POST("/upload", middleware.JWTMiddleware(), handler.Upload)
	app.POST("/deleteMyAccount", middleware.JWTMiddleware(), handler.DeleteMyAccount)
	app.GET("/csvExport/token/:token", handler.CsvExport)
	app.POST("/csvImport", middleware.JWTMiddleware(), handler.CsvImport)
}
