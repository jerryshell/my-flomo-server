package service

import (
	"github.com/jerryshell/my-flomo-server/db"
	"github.com/jerryshell/my-flomo-server/model"
)

type UrlService struct{}

func (UrlService) Page(page uint, size uint) interface{} {
	var m []model.Url
	db.DB.Offset(int((page - 1) * size)).Limit(int(size)).Find(&m)
	return m
}

func (UrlService) List() interface{} {
	var m []model.Url
	db.DB.Find(&m)
	return m
}

func (UrlService) Get(id string) (interface{}, error) {
	var m model.Url
	err := db.DB.Where("id = ?", id).First(&m).Error
	return m, err
}

func (UrlService) Create(i interface{}) {
	db.DB.Create(i)
}

func (UrlService) DeleteByID(id string) {
	db.DB.Where("id = ?", id).Delete(model.Url{})
}

func (UrlService) Delete(i interface{}) {
	db.DB.Delete(i)
}

func (UrlService) Update(i interface{}) {
	db.DB.Save(i)
}
