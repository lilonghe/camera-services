package model

type Camera struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Url         string `json:"url"`
	Company     string `json:"company"`
	TrackId     string `json:"track_id"`
	UpdateUrl   string `json:"update_url"`
}

func (Camera) TableName() string { return "cameras" }
