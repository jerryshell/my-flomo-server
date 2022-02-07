package service

import (
	"github.com/jerryshell/my-flomo-server/db"
	"github.com/jerryshell/my-flomo-server/model"
)

type MentionService struct{}

func (MentionService) Page(page uint, size uint) interface{} {
	var m []model.Mention
	db.DB.Offset(int((page - 1) * size)).Limit(int(size)).Find(&m)
	return m
}

func (MentionService) List() interface{} {
	var m []model.Mention
	db.DB.Find(&m)
	return m
}

func (MentionService) Get(id string) (interface{}, error) {
	var m model.Mention
	if err := db.DB.Where("id = ?", id).First(&m).Error; err != nil {
		return nil, err
	}
	return m, nil
}

func (MentionService) Create(i interface{}) {
	db.DB.Create(i)
}

func (MentionService) DeleteByID(id string) {
	db.DB.Where("id = ?", id).Delete(model.Mention{})
}

func (MentionService) Delete(i interface{}) {
	db.DB.Delete(i)
}

func (MentionService) Update(i interface{}) {
	db.DB.Save(i)
}
