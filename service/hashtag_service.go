package service

import (
	"github.com/jerryshell/my-flomo-server/db"
	"github.com/jerryshell/my-flomo-server/model"
)

type HashtagService struct{}

func (HashtagService) Page(page uint, size uint) interface{} {
	var m []model.Hashtag
	db.DB.Offset(int((page - 1) * size)).Limit(int(size)).Find(&m)
	return m
}

func (HashtagService) List() interface{} {
	var m []model.Hashtag
	db.DB.Find(&m)
	return m
}

func (HashtagService) Get(id string) (interface{}, error) {
	var m model.Hashtag
	if err := db.DB.Where("id = ?", id).First(&m).Error; err != nil {
		return nil, err
	}
	return m, nil
}

func (HashtagService) Create(i interface{}) {
	db.DB.Create(i)
}

func (HashtagService) DeleteByID(id string) {
	db.DB.Where("id = ?", id).Delete(model.Hashtag{})
}

func (HashtagService) Delete(i interface{}) {
	db.DB.Delete(i)
}

func (HashtagService) Update(i interface{}) {
	db.DB.Save(i)
}
