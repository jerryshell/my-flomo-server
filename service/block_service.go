package service

import (
	"github.com/jerryshell/my-flomo-server/db"
	"github.com/jerryshell/my-flomo-server/model"
)

type BlockService struct{}

func (BlockService) Page(page uint, size uint) interface{} {
	var m []model.Block
	db.DB.Offset(int((page - 1) * size)).Limit(int(size)).Find(&m)
	return m
}

func (BlockService) List() interface{} {
	var m []model.Block
	db.DB.Find(&m)
	return m
}

func (BlockService) Get(id string) (interface{}, error) {
	var m model.Block
	if err := db.DB.Where("id = ?", id).First(&m).Error; err != nil {
		return nil, err
	}
	return m, nil
}

func (BlockService) Create(i interface{}) {
	db.DB.Create(i)
}

func (BlockService) DeleteByID(id string) {
	db.DB.Where("id = ?", id).Delete(model.Block{})
}

func (BlockService) Delete(i interface{}) {
	db.DB.Delete(i)
}

func (BlockService) Update(i interface{}) {
	db.DB.Save(i)
}
