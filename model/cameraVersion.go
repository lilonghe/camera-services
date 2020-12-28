package model

import "time"

type CameraVersion struct {
	Id          int       `json:"id"`
	Version     string    `json:"version"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CameraId    int       `json:"camera_id"`
	PublishDate time.Time `json:"publish_date"`
	TrackId     string    `json:"track_id"`
}

func (CameraVersion) TableName() string { return "camera_versions" }
