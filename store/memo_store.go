package store

import (
	"log"
	"time"

	"github.com/jerryshell/my-flomo-server/db"
	"github.com/jerryshell/my-flomo-server/model"
	"github.com/jerryshell/my-flomo-server/util"
)

func MemoListByUserID(userID string) ([]model.Memo, error) {
	var memoList []model.Memo
	err := db.DB.Order("created_at desc").Where("user_id = ?", userID).Find(&memoList).Error
	return memoList, err
}

func MemoGetByID(id string) (model.Memo, error) {
	var memo model.Memo
	err := db.DB.Where("id = ?", id).First(&memo).Error
	return memo, err
}

func MemoCreate(content string, userID string) (*model.Memo, error) {
	id, err := util.NextIDStr()
	if err != nil {
		log.Println("util.NextIDStr :: err", err)
		return nil, err
	}

	memo := &model.Memo{
		BaseModel: model.BaseModel{
			ID: id,
		},
		Content: content,
		UserID:  userID,
	}
	err = db.DB.Create(memo).Error

	return memo, err
}

func MemoCreateByTime(content string, userID string, createdAt time.Time) (*model.Memo, error) {
	id, err := util.NextIDStr()
	if err != nil {
		log.Println("util.NextIDStr :: err", err)
		return nil, err
	}

	memo := &model.Memo{
		BaseModel: model.BaseModel{
			ID:        id,
			CreatedAt: createdAt,
		},
		Content: content,
		UserID:  userID,
	}
	err = db.DB.Create(memo).Error

	return memo, err
}

func MemoSave(memo *model.Memo) error {
	return db.DB.Save(memo).Error
}

func MemoDeleteByID(id string) error {
	return db.DB.Delete(&model.Memo{BaseModel: model.BaseModel{ID: id}}).Error
}
