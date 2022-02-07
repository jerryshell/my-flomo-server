package service

import (
	"github.com/jerryshell/my-flomo-server/db"
	"github.com/jerryshell/my-flomo-server/model"
)

type SeedHashtagService struct{}

func (SeedHashtagService) Page(page uint, size uint) interface{} {
	var m []model.SeedHashtag
	db.DB.Offset(int((page - 1) * size)).Limit(int(size)).Find(&m)
	return m
}

func (SeedHashtagService) List() interface{} {
	var m []model.SeedHashtag
	db.DB.Find(&m)
	return m
}

func (SeedHashtagService) Get(id string) (interface{}, error) {
	var m model.SeedHashtag
	err := db.DB.Where("id = ?", id).First(&m).Error
	return m, err
}

func (SeedHashtagService) Create(i interface{}) {
	db.DB.Create(i)
}

func (SeedHashtagService) DeleteByID(id string) {
	db.DB.Where("id = ?", id).Delete(model.SeedHashtag{})
}

func (SeedHashtagService) Delete(i interface{}) {
	db.DB.Delete(i)
}

func (SeedHashtagService) Update(i interface{}) {
	db.DB.Save(i)
}
