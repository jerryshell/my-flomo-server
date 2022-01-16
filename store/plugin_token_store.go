package store

import (
	"github.com/jerryshell/my-flomo-server/db"
	"github.com/jerryshell/my-flomo-server/model"
)

func PluginTokenGetByUserID(userID string) (*model.PluginToken, error) {
	token := &model.PluginToken{}
	err := db.DB.Where("user_id = ?", userID).First(token).Error
	return token, err
}

func PluginTokenGetByToken(token string) (*model.PluginToken, error) {
	tokenModel := &model.PluginToken{}
	err := db.DB.Where("token = ?", token).First(&tokenModel).Error
	return tokenModel, err
}

func PluginTokenDeleteByID(id string) error {
	pluginToken := &model.PluginToken{}
	return db.DB.Delete(pluginToken, id).Error
}

func PluginTokenCreate(token *model.PluginToken) error {
	return db.DB.Create(token).Error
}
