package service

import (
	"errors"
	"github.com/jerryshell/my-flomo-server/model"
	"log"
)

func Register(username, password string) (*model.User, error) {
	user, err := UserGetByUsername(username)
	if user.ID != "" {
		log.Println("UserGetByUsername :: err", err)
		return nil, errors.New("用户已存在")
	}

	user, err = UserCreate(username, password)
	if err != nil {
		log.Println("UserCreate :: err", err)
		return nil, errors.New("创建用户失败")
	}

	return user, nil
}
