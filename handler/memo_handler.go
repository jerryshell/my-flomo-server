package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jerryshell/my-flomo-server/form"
	"github.com/jerryshell/my-flomo-server/model"
	"github.com/jerryshell/my-flomo-server/result"
	"github.com/jerryshell/my-flomo-server/service"
	"log"
	"net/http"
	"strings"
)

func MemoList(c *gin.Context) {
	user := c.MustGet("user").(*model.User)
	memoList, err := service.MemoListByUserID(user.ID)
	if err != nil {
		log.Println("store.MemoListByUserID :: err", err)
		c.JSON(http.StatusInternalServerError, result.ErrorWithMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, result.SuccessWithData(memoList))
}

func MemoCreate(c *gin.Context) {
	formData := &form.MemoCreateForm{}
	if err := c.ShouldBindJSON(formData); err != nil {
		log.Println("c.ShouldBindJSON :: err", err)
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
		log.Println("service.MemoCreate :: err", err)
		c.JSON(http.StatusOK, result.ErrorWithMessage(err.Error()))
		return
	}

	c.JSON(http.StatusOK, result.SuccessWithData(memo))
}

func MemoUpdate(c *gin.Context) {
	formData := &form.MemoUpdateForm{}
	if err := c.ShouldBindJSON(&formData); err != nil {
		log.Println("c.ShouldBindJSON :: err", err)
		c.JSON(http.StatusOK, result.ErrorWithMessage(err.Error()))
		return
	}

	memo, err := service.MemoUpdate(formData.ID, formData.Content)
	if err != nil {
		log.Println("service.MemoUpdate :: err", err)
		c.JSON(http.StatusOK, result.ErrorWithMessage(err.Error()))
		return
	}

	c.JSON(http.StatusOK, result.SuccessWithData(memo))
}

func MemoDeleteByID(c *gin.Context) {
	id := c.Param("id")
	err := service.MemoDeleteByID(id)
	if err != nil {
		log.Println("service.MemoDeleteByID :: err", err)
		c.JSON(http.StatusOK, result.ErrorWithMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, result.Success())
}

func MemoDailyReview(c *gin.Context) {
	err := service.MemoDailyReview()
	if err != nil {
		log.Println("service.MemoDailyReview :: err", err)
		c.JSON(http.StatusOK, result.ErrorWithMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, result.Success())
}
