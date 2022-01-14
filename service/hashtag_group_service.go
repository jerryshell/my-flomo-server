package service

import (
	"github.com/jerryshell/my-flomo-server/db"
	"github.com/jerryshell/my-flomo-server/model"
)

type HashtagGroupService struct{}

func (HashtagGroupService) Page(page uint, size uint) interface{} {
	var m []model.HashtagGroup
	db.DB.Offset(int((page - 1) * size)).Limit(int(size)).Find(&m)
	return m
}

func (HashtagGroupService) List() interface{} {
	var m []model.HashtagGroup
	db.DB.Find(&m)
	return m
}

func (HashtagGroupService) Get(id string) (interface{}, error) {
	var m model.HashtagGroup
	err := db.DB.Where("id = ?", id).First(&m).Error
	return m, err
}

func (HashtagGroupService) Create(i interface{}) {
	db.DB.Create(i)
}

func (HashtagGroupService) DeleteByID(id string) {
	db.DB.Where("id = ?", id).Delete(model.HashtagGroup{})
}

func (HashtagGroupService) Delete(i interface{}) {
	db.DB.Delete(i)
}

func (HashtagGroupService) Update(i interface{}) {
	db.DB.Save(i)
}
