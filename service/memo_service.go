package service

import (
	"errors"
	"github.com/jerryshell/my-flomo-server/config"
	"github.com/jerryshell/my-flomo-server/model"
	"github.com/jerryshell/my-flomo-server/store"
	"gopkg.in/gomail.v2"
	"log"
	"math/rand"
	"time"
)

func MemoListByUserID(userID string) ([]model.Memo, error) {
	return store.MemoListByUserID(userID)
}

func MemoCreate(content string, userID string) (*model.Memo, error) {
	return store.MemoCreate(content, userID)
}

func MemoSave(memo *model.Memo) error {
	return store.MemoSave(memo)
}

func MemoUpdate(id string, content string) (*model.Memo, error) {
	memo, err := store.MemoGetByID(id)
	if err != nil {
		return nil, err
	}
	if memo.ID == "" {
		return nil, errors.New("找不到 memo，id: " + id)
	}

	memo.Content = content
	err = store.MemoSave(&memo)

	return &memo, err
}

func MemoDelete(id string) error {
	return store.MemoDeleteByID(id)
}

func MemoGetRandomByUserID(userID string) (*model.Memo, error) {
	memoList, err := MemoListByUserID(userID)
	if err != nil {
		return nil, err
	}
	if len(memoList) == 0 {
		return nil, errors.New("memo 数据为空")
	}

	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(memoList))

	return &memoList[index], nil
}

func MemoDailyReview() error {
	userList, err := UserListByEmailIsNotNull()
	if err != nil {
		return err
	}
	if len(userList) == 0 {
		return errors.New("用户数据为空")
	}

	for _, user := range userList {
		log.Println("MemoDailyReview() user", user)

		memo, err := MemoGetRandomByUserID(user.ID)
		log.Println("MemoDailyReview() memo", memo)
		if err != nil {
			log.Println("MemoDailyReview() err", err)
			continue
		}

		m := gomail.NewMessage()
		m.SetHeader("From", config.Data.SmtpUsername)
		m.SetHeader("To", user.Email)
		m.SetHeader("Subject", config.Data.SmtpSubject)
		m.SetBody("text/plain", memo.Content)

		d := gomail.NewDialer(
			config.Data.SmtpHost,
			config.Data.SmtpPort,
			config.Data.SmtpUsername,
			config.Data.SMTPPassword,
		)
		if err = d.DialAndSend(m); err != nil {
			log.Println("发送失败", err)
			continue
		}
	}

	return nil
}
