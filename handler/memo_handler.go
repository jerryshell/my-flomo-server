package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jerryshell/my-flomo-server/form"
	"github.com/jerryshell/my-flomo-server/model"
	"github.com/jerryshell/my-flomo-server/service"
	"log"
	"strings"
)

func MemoList(c *gin.Context) {
	memoList := service.MemoList()

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

	token := formData.Token
	log.Println("[MemoForPlugin] token: ", token)
	// TODO: 此处执行 token 解析，获取 userId
	userID := "这里是最后解析出来的userId"

	content := strings.TrimSpace(formData.Content)
	if len(content) == 0 {
		c.JSON(400, gin.H{
			"success": false,
			"message": "内容不能为空",
		})
		return
	}

	memo := model.Memo{
		Content: content,
		UserId:  userID,
	}

	res, err := service.MemoCreate(memo)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "ok",
		"data":    res,
	})
}

// MemoForPlugin 这里是兼容 flomo 生态的接口
func MemoForPlugin(c *gin.Context) {
	// TODO: 此处执行 token 解析，获取 userId
	token := c.Param("token")
	log.Println("[MemoForPlugin] token: ", token)
	userID := "这里是最后解析出来的userId"

	var formData form.MemoCreateForm
	if err := c.ShouldBindJSON(&formData); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	content := strings.TrimSpace(formData.Content)
	if len(content) == 0 {
		c.JSON(400, gin.H{
			"success": false,
			"message": "内容不能为空",
		})
		return
	}

	memo := model.Memo{
		Content: content,
		UserId:  userID,
	}

	res, err := service.MemoCreate(memo)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "ok",
		"data":    res,
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

	memo, err := service.MemoUpdate(formData.ID, formData.Content)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "ok",
		"data":    memo,
	})
}

func MemoDelete(c *gin.Context) {
	id := c.Param("id")

	service.MemoDelete(id)

	c.JSON(200, gin.H{
		"success": true,
		"message": "ok",
	})
}

func SendRandomMemo(c *gin.Context) {
	memo, err := service.SendRandomMemo()
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "ok",
		"data":    memo,
	})
}
