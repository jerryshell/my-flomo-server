package service

import (
	"errors"
	"github.com/jerryshell/my-flomo-server/config"
	"github.com/jerryshell/my-flomo-server/db"
	"github.com/jerryshell/my-flomo-server/model"
	"github.com/jerryshell/my-flomo-server/util"
	"gopkg.in/gomail.v2"
)

func MemoList() *[]model.Memo {
	var memoList []model.Memo
	_ = db.DB.Order("created_at desc").Find(&memoList)
	return &memoList
}

func MemoCreate(content string) (*model.Memo, error) {
	id, err := util.NextIDStr()
	if err != nil {
		return nil, err
	}

	memo := model.Memo{
		BaseModel: model.BaseModel{
			ID: id,
		},
		Content: content,
	}
	_ = db.DB.Create(&memo)

	return &memo, nil
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

func SendRandomMemo() error {
	// TODO 随机选择一个 Memo
	smtpContent := "test memo"

	m := gomail.NewMessage()
	m.SetHeader("From", config.Data.SmtpUsername)
	m.SetHeader("To", config.Data.SmtpTo)
	m.SetHeader("Subject", config.Data.SmtpSubject)
	m.SetBody("text/html", smtpContent)

	d := gomail.NewDialer(
		config.Data.SmtpHost,
		config.Data.SmtpPort,
		config.Data.SmtpUsername,
		config.Data.SmtpPassword,
	)
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
