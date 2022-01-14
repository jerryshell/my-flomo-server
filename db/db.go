package db

import (
	"github.com/jerryshell/my-flomo-server/config"
	"github.com/jerryshell/my-flomo-server/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	// 初始化 MySQL
	dsn := config.Data.DSN
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB = db

	modelList := []interface{}{
		&model.Block{},
		&model.BlockSubscriber{},
		&model.Hashtag{},
		&model.HashtagGroup{},
		&model.Memo{},
		&model.Mention{},
		&model.PluginToken{},
		&model.SeedHashtag{},
		&model.SeedMention{},
		&model.SeedUrl{},
		&model.Url{},
		&model.User{},
		&model.UserBlock{},
	}
	for _, newModel := range modelList {
		err := db.AutoMigrate(newModel)
		if err != nil {
			panic("Memo autoMigrate failed")
		}
	}
}
