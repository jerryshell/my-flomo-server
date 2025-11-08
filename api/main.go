package main

import (
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jerryshell/my-flomo/api/config"
	"github.com/jerryshell/my-flomo/api/route"
	"github.com/jerryshell/my-flomo/api/service"
	"github.com/jerryshell/my-flomo/api/util"
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
	logger := util.NewLogger("cron")

	c := cron.New()
	_, err := c.AddFunc(config.Data.CronSpec, func() {
		logger.Info("cron job started", util.StringField("cron_spec", config.Data.CronSpec))
		err := service.MemoDailyReview()
		if err != nil {
			logger.Error("memo daily review failed", util.ErrorField(err))
			return
		}
		logger.Info("cron job completed successfully")
	})
	if err != nil {
		logger.Fatal("failed to add cron function", util.ErrorField(err))
	}
	c.Start()

	logger.Info("cron scheduler started")
}
