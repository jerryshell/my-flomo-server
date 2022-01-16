package service

import (
	"errors"
	"github.com/jerryshell/my-flomo-server/model"
	"github.com/jerryshell/my-flomo-server/store"
	"github.com/jerryshell/my-flomo-server/util"
	"golang.org/x/crypto/bcrypt"
)

func UserListByEmailIsNotNull() ([]model.User, error) {
	return store.UserListByEmailIsNotNull()
}

func UserGetByUsername(username string) (*model.User, error) {
	return store.UserGetByUsername(username)
}

func UserCreate(username string, password string) (*model.User, error) {
	passwordBcrypt, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("密码加密失败")
	}

	id, err := util.NextIDStr()
	if err != nil {
		return nil, err
	}

	user := &model.User{
		BaseModel: model.BaseModel{
			ID: id,
		},
		Username: username,
		Password: string(passwordBcrypt),
	}

	err = store.UserCreate(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func UserUpdateEmail(userID string, email string) (*model.User, error) {
	user, err := store.UserGetByID(userID)
	if err != nil {
		return nil, err
	}
	if user.ID == "" {
		return nil, errors.New("找不到 user，id: " + userID)
	}

	user.Email = email
	err = store.UserSave(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
