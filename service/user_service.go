package service

import (
	"errors"
	"github.com/jerryshell/my-flomo-server/model"
	"github.com/jerryshell/my-flomo-server/store"
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

	user, err := store.UserCreate(username, string(passwordBcrypt))
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
