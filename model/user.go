package model

type User struct {
	BaseModel
	Username string `json:"username"`
	Password string `json:"-"`
	Email    string `json:"email"`
}
