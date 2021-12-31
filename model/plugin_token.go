package model

type PluginToken struct {
	BaseModel
	UserId string `json:"userId"`
	Token  string `json:"token"`
}
