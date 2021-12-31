package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jerryshell/my-flomo-server/form"
	"github.com/jerryshell/my-flomo-server/model"
	"github.com/jerryshell/my-flomo-server/result"
	"github.com/jerryshell/my-flomo-server/service"
	"net/http"
	"strings"
)

func MemoList(c *gin.Context) {
	memoList := service.MemoList()

	c.JSON(http.StatusOK, result.SuccessWithData(memoList))
}

func MemoCreate(c *gin.Context) {
	formData := &form.MemoCreateForm{}
	if err := c.ShouldBindJSON(formData); err != nil {
		c.JSON(http.StatusOK, result.ErrorWithMessage(err.Error()))
		return
	}

	content := strings.TrimSpace(formData.Content)
	if len(content) == 0 {
		c.JSON(http.StatusOK, result.ErrorWithMessage("内容不能为空"))
		return
	}

	user := c.MustGet("user").(*model.User)

	memo, err := service.MemoCreate(content, user.ID)
	if err != nil {
		c.JSON(http.StatusOK, result.ErrorWithMessage(err.Error()))
		return
	}

	c.JSON(http.StatusOK, result.SuccessWithData(memo))
}

func MemoUpdate(c *gin.Context) {
	formData := &form.MemoUpdateForm{}
	if err := c.ShouldBindJSON(&formData); err != nil {
		c.JSON(http.StatusOK, result.ErrorWithMessage(err.Error()))
		return
	}

	memo, err := service.MemoUpdate(formData.ID, formData.Content)
	if err != nil {
		c.JSON(http.StatusOK, result.ErrorWithMessage(err.Error()))
		return
	}

	c.JSON(http.StatusOK, result.SuccessWithData(memo))
}

func MemoDelete(c *gin.Context) {
	id := c.Param("id")

	service.MemoDelete(id)

	c.JSON(http.StatusOK, result.Success())
}

func SendRandomMemo(c *gin.Context) {
	memo, err := service.MemoSendRandom()
	if err != nil {
		c.JSON(http.StatusOK, result.ErrorWithMessage(err.Error()))
		return
	}

	c.JSON(http.StatusOK, result.SuccessWithData(memo))
}
