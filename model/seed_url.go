package model

type SeedUrl struct {
	BaseModel
	SeedID string `json:"seedId"`
	Url    string `json:"url"`
}
