package service

import (
	"github.com/jerryshell/my-flomo-server/db"
	"github.com/jerryshell/my-flomo-server/model"
)

type UserBlockService struct{}

func (UserBlockService) Page(page uint, size uint) interface{} {
	var m []model.UserBlock
	db.DB.Offset(int((page - 1) * size)).Limit(int(size)).Find(&m)
	return m
}

func (UserBlockService) List() interface{} {
	var m []model.UserBlock
	db.DB.Find(&m)
	return m
}

func (UserBlockService) Get(id string) (interface{}, error) {
	var m model.UserBlock
	err := db.DB.Where("id = ?", id).First(&m).Error
	return m, err
}

func (UserBlockService) Create(i interface{}) {
	db.DB.Create(i)
}

func (UserBlockService) DeleteByID(id string) {
	db.DB.Where("id = ?", id).Delete(model.UserBlock{})
}

func (UserBlockService) Delete(i interface{}) {
	db.DB.Delete(i)
}

func (UserBlockService) Update(i interface{}) {
	db.DB.Save(i)
}
