package main

import (
	"log"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jerryshell/my-flomo/api/config"
	"github.com/jerryshell/my-flomo/api/route"
	"github.com/jerryshell/my-flomo/api/service"
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
	_ = app.Run(config.Data.Host + ":" + strconv.Itoa(config.Data.Port))
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
