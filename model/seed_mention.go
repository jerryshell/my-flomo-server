package model

type SeedMention struct {
	BaseModel
	SeedID   string `json:"seedId"`
	Username string `json:"username"`
}
