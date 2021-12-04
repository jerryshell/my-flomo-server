package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jerryshell/my-flomo-server/route"
)

func main() {
	app := gin.Default()
	// cors
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))
	// route setup
	route.Setup(app)
	// run
	_ = app.Run("0.0.0.0:8080")
}
