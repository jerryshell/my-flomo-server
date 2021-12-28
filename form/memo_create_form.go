package form

type MemoCreateForm struct {
	Content string `json:"content" required:"true"`
	Token   string `json:"token"`
}

type MemoUpdateForm struct {
	ID      string `json:"id" required:"true"`
	Content string `json:"content" required:"true"`
	Token   string `json:"token" required:"true"`
}
