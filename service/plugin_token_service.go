package service

import (
	"github.com/jerryshell/my-flomo-server/db"
	"github.com/jerryshell/my-flomo-server/model"
)

func PluginTokenGetByUserId(userId string) *model.PluginToken {
	token := &model.PluginToken{}
	db.DB.Where("user_id = ?", userId).First(token)
	return token
}

func PluginTokenGetByToken(token string) (*model.PluginToken, error) {
	var tokenModel model.PluginToken
	if err := db.DB.Where("token = ?", token).First(&tokenModel).Error; err != nil {
		return nil, err
	}
	return &tokenModel, nil
}

func PluginTokenDeleteById(id string) {
	pluginToken := model.PluginToken{}
	_ = db.DB.First(&pluginToken, id)
	if pluginToken.ID == "" {
		return
	}
	_ = db.DB.Delete(&pluginToken)
}
