package model

type PluginToken struct {
	BaseModel
	UserID string `json:"userId"`
	Token  string `json:"token"`
}
