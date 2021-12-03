package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sony/sonyflake"
	"time"
)

var sf *sonyflake.Sonyflake

func GetMachineID() (uint16, error) {
	return uint16(1), nil
}

func init() {
	var st = sonyflake.Settings{
		MachineID: GetMachineID,
	}
	sf = sonyflake.NewSonyflake(st)
	if sf == nil {
		panic("sonyflake not created")
	}
}

var memoList = []Memo{}

type MemoCreateForm struct {
	Content string `json:"content" required:"true"`
}

type Memo struct {
	Id         uint64 `json:"id"`
	Content    string `json:"content"`
	CreateTime string `json:"createTime"`
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
		c.JSON(200, gin.H{
			"success": true,
			"message": "success",
			"data":    memoList,
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
			Id:         id,
			Content:    form.Content,
			CreateTime: time.Now().Format("2006-01-02 15:04:05"),
		}
		memoList = append(memoList, memo)
		c.JSON(200, gin.H{
			"success": true,
			"message": "success",
		})
	})
	err := r.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
