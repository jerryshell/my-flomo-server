package service

import (
	"errors"
	"github.com/jerryshell/my-flomo-server/db"
	"github.com/jerryshell/my-flomo-server/model"
	"github.com/jerryshell/my-flomo-server/util"
	"log"
)

func UserListByEmailIsNotNull() []model.User {
	var userList []model.User
	_ = db.DB.Where("email is not null").Order("created_at desc").Find(&userList)
	return userList
}

func UserGetByUsername(username string) *model.User {
	user := &model.User{}
	db.DB.Where("username = ?", username).First(user)
	return user
}

func UserCreate(username string, password string) (*model.User, error) {
	id, err := util.NextIDStr()
	if err != nil {
		return nil, err
	}

	user := &model.User{
		BaseModel: model.BaseModel{
			ID: id,
		},
		Username: username,
		Password: password,
	}
	log.Println("user", user)
	_ = db.DB.Create(user)

	return user, nil
}

func UserUpdateEmail(userID string, email string) (*model.User, error) {
	user := model.User{}
	_ = db.DB.First(&user, userID)
	if user.ID == "" {
		return nil, errors.New("找不到 user，id: " + userID)
	}

	user.Email = email
	_ = db.DB.Save(&user)

	return &user, nil
}
