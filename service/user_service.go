package service

import (
	"errors"
	"github.com/jerryshell/my-flomo-server/db"
	"github.com/jerryshell/my-flomo-server/model"
	"github.com/jerryshell/my-flomo-server/util"
	"log"
)

func UserList() []model.User {
	var userList []model.User
	_ = db.DB.Order("created_at desc").Find(&userList)
	return userList
}

func GetUserByUsername(username string) model.User {
	var user model.User
	db.DB.Where("username = ?", username).First(&user)
	return user
}

func UserSave(user *model.User) error {
	db.DB.Save(user)
	return nil
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

func UserUpdate(userID string, password string) (*model.User, error) {
	user := model.User{}
	_ = db.DB.First(&user, userID)
	if user.ID == "" {
		return nil, errors.New("找不到 user，id: " + userID)
	}

	user.Password = password
	_ = db.DB.Save(&user)

	return &user, nil
}

func UserDelete(id string) {
	user := model.User{}
	_ = db.DB.First(&user, id)
	_ = db.DB.Delete(&user)
}
