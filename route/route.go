package route

import (
	"github.com/gin-gonic/gin"
	"github.com/jerryshell/my-flomo-server/db"
	"github.com/jerryshell/my-flomo-server/form"
	"github.com/jerryshell/my-flomo-server/model"
	"github.com/jerryshell/my-flomo-server/util"
)

func Setup(app *gin.Engine) {
	app.GET("/memo/list", func(c *gin.Context) {
		var memo []model.Memo
		_ = db.DB.Find(&memo)
		c.JSON(200, gin.H{
			"success": true,
			"message": "success",
			"data":    memo,
		})
	})
	app.POST("/memo/create", func(c *gin.Context) {
		var formData form.MemoCreateForm
		if err := c.ShouldBindJSON(&formData); err != nil {
			c.JSON(400, gin.H{
				"success": false,
				"message": err.Error(),
			})
			return
		}
		id, err := util.NextID()
		if err != nil {
			c.JSON(500, gin.H{
				"success": false,
				"message": err.Error(),
			})
			return
		}
		memo := model.Memo{
			Content: formData.Content,
		}
		memo.ID = id
		_ = db.DB.Create(&memo)
		c.JSON(200, gin.H{
			"success": true,
			"message": id,
		})
	})
}
