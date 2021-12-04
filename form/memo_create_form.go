package form

type MemoCreateForm struct {
	Content string `json:"content" required:"true"`
}
