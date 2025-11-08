package service

import (
	"errors"
	"math/rand"
	"strconv"
	"time"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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

// sendTelegramMessage 发送Telegram消息
func sendTelegramMessage(botToken string, chatID string, message string) error {
	logger := util.NewLogger("telegram_service")

	bot, err := telegram.NewBotAPI(botToken)
	if err != nil {
		logger.Error("failed to create telegram bot", util.ErrorField(err))
		return err
	}

	chatIDInt, err := strconv.ParseInt(chatID, 10, 64)
	if err != nil {
		logger.Error("failed to parse chat id", util.ErrorField(err), util.StringField("chat_id", chatID))
		return err
	}

	msg := telegram.NewMessage(chatIDInt, message)
	_, err = bot.Send(msg)
	if err != nil {
		logger.Error("failed to send telegram message", util.ErrorField(err), util.StringField("chat_id", chatID))
		return err
	}

	logger.Info("telegram message sent successfully", util.StringField("chat_id", chatID))
	return nil
}

func MemoDailyReview() error {
	logger := util.NewLogger("memo_service")

	// 处理邮件每日回顾
	emailUserList, err := UserListByEmailIsNotNull()
	if err != nil {
		logger.Error("failed to get user list with email", util.ErrorField(err))
		return err
	}

	// 处理Telegram每日回顾
	telegramUserList, err := UserListWithTelegramConfig()
	if err != nil {
		logger.Error("failed to get user list with telegram config", util.ErrorField(err))
		return err
	}

	if len(emailUserList) == 0 && len(telegramUserList) == 0 {
		logger.Warn("both email and telegram user lists are empty")
		return errors.New("用户数据为空")
	}

	logger.Info("starting daily review for users",
		util.IntField("email_user_count", len(emailUserList)),
		util.IntField("telegram_user_count", len(telegramUserList)))

	emailSuccessCount := 0
	emailFailCount := 0
	telegramSuccessCount := 0
	telegramFailCount := 0

	// 检查SMTP配置是否完整
	smtpConfigured := config.Data.SmtpUsername != "" && config.Data.SmtpPassword != ""
	if !smtpConfigured {
		logger.Warn("SMTP credentials not configured, skipping email sending")
	}

	// 处理邮件用户
	if smtpConfigured {
		// 邮件发送配置
		dialer := gomail.NewDialer(
			config.Data.SmtpHost,
			config.Data.SmtpPort,
			config.Data.SmtpUsername,
			config.Data.SmtpPassword,
		)

		for _, user := range emailUserList {
			logger.Debug("processing user for email daily review", util.StringField("user_id", user.ID), util.StringField("user_email", user.Email))

			// 检查用户是否开启了每日回顾功能
			if !user.DailyReviewEnabled {
				logger.Debug("daily review disabled for user, skipping", util.StringField("user_id", user.ID), util.StringField("user_email", user.Email))
				continue
			}

			memo, err := MemoGetRandomByUserID(user.ID)
			if err != nil {
				logger.Error("failed to get random memo for user", util.ErrorField(err), util.StringField("user_id", user.ID))
				emailFailCount++
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
				emailFailCount++
				continue
			}

			logger.Info("daily review email sent successfully", util.StringField("user_id", user.ID), util.StringField("user_email", user.Email))
			emailSuccessCount++
		}
	} else {
		// SMTP未配置，记录日志
		for _, user := range emailUserList {
			if user.DailyReviewEnabled {
				logger.Info("email daily review skipped due to missing SMTP configuration", util.StringField("user_id", user.ID), util.StringField("user_email", user.Email))
			}
		}
	}

	// 处理Telegram用户
	for _, user := range telegramUserList {
		logger.Debug("processing user for telegram daily review", util.StringField("user_id", user.ID), util.StringField("telegram_chat_id", user.TelegramChatID))

		// 检查用户是否开启了每日回顾功能
		if !user.DailyReviewEnabled {
			logger.Debug("daily review disabled for user, skipping", util.StringField("user_id", user.ID), util.StringField("telegram_chat_id", user.TelegramChatID))
			continue
		}

		memo, err := MemoGetRandomByUserID(user.ID)
		if err != nil {
			logger.Error("failed to get random memo for user", util.ErrorField(err), util.StringField("user_id", user.ID))
			telegramFailCount++
			continue
		}

		logger.Debug("selected memo for user", util.StringField("user_id", user.ID), util.StringField("memo_id", memo.ID))

		// 构建Telegram消息
		telegramMessage := "📝 My Flomo 每日回顾\n\n" + memo.Content

		if err = sendTelegramMessage(user.TelegramBotToken, user.TelegramChatID, telegramMessage); err != nil {
			logger.Error("failed to send telegram message to user", util.ErrorField(err), util.StringField("user_id", user.ID), util.StringField("telegram_chat_id", user.TelegramChatID))
			telegramFailCount++
			continue
		}

		logger.Info("daily review telegram message sent successfully", util.StringField("user_id", user.ID), util.StringField("telegram_chat_id", user.TelegramChatID))
		telegramSuccessCount++
	}

	logger.Info("daily review completed",
		util.IntField("email_success_count", emailSuccessCount),
		util.IntField("email_fail_count", emailFailCount),
		util.IntField("telegram_success_count", telegramSuccessCount),
		util.IntField("telegram_fail_count", telegramFailCount))

	return nil
}
