package form

type UserLoginOrRegisterForm struct {
	Email    string `json:"email" required:"true"`
	Password string `json:"password" required:"true"`
}
