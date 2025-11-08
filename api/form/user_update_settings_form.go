package form

type UserUpdateSettingsForm struct {
	DailyReviewEnabled bool   `json:"dailyReviewEnabled"`
	TelegramChatID     string `json:"telegramChatId"`
	TelegramBotToken   string `json:"telegramBotToken"`
}
