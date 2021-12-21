package db

import (
	"github.com/jerryshell/my-flomo-server/config"
	"github.com/jerryshell/my-flomo-server/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	// 初始化 MySQL
	dsn := config.Map.DSN
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	DB = db

	// 迁移 schema
	err := db.AutoMigrate(&model.Memo{})
	if err != nil {
		panic("Memo autoMigrate failed")
	}
}
