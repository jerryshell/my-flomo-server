package store

import (
	"log"

	"github.com/jerryshell/my-flomo/api/db"
	"github.com/jerryshell/my-flomo/api/model"
	"github.com/jerryshell/my-flomo/api/util"
)

func UserListByEmailIsNotNull() ([]model.User, error) {
	var userList []model.User
	err := db.DB.Where("email is not null").Order("created_at desc").Find(&userList).Error
	return userList, err
}

func UserGetByID(id string) (*model.User, error) {
	user := &model.User{}
	if err := db.DB.First(user, id).Error; err != nil {
		log.Println("db.DB.First :: err", err)
		return nil, err
	}
	return user, nil
}

func UserGetByEmail(email string) (*model.User, error) {
	user := &model.User{}
	err := db.DB.Where("email = ?", email).First(user).Error
	return user, err
}

func UserGetByPluginToken(token string) (*model.User, error) {
	user := &model.User{}
	err := db.DB.Where("plugin_token = ?", token).First(user).Error
	return user, err
}

func UserCreate(email string, password string) (*model.User, error) {
	id, err := util.NextIDStr()
	if err != nil {
		log.Println("util.NextIDStr :: err", err)
		return nil, err
	}

	user := &model.User{
		BaseModel: model.BaseModel{
			ID: id,
		},
		Email:    email,
		Password: password,
	}
	err = db.DB.Create(user).Error

	return user, err
}

func UserSave(user *model.User) error {
	return db.DB.Save(user).Error
}
