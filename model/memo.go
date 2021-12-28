package model

type Memo struct {
	BaseModel
	Content string `json:"content"`
	UserId  string `json:"userId"`
}
