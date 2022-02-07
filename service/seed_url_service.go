package service

import (
	"github.com/jerryshell/my-flomo-server/db"
	"github.com/jerryshell/my-flomo-server/model"
)

type SeedUrlService struct{}

func (SeedUrlService) Page(page uint, size uint) interface{} {
	var m []model.SeedUrl
	db.DB.Offset(int((page - 1) * size)).Limit(int(size)).Find(&m)
	return m
}

func (SeedUrlService) List() interface{} {
	var m []model.SeedUrl
	db.DB.Find(&m)
	return m
}

func (SeedUrlService) Get(id string) (interface{}, error) {
	var m model.SeedUrl
	err := db.DB.Where("id = ?", id).First(&m).Error
	return m, err
}

func (SeedUrlService) Create(i interface{}) {
	db.DB.Create(i)
}

func (SeedUrlService) DeleteByID(id string) {
	db.DB.Where("id = ?", id).Delete(model.SeedUrl{})
}

func (SeedUrlService) Delete(i interface{}) {
	db.DB.Delete(i)
}

func (SeedUrlService) Update(i interface{}) {
	db.DB.Save(i)
}
