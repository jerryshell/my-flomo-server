package model

type UserBlock struct {
	BaseModel
	UserID  string `json:"userId"`
	BlockID string `json:"blockId"`
}
