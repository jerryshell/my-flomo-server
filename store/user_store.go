package store

import (
	"github.com/jerryshell/my-flomo-server/db"
	"github.com/jerryshell/my-flomo-server/model"
)

func UserListByEmailIsNotNull() ([]model.User, error) {
	var userList []model.User
	err := db.DB.Where("email is not null").Order("created_at desc").Find(&userList).Error
	return userList, err
}

func UserGetByID(id string) (*model.User, error) {
	user := &model.User{}
	err := db.DB.First(user, id).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func UserGetByUsername(username string) (*model.User, error) {
	user := &model.User{}
	err := db.DB.Where("username = ?", username).First(user).Error
	return user, err
}

func UserCreate(user *model.User) error {
	return db.DB.Create(user).Error
}

func UserSave(user *model.User) error {
	return db.DB.Save(user).Error
}
