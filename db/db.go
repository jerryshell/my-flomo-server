package db

import (
	"github.com/jerryshell/my-flomo-server/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	// 初始化 MySQL
	dsn := "root:toor@tcp(devenv.d8s.fun:3306)/my_flomo?charset=utf8mb4&parseTime=True&loc=Asia%2fShanghai"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	DB = db

	// 迁移 schema
	err := db.AutoMigrate(&model.Memo{})
	if err != nil {
		panic("Memo autoMigrate failed")
	}
}
