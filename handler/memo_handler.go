package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jerryshell/my-flomo-server/form"
	"github.com/jerryshell/my-flomo-server/result"
	"github.com/jerryshell/my-flomo-server/service"
	"log"
	"net/http"
	"strings"
)

func MemoList(c *gin.Context) {
	memoList := service.MemoList()

	c.JSON(200, result.SuccessWithData(memoList))
}

func MemoCreate(c *gin.Context) {
	//userID := "1"

	var formData form.MemoCreateForm
	if err := c.ShouldBindJSON(&formData); err != nil {
		c.JSON(http.StatusBadRequest, result.ErrorWithMessage(err.Error()))
		return
	}

	content := strings.TrimSpace(formData.Content)
	if len(content) == 0 {
		c.JSON(http.StatusBadRequest, result.ErrorWithMessage("内容不能为空"))
		return
	}

	memo, err := service.MemoCreate(content)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.ErrorWithMessage(err.Error()))
		return
	}

	c.JSON(200, result.SuccessWithData(memo))
}

func MemoUpdate(c *gin.Context) {
	// TODO: 在此处校验 token
	token := "这里是从 header 中拿到的 token"
	log.Println("[MemoForPlugin] token: ", token)

	var formData form.MemoUpdateForm
	if err := c.ShouldBindJSON(&formData); err != nil {
		c.JSON(http.StatusBadRequest, result.ErrorWithMessage(err.Error()))
		return
	}

	memo, err := service.MemoUpdate(formData.ID, formData.Content)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.ErrorWithMessage(err.Error()))
		return
	}

	c.JSON(200, result.SuccessWithData(memo))
}

func MemoDelete(c *gin.Context) {
	id := c.Param("id")

	service.MemoDelete(id)

	c.JSON(200, result.Success())
}

func SendRandomMemo(c *gin.Context) {
	memo, err := service.SendRandomMemo()
	if err != nil {
		c.JSON(http.StatusBadRequest, result.ErrorWithMessage(err.Error()))
		return
	}

	c.JSON(200, result.SuccessWithData(memo))
}
