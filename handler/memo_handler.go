package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jerryshell/my-flomo-server/config"
	"github.com/jerryshell/my-flomo-server/db"
	"github.com/jerryshell/my-flomo-server/form"
	"github.com/jerryshell/my-flomo-server/model"
	"github.com/jerryshell/my-flomo-server/util"
	"gopkg.in/gomail.v2"
)

func MemoList(c *gin.Context) {
	var memoList []model.Memo
	_ = db.DB.Order("created_at desc").Find(&memoList)

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

	id, err := util.NextIDStr()
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

func MemoUpdate(c *gin.Context) {
	var formData form.MemoUpdateForm
	if err := c.ShouldBindJSON(&formData); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	memo := model.Memo{}
	_ = db.DB.First(&memo, formData.ID)
	if memo.ID == "" {
		c.JSON(404, gin.H{
			"success": false,
			"message": "not found",
		})
		return
	}

	memo.Content = formData.Content
	_ = db.DB.Save(&memo)

	c.JSON(200, gin.H{
		"success": true,
		"message": "ok",
		"data":    memo,
	})
}

func MemoDelete(c *gin.Context) {
	id := c.Param("id")
	memo := model.Memo{}
	_ = db.DB.First(&memo, id)

	_ = db.DB.Delete(&memo)

	c.JSON(200, gin.H{
		"success": true,
		"message": "ok",
	})
}

func SendRandomMemo(c *gin.Context) {
	// TODO 随机选择一个 Memo
	smtpContent := "test memo"

	m := gomail.NewMessage()
	m.SetHeader("From", config.Data.SmtpUsername)
	m.SetHeader("To", config.Data.SmtpTo)
	m.SetHeader("Subject", config.Data.SmtpSubject)
	m.SetBody("text/html", smtpContent)

	d := gomail.NewDialer(
		config.Data.SmtpHost,
		config.Data.SmtpPort,
		config.Data.SmtpUsername,
		config.Data.SmtpPassword,
	)
	if err := d.DialAndSend(m); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "ok",
	})
}
