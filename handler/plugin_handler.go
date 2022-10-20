package handler

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jerryshell/my-flomo-server/form"
	"github.com/jerryshell/my-flomo-server/model"
	"github.com/jerryshell/my-flomo-server/result"
	"github.com/jerryshell/my-flomo-server/service"
)

func PluginTokenGet(c *gin.Context) {
	user := c.MustGet("user").(*model.User)

	pluginToken, err := service.PluginTokenGetByUserID(user.ID)
	if pluginToken.ID == "" {
		log.Println("service.PluginTokenGetByUserID :: err", err)
		c.JSON(http.StatusOK, result.ErrorWithMessage("当前没有插件令牌，请重新生成"))
		return
	}

	c.JSON(http.StatusOK, result.SuccessWithData(pluginToken.Token))
}

// PluginTokenCreate 这里是兼容 flomo 生态的接口
func PluginTokenCreate(c *gin.Context) {
	user := c.MustGet("user").(*model.User)

	pluginToken, err := service.PluginTokenCreateByUserID(user.ID)
	if err != nil {
		log.Println("service.PluginTokenCreateByUserID :: err", err)
		c.JSON(http.StatusOK, result.ErrorWithMessage(err.Error()))
		return
	}

	c.JSON(http.StatusOK, result.SuccessWithData(pluginToken.Token))
}

// PluginTokenCreateMemo 这里是兼容 flomo 生态的接口
func PluginTokenCreateMemo(c *gin.Context) {
	token := c.Param("pluginToken")
	if token == "" {
		c.JSON(http.StatusOK, result.BaseResult{
			Code:    -1,
			Success: false,
			Message: "pluginToken 为空",
		})
		return
	}

	pluginToken, err := service.PluginTokenGetByToken(token)
	if err != nil {
		log.Println("service.PluginTokenGetByToken :: err", err)
		c.JSON(http.StatusOK, result.BaseResult{
			Code:    -1,
			Success: false,
			Message: "token 无效",
		})
		return
	}

	formData := &form.MemoCreateForm{}
	if err := c.ShouldBindJSON(&formData); err != nil {
		log.Println("c.ShouldBindJSON :: err", err)
		c.JSON(http.StatusOK, result.BaseResult{
			Code:    -1,
			Success: false,
			Message: err.Error(),
		})
		return
	}

	content := strings.TrimSpace(formData.Content)
	if len(content) == 0 {
		c.JSON(http.StatusOK, result.BaseResult{
			Code:    -1,
			Success: false,
			Message: "内容不能为空",
		})
		return
	}

	memo, err := service.MemoCreate(content, pluginToken.UserID)
	if err != nil {
		log.Println("service.MemoCreate :: err", err)
		c.JSON(http.StatusOK, result.BaseResult{
			Code:    -1,
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result.BaseResult{
		Code:    0,
		Success: true,
		Message: "已记录",
		Data:    memo,
	})
}

// PluginTokenRandomMemo 这里是 my-flomo 随机一条接口
func PluginTokenRandomMemo(c *gin.Context) {
	token := c.Param("pluginToken")
	if token == "" {
		c.JSON(http.StatusOK, result.BaseResult{
			Code:    -1,
			Success: false,
			Message: "pluginToken 为空",
		})
		return
	}

	pluginToken, err := service.PluginTokenGetByToken(token)
	if err != nil {
		log.Println("service.PluginTokenGetByToken :: err", err)
		c.JSON(http.StatusOK, result.BaseResult{
			Code:    -1,
			Success: false,
			Message: "token 无效",
		})
		return
	}

	memo, err := service.MemoGetRandomByUserID(pluginToken.UserID)
	if err != nil {
		log.Println("service.MemoGetRandomByUserID :: err", err)
		c.JSON(http.StatusOK, result.BaseResult{
			Code:    -1,
			Success: false,
			Message: "获取失败",
		})
		return
	}

	c.JSON(http.StatusOK, result.BaseResult{
		Code:    0,
		Success: true,
		Message: "获取成功",
		Data:    memo,
	})
}
