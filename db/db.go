package db

import (
	"github.com/jerryshell/my-flomo-server/config"
	"github.com/jerryshell/my-flomo-server/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	// 初始化 SQLite
	dsn := config.Data.DSN
	db, _ := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	DB = db

	// 迁移 schema
	if err := db.AutoMigrate(&model.Memo{}); err != nil {
		panic("Memo autoMigrate failed")
	}

	if err := db.AutoMigrate(&model.User{}); err != nil {
		panic("User autoMigrate failed")
	}

	if err := db.AutoMigrate(&model.PluginToken{}); err != nil {
		panic("PluginToken autoMigrate failed")
	}
}
