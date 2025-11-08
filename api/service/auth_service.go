package service

import (
	"errors"
	"log"

	"github.com/jerryshell/my-flomo/api/model"
)

func Register(email, password string) (*model.User, error) {
	user, err := UserGetByEmail(email)
	if user.ID != "" {
		log.Println("UserGetByEmail :: err", err)
		return nil, errors.New("用户已存在")
	}

	user, err = UserCreate(email, password)
	if err != nil {
		log.Println("UserCreate :: err", err)
		return nil, errors.New("创建用户失败")
	}

	return user, nil
}
