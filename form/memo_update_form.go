package form

type MemoUpdateForm struct {
	ID      string `json:"id" required:"true"`
	Content string `json:"content" required:"true"`
}
