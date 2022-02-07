package model

type BlockSubscriber struct {
	BaseModel
	BlockID string `json:"blockId"`
	UserID  string `json:"userId"`
}
