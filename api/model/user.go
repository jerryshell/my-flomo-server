package model

type User struct {
	BaseModel
	Password           string `json:"-"`
	Email              string `json:"email"`
	PluginToken        string `json:"pluginToken"`
	DailyReviewEnabled bool   `json:"dailyReviewEnabled" gorm:"default:false"`
}
