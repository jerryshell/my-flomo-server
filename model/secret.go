package model

type Secret struct {
	BaseModel
	UserId string `json:"userId"`
	Secret string `json:"secret"`
}
