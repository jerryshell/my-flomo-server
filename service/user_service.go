package service

import (
	"encoding/base64"
	"errors"
	"log"

	"github.com/jerryshell/my-flomo-server/model"
	"github.com/jerryshell/my-flomo-server/store"
	"golang.org/x/crypto/bcrypt"
)

func UserListByEmailIsNotNull() ([]model.User, error) {
	return store.UserListByEmailIsNotNull()
}

func UserGetByEmail(email string) (*model.User, error) {
	return store.UserGetByEmail(email)
}

func UserCreate(email string, password string) (*model.User, error) {
	// 对密码进行bcrypt加密
	passwordBcrypt, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("bcrypt.GenerateFromPassword :: err", err)
		return nil, errors.New("密码加密失败")
	}

	user, err := store.UserCreate(email, string(passwordBcrypt))
	if err != nil {
		log.Println("store.UserCreate :: err", err)
		return nil, err
	}
	return user, nil
}

func UserUpdatePassword(userID string, password string) (*model.User, error) {
	user, err := store.UserGetByID(userID)
	if err != nil {
		log.Println("store.UserGetByID :: err", err)
		return nil, err
	}
	if user.ID == "" {
		return nil, errors.New("找不到 user，id: " + userID)
	}

	passwordBcrypt, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("bcrypt.GenerateFromPassword :: err", err)
		return nil, errors.New("密码加密失败")
	}

	user.Password = string(passwordBcrypt)
	if err = store.UserSave(user); err != nil {
		log.Println("store.UserSave :: err", err)
		return nil, err
	}

	return user, nil
}

func UserGetByPluginToken(token string) (*model.User, error) {
	return store.UserGetByPluginToken(token)
}

func UserUpdatePluginToken(userID string) (*model.User, error) {
	user, err := store.UserGetByID(userID)
	if err != nil {
		log.Println("store.UserGetByID :: err", err)
		return nil, err
	}
	if user.ID == "" {
		return nil, errors.New("找不到 user，id: " + userID)
	}

	// 生成新的插件令牌
	user.PluginToken = base64.RawStdEncoding.EncodeToString([]byte(userID))
	
	if err = store.UserSave(user); err != nil {
		log.Println("store.UserSave :: err", err)
		return nil, err
	}

	return user, nil
}
