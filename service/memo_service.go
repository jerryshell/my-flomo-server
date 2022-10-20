package service

import (
	"errors"
	"log"
	"math/rand"
	"time"

	"github.com/jerryshell/my-flomo-server/config"
	"github.com/jerryshell/my-flomo-server/model"
	"github.com/jerryshell/my-flomo-server/store"
	"gopkg.in/gomail.v2"
)

func MemoListByUserID(userID string) ([]model.Memo, error) {
	return store.MemoListByUserID(userID)
}

func MemoCreate(content string, userID string) (*model.Memo, error) {
	return store.MemoCreate(content, userID)
}

func MemoCreateByTime(content string, userID string, createdAt time.Time) (*model.Memo, error) {
	return store.MemoCreateByTime(content, userID, createdAt)
}

func MemoSave(memo *model.Memo) error {
	return store.MemoSave(memo)
}

func MemoUpdate(id string, content string) (*model.Memo, error) {
	memo, err := store.MemoGetByID(id)
	if err != nil {
		log.Println("store.MemoGetByID :: err", err)
		return nil, err
	}
	if memo.ID == "" {
		return nil, errors.New("找不到 memo，id: " + id)
	}

	memo.Content = content
	err = store.MemoSave(&memo)

	return &memo, err
}

func MemoDeleteByID(id string) error {
	return store.MemoDeleteByID(id)
}

func MemoGetRandomByUserID(userID string) (*model.Memo, error) {
	memoList, err := MemoListByUserID(userID)
	if err != nil {
		log.Println("MemoListByUserID :: err", err)
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
		log.Println("UserListByEmailIsNotNull :: err", err)
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
			log.Println("MemoGetRandomByUserID :: err", err)
			continue
		}

		message := gomail.NewMessage()
		message.SetHeader("From", config.Data.SmtpUsername)
		message.SetHeader("To", user.Email)
		message.SetHeader("Subject", config.Data.SmtpSubject)
		message.SetBody("text/plain", memo.Content)

		dialer := gomail.NewDialer(
			config.Data.SmtpHost,
			config.Data.SmtpPort,
			config.Data.SmtpUsername,
			config.Data.SMTPPassword,
		)
		if err = dialer.DialAndSend(message); err != nil {
			log.Println("dialer.DialAndSend :: err", err)
			continue
		}
	}

	return nil
}
