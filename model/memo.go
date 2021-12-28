package model

type Memo struct {
	BaseModel
	Content string `json:"content"`
	UserID  string `json:"userId"`
}
