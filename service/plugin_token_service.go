package service

import (
	"errors"
	"github.com/jerryshell/my-flomo-server/db"
	"github.com/jerryshell/my-flomo-server/model"
	"github.com/jerryshell/my-flomo-server/util"
)

func List() []model.PluginToken {
	var list []model.PluginToken
	_ = db.DB.Order("created_at desc").Find(&list)
	return list
}

func GetByUserId(userId string) (*model.PluginToken, error) {
	var token model.PluginToken
	if err := db.DB.Where("user_id = ?", userId).First(&token).Error; err != nil {
		return nil, err
	}
	return &token, nil
}

func GetByToken(token string) (*model.PluginToken, error) {
	var tokenModel model.PluginToken
	if err := db.DB.Where("token = ?", token).First(&tokenModel).Error; err != nil {
		return nil, err
	}
	return &tokenModel, nil
}

func Save(pluginToken *model.PluginToken) error {
	if err := db.DB.Save(pluginToken).Error; err != nil {
		return err
	}
	return nil
}

func Create(userId string, token string) (*model.PluginToken, error) {
	id, err := util.NextIDStr()
	if err != nil {
		return nil, err
	}

	pluginToken := &model.PluginToken{
		BaseModel: model.BaseModel{
			ID: id,
		},
		UserId: userId,
		Token:  token,
	}

	_ = db.DB.Create(pluginToken)

	return pluginToken, nil
}

func Update(id string, token string) (*model.PluginToken, error) {
	pluginToken := model.PluginToken{}

	_ = db.DB.First(&pluginToken, id)
	if pluginToken.ID == "" {
		return nil, errors.New("找不到 pluginToken，id: " + id)
	}

	pluginToken.Token = token
	_ = db.DB.Save(&pluginToken)

	return &pluginToken, nil
}

func Delete(id string) {
	pluginToken := model.PluginToken{}
	_ = db.DB.First(&pluginToken, id)
	if pluginToken.ID == "" {
		return
	}
	_ = db.DB.Delete(&pluginToken)
}
