package model

type Mention struct {
	BaseModel
	Username string `json:"username"`
}
