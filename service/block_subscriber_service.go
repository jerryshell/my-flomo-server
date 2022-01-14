package service

import (
	"github.com/jerryshell/my-flomo-server/db"
	"github.com/jerryshell/my-flomo-server/model"
)

type BlockSubscriberService struct{}

func (BlockSubscriberService) Page(page uint, size uint) interface{} {
	var m []model.BlockSubscriber
	db.DB.Offset(int((page - 1) * size)).Limit(int(size)).Find(&m)
	return m
}

func (BlockSubscriberService) List() interface{} {
	var m []model.BlockSubscriber
	db.DB.Find(&m)
	return m
}

func (BlockSubscriberService) Get(id string) (interface{}, error) {
	var m model.BlockSubscriber
	if err := db.DB.Where("id = ?", id).First(&m).Error; err != nil {
		return nil, err
	}
	return m, nil
}

func (BlockSubscriberService) Create(i interface{}) {
	db.DB.Create(i)
}

func (BlockSubscriberService) DeleteByID(id string) {
	db.DB.Where("id = ?", id).Delete(model.BlockSubscriber{})
}

func (BlockSubscriberService) Delete(i interface{}) {
	db.DB.Delete(i)
}

func (BlockSubscriberService) Update(i interface{}) {
	db.DB.Save(i)
}
