package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-basic/uuid"
	"github.com/jerryshell/my-flomo-server/db"
	"github.com/jerryshell/my-flomo-server/form"
	"github.com/jerryshell/my-flomo-server/model"
	"github.com/jerryshell/my-flomo-server/model/gen"
	"github.com/jerryshell/my-flomo-server/result"
	"github.com/jerryshell/my-flomo-server/service"
	"github.com/jerryshell/my-flomo-server/util"
	"log"
	"strings"
)

// CreatePluginSecret 这里是兼容 flomo 生态的接口
func CreatePluginSecret(c *gin.Context) {
	// TODO: 此处执行 token 解析，获取 userId
	token := "这里是从 header 中拿到的 token"
	log.Println("[MemoForPlugin] token: ", token)
	// TODO: 这里是从 token 解析出来的userId
	userID := "testUserID"

	pluginSecretMgr := gen.PluginSecretMgr(db.DB)
	userSecret, _ := pluginSecretMgr.GetFromUserID(userID)
	log.Println("[MemoCreateForPlugin][WithUserSecret]", userSecret)
	if len(userSecret) != 0 {
		log.Println("[MemoCreateForPlugin] secret: 已存在")
		c.JSON(200, result.ErrorWithMessage(userSecret[0].UserSecret))
		return
	}
	id, _ := util.NextIDStr()
	Secret := gen.PluginSecret{
		ID:         id,
		UserID:     userID,
		UserSecret: uuid.New(),
	}

	// TODO：db.DB.Create(Secret) 这里需要写入数据库
	log.Printf("[MemoCreateForPlugin] 创建 secret 用户: %s,secret: %s", userID, Secret)
	c.JSON(200, result.ErrorWithMessage(Secret.UserSecret))
}

// CreateMemoByPluginSecret 这里是兼容 flomo 生态的接口
func CreateMemoByPluginSecret(c *gin.Context) {
	secret := c.Param("secret")
	log.Println("[MemoForPlugin] secret: ", secret)
	if secret == "" {
		log.Println("[MemoForPlugin] secret: 空")
		c.JSON(400, result.ErrorWithMessage("secret 为空"))
		return
	}

	pluginSecretMgr := gen.PluginSecretMgr(db.DB)
	userSecret, _ := pluginSecretMgr.GetFromUserSecret(secret)
	log.Println("[MemoCreateForPlugin][WithUserSecret]", userSecret)
	if len(userSecret) == 0 {
		log.Println("[MemoCreateForPlugin] secret: 不存在")
		c.JSON(400, result.ErrorWithMessage("secret 不存在"))
		return
	}

	userID := userSecret[0].UserID
	var formData form.MemoCreateForm
	if err := c.ShouldBindJSON(&formData); err != nil {
		c.JSON(400, result.ErrorWithMessage(err.Error()))
		return
	}

	content := strings.TrimSpace(formData.Content)
	if len(content) == 0 {
		c.JSON(400, result.ErrorWithMessage("内容不能为空"))
		return
	}

	memo := model.Memo{
		Content: content,
		UserID:  userID,
	}

	err := service.MemoCreate(memo)
	if err != nil {
		c.JSON(400, result.ErrorWithMessage(err.Error()))
		return
	}

	c.JSON(200, result.SuccessWithData(memo))
}
