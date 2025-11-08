package service

import (
	"crypto/rand"
	"errors"
	"log"
	"math/big"

	"github.com/jerryshell/my-flomo/api/model"
	"github.com/jerryshell/my-flomo/api/store"
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

	// 定义字符集：小写字母、大写字母和数字
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	tokenLength := 32
	token := make([]byte, tokenLength)

	// 生成只包含字母和数字的随机令牌
	for i := range token {
		randomIndex, gen_err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if gen_err != nil {
			log.Println("rand.Int :: err", gen_err)
			return nil, errors.New("生成随机令牌失败")
		}
		token[i] = charset[randomIndex.Int64()]
	}

	user.PluginToken = string(token)

	if err = store.UserSave(user); err != nil {
		log.Println("store.UserSave :: err", err)
		return nil, err
	}

	return user, nil
}
