package service

import (
	"github.com/jerryshell/my-flomo-server/db"
	"github.com/jerryshell/my-flomo-server/model"
)

type SeedMentionService struct{}

func (SeedMentionService) Page(page uint, size uint) interface{} {
	var m []model.SeedMention
	db.DB.Offset(int((page - 1) * size)).Limit(int(size)).Find(&m)
	return m
}

func (SeedMentionService) List() interface{} {
	var m []model.SeedMention
	db.DB.Find(&m)
	return m
}

func (SeedMentionService) Get(id string) (interface{}, error) {
	var m model.SeedMention
	err := db.DB.Where("id = ?", id).First(&m).Error
	return m, err
}

func (SeedMentionService) Create(i interface{}) {
	db.DB.Create(i)
}

func (SeedMentionService) DeleteByID(id string) {
	db.DB.Where("id = ?", id).Delete(model.SeedMention{})
}

func (SeedMentionService) Delete(i interface{}) {
	db.DB.Delete(i)
}

func (SeedMentionService) Update(i interface{}) {
	db.DB.Save(i)
}
