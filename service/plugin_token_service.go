package service

import (
	"github.com/jerryshell/my-flomo-server/model"
	"github.com/jerryshell/my-flomo-server/store"
)

func PluginTokenGetByUserID(userID string) (*model.PluginToken, error) {
	return store.PluginTokenGetByUserID(userID)
}

func PluginTokenGetByToken(token string) (*model.PluginToken, error) {
	return store.PluginTokenGetByToken(token)
}

func PluginTokenCreateByUserID(userID string) (*model.PluginToken, error) {
	// 获取旧插件令牌
	pluginTokenGetByUserID, _ := PluginTokenGetByUserID(userID)

	// 删除旧插件令牌
	if pluginTokenGetByUserID != nil {
		_ = PluginTokenDeleteByID(pluginTokenGetByUserID.ID)
	}

	// 创建新插件令牌
	pluginToken, err := store.PluginTokenCreate(userID)

	return pluginToken, err
}

func PluginTokenDeleteByID(id string) error {
	return store.PluginTokenDeleteByID(id)
}
