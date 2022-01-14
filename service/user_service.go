package service

import (
	"errors"
	"github.com/jerryshell/my-flomo-server/db"
	"github.com/jerryshell/my-flomo-server/model"
	"github.com/jerryshell/my-flomo-server/util"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func UserList() []model.User {
	var userList []model.User
	_ = db.DB.Order("created_at desc").Find(&userList)
	return userList
}

func UserGetByUsername(username string) *model.User {
	user := &model.User{}
	db.DB.Where("username = ?", username).First(user)
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

	passwordByte, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Password = string(passwordByte)
	_ = db.DB.Save(&user)

	return &user, nil
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

func UserDeleteById(id string) {
	user := model.User{}
	_ = db.DB.First(&user, id)
	if user.ID == "" {
		return
	}
	_ = db.DB.Delete(&user)
}

type UserService struct{}

func (UserService) Page(page uint, size uint) interface{} {
	var m []model.User
	db.DB.Offset(int((page - 1) * size)).Limit(int(size)).Find(&m)
	return m
}

func (UserService) List() interface{} {
	var m []model.User
	db.DB.Find(&m)
	return m
}

func (UserService) Get(id string) (interface{}, error) {
	var m model.User
	err := db.DB.First(&m, id).Error
	return m, err
}

func (UserService) Create(i interface{}) {
	db.DB.Create(i)
}

func (UserService) DeleteByID(id string) {
	db.DB.Delete(model.User{}, id)
}

func (UserService) Delete(i interface{}) {
	db.DB.Delete(i)
}

func (UserService) Update(i interface{}) {
	db.DB.Save(i)
}
