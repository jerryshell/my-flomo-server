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

// CreatePluginSecret 这里是兼容 flomo 生态的接口
func CreatePluginSecret(c *gin.Context) {
	// TODO: 此处执行 token 解析，获取 userId
	token := "这里是从 header 中拿到的 token"
	log.Println("[MemoForPlugin] token: ", token)
	// TODO: 这里是从 token 解析出来的userId
	userID := "2"

	userSecret, err := service.GetSecretByUserId(userID)
	log.Println("[MemoCreateForPlugin][WithUserSecret]", userSecret)
	if err == nil {
		log.Println("[CreatePluginSecret]", userSecret)
		c.JSON(http.StatusOK, result.ErrorWithMessage(userSecret.Secret))
		return
	}
	id, _ := util.NextIDStr()
	Secret := model.Secret{
		BaseModel: model.BaseModel{
			ID: id,
		},
		UserId: userID,
		Secret: uuid.NewV4().String(),
	}

	res := db.DB.Create(&Secret)
	log.Printf("[CreatePluginSecret] 创建 secret 用户: %s,secret: %s 影响行数：%s", userID, Secret, res.RowsAffected)
	c.JSON(http.StatusOK, result.ErrorWithMessage(Secret.Secret))
}

// CreateMemoByPluginSecret 这里是兼容 flomo 生态的接口
func CreateMemoByPluginSecret(c *gin.Context) {
	secret := c.Param("secret")
	log.Println("[CreateMemoByPluginSecret] secret: ", secret)

	if secret == "" {
		log.Println("[CreateMemoByPluginSecret] secret: 空")
		c.JSON(http.StatusOK, result.ErrorWithMessage("secret 为空"))
		return
	}

	userSecret, err := service.GetSecretBySecret(secret)

	if err != nil {
		log.Println("[CreateMemoByPluginSecret][GetSecretBySecret] secret 不存在")
		c.JSON(http.StatusOK, result.ErrorWithMessage("secret 为空"))
		return
	}
	log.Printf("[CreateMemoByPluginSecret][GetSecretBySecret] userId：%s,secret：%s", userSecret.UserId, userSecret.Secret)
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
