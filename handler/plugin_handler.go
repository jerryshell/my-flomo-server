package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jerryshell/my-flomo-server/db"
	"github.com/jerryshell/my-flomo-server/form"
	"github.com/jerryshell/my-flomo-server/model"
	"github.com/jerryshell/my-flomo-server/result"
	"github.com/jerryshell/my-flomo-server/service"
	"github.com/jerryshell/my-flomo-server/util"
	"github.com/satori/go.uuid"
	"net/http"
	"strings"
)

// CreatePluginToken 这里是兼容 flomo 生态的接口
func CreatePluginToken(c *gin.Context) {
	user := c.MustGet("user").(*model.User)

	pluginToken, err := service.PluginTokenGetByUserId(user.ID)
	if err == nil {
		c.JSON(http.StatusOK, result.ErrorWithMessage(pluginToken.Token))
		return
	}

	id, _ := util.NextIDStr()
	pluginToken = &model.PluginToken{
		BaseModel: model.BaseModel{
			ID: id,
		},
		UserId: user.ID,
		Token:  uuid.NewV4().String(),
	}

	_ = db.DB.Create(pluginToken)

	c.JSON(http.StatusOK, result.ErrorWithMessage(pluginToken.Token))
}

// CreateMemoByPluginToken 这里是兼容 flomo 生态的接口
func CreateMemoByPluginToken(c *gin.Context) {
	token := c.Param("pluginToken")

	if token == "" {
		c.JSON(http.StatusOK, result.ErrorWithMessage("pluginToken 为空"))
		return
	}

	pluginToken, err := service.PluginTokenGetByToken(token)
	if err != nil {
		c.JSON(http.StatusOK, result.ErrorWithMessage("token 无效"))
		return
	}

	formData := &form.MemoCreateForm{}
	if err := c.ShouldBindJSON(&formData); err != nil {
		c.JSON(http.StatusOK, result.ErrorWithMessage(err.Error()))
		return
	}

	content := strings.TrimSpace(formData.Content)
	if len(content) == 0 {
		c.JSON(http.StatusOK, result.ErrorWithMessage("内容不能为空"))
		return
	}

	res, err := service.MemoCreate(content, pluginToken.UserId)

	if err != nil {
		c.JSON(http.StatusOK, result.ErrorWithMessage(err.Error()))
		return
	}

	c.JSON(http.StatusOK, result.SuccessWithData(res))
}
