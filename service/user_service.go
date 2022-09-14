package service

import (
	"errors"
	"github.com/jerryshell/my-flomo-server/model"
	"github.com/jerryshell/my-flomo-server/store"
	"golang.org/x/crypto/bcrypt"
	"log"
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
		log.Println("bcrypt.GenerateFromPassword :: err", err)
		return nil, errors.New("密码加密失败")
	}

	user, err := store.UserCreate(username, string(passwordBcrypt))
	if err != nil {
		log.Println("store.UserCreate :: err", err)
		return nil, err
	}

	return user, nil
}

func UserUpdateEmail(userID string, email string) (*model.User, error) {
	user, err := store.UserGetByID(userID)
	if err != nil {
		log.Println("store.UserGetByID :: err", err)
		return nil, err
	}
	if user.ID == "" {
		return nil, errors.New("找不到 user，id: " + userID)
	}

	user.Email = email
	if err = store.UserSave(user); err != nil {
		log.Println("store.UserSave :: err", err)
		return nil, err
	}

	return user, nil
}
