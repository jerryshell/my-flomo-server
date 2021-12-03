package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sony/sonyflake"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var sf *sonyflake.Sonyflake
var DB *gorm.DB
var Memos []Memo

func GetMachineID() (uint16, error) {
	return uint16(1), nil
}

func init() {
	// 初始化雪花id
	var st = sonyflake.Settings{
		MachineID: GetMachineID,
	}
	sf = sonyflake.NewSonyflake(st)
	if sf == nil {
		panic("sonyflake not created")
	}

	// 初始化 MySQL
	dsn := "root:toor@tcp(devenv.d8s.fun:3306)/flomo?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	DB = db
}

type MemoCreateForm struct {
	Content string `json:"content" required:"true"`
}

type Memo struct {
	gorm.Model
	Content string `json:"content"`
}

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	r.GET("/memo/list", func(c *gin.Context) {

		_ = DB.Find(&Memos)
		fmt.Println(Memos)
		c.JSON(200, gin.H{
			"success": true,
			"message": "success",
			"data":    Memos,
		})
	})
	r.POST("/memo/create", func(c *gin.Context) {
		var form MemoCreateForm
		if err := c.ShouldBindJSON(&form); err != nil {
			c.JSON(400, gin.H{
				"success": false,
				"message": err.Error(),
			})
			return
		}
		id, err := sf.NextID()
		if err != nil {
			c.JSON(500, gin.H{
				"success": false,
				"message": err.Error(),
			})
			return
		}
		memo := Memo{
			Content: form.Content,
		}
		memo.ID = uint(id)
		//memoList = append(memoList, memo)
		_ = DB.Create(&memo)
		c.JSON(200, gin.H{
			"success": true,
			"message": id,
		})
	})
	err := r.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
