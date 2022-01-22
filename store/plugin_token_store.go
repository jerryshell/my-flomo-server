package store

import (
	"github.com/jerryshell/my-flomo-server/db"
	"github.com/jerryshell/my-flomo-server/model"
	"github.com/jerryshell/my-flomo-server/util"
	"github.com/satori/go.uuid"
	"log"
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

func PluginTokenCreate(userID string) (*model.PluginToken, error) {
	id, err := util.NextIDStr()
	if err != nil {
		log.Println("util.NextIDStr :: err", err)
		return nil, err
	}

	pluginToken := &model.PluginToken{
		BaseModel: model.BaseModel{
			ID: id,
		},
		UserID: userID,
		Token:  uuid.NewV4().String(),
	}
	err = db.DB.Create(pluginToken).Error

	return pluginToken, err
}
