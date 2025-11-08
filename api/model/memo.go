package model

type Memo struct {
	BaseModel
	UserID  string `json:"userId"`
	Content string `json:"content"`
}
