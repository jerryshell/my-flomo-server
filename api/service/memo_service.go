package service

import (
	"errors"
	"math/rand"
	"time"

	"github.com/jerryshell/my-flomo/api/config"
	"github.com/jerryshell/my-flomo/api/model"
	"github.com/jerryshell/my-flomo/api/store"
	"github.com/jerryshell/my-flomo/api/util"
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
	logger := util.NewLogger("memo_service")
	
	memo, err := store.MemoGetByID(id)
	if err != nil {
		logger.Error("failed to get memo by id", util.ErrorField(err), util.StringField("memo_id", id))
		return nil, err
	}
	if memo.ID == "" {
		logger.Warn("memo not found", util.StringField("memo_id", id))
		return nil, errors.New("找不到 memo，id: " + id)
	}

	memo.Content = content
	err = store.MemoSave(&memo)

	if err != nil {
		logger.Error("failed to save memo", util.ErrorField(err), util.StringField("memo_id", id))
	} else {
		logger.Info("memo updated successfully", util.StringField("memo_id", id))
	}

	return &memo, err
}

func MemoDeleteByID(id string) error {
	return store.MemoDeleteByID(id)
}

func MemoGetRandomByUserID(userID string) (*model.Memo, error) {
	logger := util.NewLogger("memo_service")
	
	memoList, err := MemoListByUserID(userID)
	if err != nil {
		logger.Error("failed to get memo list by user id", util.ErrorField(err), util.StringField("user_id", userID))
		return nil, err
	}
	if len(memoList) == 0 {
		logger.Warn("memo list is empty", util.StringField("user_id", userID))
		return nil, errors.New("memo 数据为空")
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	index := r.Intn(len(memoList))
	
	logger.Debug("random memo selected", util.StringField("user_id", userID), util.IntField("memo_count", len(memoList)), util.IntField("selected_index", index))

	return &memoList[index], nil
}

func MemoDailyReview() error {
	logger := util.NewLogger("memo_service")
	
	userList, err := UserListByEmailIsNotNull()
	if err != nil {
		logger.Error("failed to get user list with email", util.ErrorField(err))
		return err
	}
	if len(userList) == 0 {
		logger.Warn("user list with email is empty")
		return errors.New("用户数据为空")
	}

	logger.Info("starting daily review for users", util.IntField("user_count", len(userList)))

	dialer := gomail.NewDialer(
		config.Data.SmtpHost,
		config.Data.SmtpPort,
		config.Data.SmtpUsername,
		config.Data.SMTPPassword,
	)

	successCount := 0
	failCount := 0

	for _, user := range userList {
		logger.Debug("processing user for daily review", util.StringField("user_id", user.ID), util.StringField("user_email", user.Email))

		// 检查用户是否开启了每日回顾功能
		if !user.DailyReviewEnabled {
			logger.Debug("daily review disabled for user, skipping", util.StringField("user_id", user.ID), util.StringField("user_email", user.Email))
			continue
		}

		memo, err := MemoGetRandomByUserID(user.ID)
		if err != nil {
			logger.Error("failed to get random memo for user", util.ErrorField(err), util.StringField("user_id", user.ID))
			failCount++
			continue
		}

		logger.Debug("selected memo for user", util.StringField("user_id", user.ID), util.StringField("memo_id", memo.ID))

		message := gomail.NewMessage()
		message.SetHeader("From", config.Data.SmtpUsername)
		message.SetHeader("To", user.Email)
		message.SetHeader("Subject", config.Data.SmtpSubject)
		message.SetBody("text/plain", memo.Content)

		if err = dialer.DialAndSend(message); err != nil {
			logger.Error("failed to send email to user", util.ErrorField(err), util.StringField("user_id", user.ID), util.StringField("user_email", user.Email))
			failCount++
			continue
		}

		logger.Info("daily review email sent successfully", util.StringField("user_id", user.ID), util.StringField("user_email", user.Email))
		successCount++
	}

	logger.Info("daily review completed", util.IntField("success_count", successCount), util.IntField("fail_count", failCount))

	return nil
}
