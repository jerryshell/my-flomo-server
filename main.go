package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jerryshell/my-flomo-server/route"
	"github.com/robfig/cron/v3"
	"log"
)

func main() {
	// cron
	initCron()

	// gin app
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

func initCron() {
	c := cron.New(cron.WithSeconds())
	_, err := c.AddFunc("* * * * * *", func() {
		log.Println("cron test")
	})
	if err != nil {
		log.Fatal(err)
	}
	c.Start()
}
