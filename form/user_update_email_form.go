package form

type UserUpdateEmailForm struct {
	Email string `json:"email" required:"true"`
}
