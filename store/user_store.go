package store

import (
	"log"

	"github.com/jerryshell/my-flomo-server/db"
	"github.com/jerryshell/my-flomo-server/model"
	"github.com/jerryshell/my-flomo-server/util"
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

func UserGetByUsername(username string) (*model.User, error) {
	user := &model.User{}
	err := db.DB.Where("username = ?", username).First(user).Error
	return user, err
}

func UserCreate(username string, password string) (*model.User, error) {
	id, err := util.NextIDStr()
	if err != nil {
		log.Println("util.NextIDStr :: err", err)
		return nil, err
	}

	user := &model.User{
		BaseModel: model.BaseModel{
			ID: id,
		},
		Username: username,
		Password: password,
	}
	err = db.DB.Create(user).Error

	return user, err
}

func UserSave(user *model.User) error {
	return db.DB.Save(user).Error
}
