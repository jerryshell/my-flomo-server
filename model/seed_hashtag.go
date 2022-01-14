package model

type SeedHashtag struct {
	BaseModel
	SeedID string `json:"seedId"`
	Name   string `json:"name"`
}
