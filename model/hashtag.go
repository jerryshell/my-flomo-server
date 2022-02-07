package model

type Hashtag struct {
	BaseModel
	HashtagGroupID string `json:"hashtagGroupId"`
	Name           string `json:"name"`
}
