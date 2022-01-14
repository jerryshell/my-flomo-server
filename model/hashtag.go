package model

type Hashtag struct {
	BaseModel
	HashtagGroupID string `json:"hashtag_group_id"`
	Name           string `json:"name"`
}
