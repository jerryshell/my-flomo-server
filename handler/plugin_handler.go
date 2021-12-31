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
	"log"
	"net/http"
	"strings"
)

// CreatePluginToken 这里是兼容 flomo 生态的接口
func CreatePluginToken(c *gin.Context) {
	// TODO: 此处执行 token 解析，获取 userId
	token := "这里是从 header 中拿到的 token"
	log.Println("[MemoForPlugin] token: ", token)
	// TODO: 这里是从 token 解析出来的userId
	userID := "2"

	userSecret, err := service.GetByUserId(userID)
	if err == nil {
		c.JSON(http.StatusOK, result.ErrorWithMessage(userSecret.Token))
		return
	}

	id, _ := util.NextIDStr()
	pluginToken := model.PluginToken{
		BaseModel: model.BaseModel{
			ID: id,
		},
		UserId: userID,
		Token:  uuid.NewV4().String(),
	}

	_ = db.DB.Create(&pluginToken)

	c.JSON(http.StatusOK, result.ErrorWithMessage(pluginToken.Token))
}

// CreateMemoByPluginToken 这里是兼容 flomo 生态的接口
func CreateMemoByPluginToken(c *gin.Context) {
	pluginToken := c.Param("pluginToken")

	if pluginToken == "" {
		c.JSON(http.StatusOK, result.ErrorWithMessage("pluginToken 为空"))
		return
	}

	userSecret, err := service.GetByToken(pluginToken)

	if err != nil {
		log.Println("[CreateMemoByPluginToken][GetByToken] pluginToken 不存在")
		c.JSON(http.StatusOK, result.ErrorWithMessage("pluginToken 为空"))
		return
	}
	log.Printf("[CreateMemoByPluginToken][GetByToken] userId：%s,pluginToken：%s", userSecret.UserId, userSecret.Token)
	var formData form.MemoCreateForm
	if err := c.ShouldBindJSON(&formData); err != nil {
		c.JSON(http.StatusOK, result.ErrorWithMessage(err.Error()))
		return
	}

	content := strings.TrimSpace(formData.Content)
	if len(content) == 0 {
		c.JSON(http.StatusOK, result.ErrorWithMessage("内容不能为空"))
		return
	}

	res, err := service.MemoCreate(content, userSecret.UserId)

	if err != nil {
		c.JSON(http.StatusOK, result.ErrorWithMessage(err.Error()))
		return
	}

	c.JSON(http.StatusOK, result.SuccessWithData(res))
}
