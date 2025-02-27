package main

import (
	"log"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jerryshell/my-flomo-server/config"
	"github.com/jerryshell/my-flomo-server/route"
	"github.com/jerryshell/my-flomo-server/service"
	"github.com/robfig/cron/v3"
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
	_ = app.Run("0.0.0.0:" + strconv.Itoa(config.Data.Port))
}

func initCron() {
	c := cron.New()
	_, err := c.AddFunc(config.Data.CronSpec, func() {
		log.Println("cron job: " + config.Data.CronSpec)
		err := service.MemoDailyReview()
		if err != nil {
			log.Println("service.MemoDailyReview :: err", err)
			return
		}
	})
	if err != nil {
		log.Fatal("c.AddFunc :: err", err)
	}
	c.Start()
}
