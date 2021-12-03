package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

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
			"data": []gin.H{
				{
					"id":         "3",
					"content":    "2016年，WhatsApp 的用户超过10亿，但是只有50个工程师。每个小团队由1到3名工程师组成，拥有很大的自主权。\n-- https://www.quastor.org/p/how-whatsapp-scaled-to-1-billion",
					"createTime": "2021-12-01 14:29:24",
				},
				{
					"id":         "2",
					"content":    "一个可运行的复杂系统，总是从一个简单系统演变而来的。似乎可以因此推断：从头开始设计一个复杂系统，永远不会奏效，必须从一个简单系统开始设计。\n-- https://www.ivanmontilla.com/blog/galls-law-and-how-i-ignored-it",
					"createTime": "2021-10-25 17:51:25",
				},
				{
					"id":         "1",
					"content":    "切勿交浅言深",
					"createTime": "2021-10-06 17:49:11",
				},
			},
		})
	})
	err := r.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
