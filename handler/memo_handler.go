package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jerryshell/my-flomo-server/db"
	"github.com/jerryshell/my-flomo-server/form"
	"github.com/jerryshell/my-flomo-server/model"
	"github.com/jerryshell/my-flomo-server/util"
)

func MemoList(c *gin.Context) {
	var memoList []model.Memo
	_ = db.DB.Find(&memoList)

	c.JSON(200, gin.H{
		"success": true,
		"message": "success",
		"data":    memoList,
	})
}

func MemoCreate(c *gin.Context) {
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
		BaseModel: model.BaseModel{
			ID: id,
		},
		Content: formData.Content,
	}
	_ = db.DB.Create(&memo)

	c.JSON(200, gin.H{
		"success": true,
		"message": "ok",
		"data":    memo,
	})
}
