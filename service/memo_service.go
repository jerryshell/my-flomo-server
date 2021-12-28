package service

import (
	"errors"
	"github.com/jerryshell/my-flomo-server/config"
	"github.com/jerryshell/my-flomo-server/db"
	"github.com/jerryshell/my-flomo-server/model"
	"github.com/jerryshell/my-flomo-server/util"
	"gopkg.in/gomail.v2"
	"math/rand"
	"time"
)

func MemoList() []model.Memo {
	var memoList []model.Memo
	_ = db.DB.Order("created_at desc").Find(&memoList)
	return memoList
}

func MemoSave(memo *model.Memo) error {
	db.DB.Save(memo)
	return nil
}

func MemoCreate(memo model.Memo) error {
	id, err := util.NextIDStr()
	if err != nil {
		return err
	}
	memo.ID = id
	_ = db.DB.Create(memo)

	return nil
}

func MemoUpdate(id string, content string) (*model.Memo, error) {
	memo := model.Memo{}
	_ = db.DB.First(&memo, id)
	if memo.ID == "" {
		return nil, errors.New("找不到 memo，id: " + id)
	}

	memo.Content = content
	_ = db.DB.Save(&memo)

	return &memo, nil
}

func MemoDelete(id string) {
	memo := model.Memo{}
	_ = db.DB.First(&memo, id)
	_ = db.DB.Delete(&memo)
}

func GetRandomMemo() (*model.Memo, error) {
	memoList := MemoList()
	if len(memoList) <= 0 {
		return nil, errors.New("memo 数据为空")
	}
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(memoList))
	return &memoList[index], nil
}

func SendRandomMemo() (*model.Memo, error) {
	memo, err := GetRandomMemo()
	if err != nil {
		return nil, err
	}

	m := gomail.NewMessage()
	m.SetHeader("From", config.Data.SmtpUsername)
	m.SetHeader("To", config.Data.SmtpTo)
	m.SetHeader("Subject", config.Data.SmtpSubject)
	m.SetBody("text/plain", memo.Content)

	d := gomail.NewDialer(
		config.Data.SmtpHost,
		config.Data.SmtpPort,
		config.Data.SmtpUsername,
		config.Data.SmtpPassword,
	)
	if err = d.DialAndSend(m); err != nil {
		return nil, err
	}

	return memo, nil
}
