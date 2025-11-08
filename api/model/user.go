package model

type User struct {
	BaseModel
	Password           string `json:"-"`
	Email              string `json:"email"`
	PluginToken        string `json:"pluginToken"`
	DailyReviewEnabled bool   `json:"dailyReviewEnabled" gorm:"default:false"`
	TelegramChatID     string `json:"telegramChatId" gorm:"default:''"`
	TelegramBotToken   string `json:"telegramBotToken" gorm:"default:''"`
}
