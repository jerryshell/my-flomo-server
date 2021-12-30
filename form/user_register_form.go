package form

type UserRegisterForm struct {
	Username string `json:"username" required:"true"`
	Password string `json:"password" required:"true"`
}
