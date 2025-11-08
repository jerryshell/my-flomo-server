package model

type User struct {
	BaseModel
	Password string `json:"-"`
	Email    string `json:"email"`
}
