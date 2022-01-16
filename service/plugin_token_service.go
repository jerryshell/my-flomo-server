package service

import (
	"github.com/jerryshell/my-flomo-server/model"
	"github.com/jerryshell/my-flomo-server/store"
	"github.com/jerryshell/my-flomo-server/util"
	"github.com/satori/go.uuid"
	"log"
)

func PluginTokenGetByUserID(userID string) (*model.PluginToken, error) {
	return store.PluginTokenGetByUserID(userID)
}

func PluginTokenGetByToken(token string) (*model.PluginToken, error) {
	return store.PluginTokenGetByToken(token)
}

func PluginTokenCreateByUserID(userID string) (*model.PluginToken, error) {
	// 获取旧令牌
	pluginTokenGetByUserID, _ := PluginTokenGetByUserID(userID)

	// 删除旧插件令牌
	if pluginTokenGetByUserID != nil {
		err := PluginTokenDeleteByID(pluginTokenGetByUserID.ID)
		if err != nil {
			log.Println("删除旧插件令牌失败", err)
		}
	}

	// 创建新插件令牌
	id, err := util.NextIDStr()
	if err != nil {
		return nil, err
	}
	pluginToken := &model.PluginToken{
		BaseModel: model.BaseModel{
			ID: id,
		},
		UserID: userID,
		Token:  uuid.NewV4().String(),
	}
	err = store.PluginTokenCreate(pluginToken)

	return pluginToken, err
}

func PluginTokenDeleteByID(id string) error {
	return store.PluginTokenDeleteByID(id)
}
