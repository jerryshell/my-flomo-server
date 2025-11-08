package form

type UserUpdatePasswordForm struct {
	Password string `json:"password" required:"true"`
}
