package model

type User struct {
	BaseModel
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
