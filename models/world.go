package models

import "time"

type WorldUpdate struct {
	AssetURL   string    `json:"asset_url" bson:"asset_url"`
	ImageURL   string    `json:"image_url" bson:"image_url"`
	UploadedAt time.Time `json:"uploaded_at" bson:"uploaded_at"`
}

type WorldDbRecord struct {
	Id         string      `json:"id" bson:"_id"`
	Name       string      `json:"name" bson:"name"`
	Capacity   int         `json:"capacity" bson:"capacity"`
	AuthorId   string      `json:"author_id" bson:"author_id"`
	Updates    WorldUpdate `json:"updates" bson:"updates"`
	CreatedAt  time.Time   `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at" bson:"updated_at"`
	DetectedAt time.Time   `json:"detected_at" bson:"detected_at"`
}
type WorldUpdateCommand struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	AssetURL     string `json:"asset_url"`
	ImageURL     string `json:"image_url"`
	ThumbnailURL string `json:"thumb_url"`
	AuthorId     string `json:"author_id"`
}
