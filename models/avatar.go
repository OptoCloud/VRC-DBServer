package models

import "time"

type AvatarDbRecord struct {
	Id            string    `json:"id" bson:"_id"`
	Name          string    `json:"name" bson:"name"`
	AssetURL      string    `json:"asset_url" bson:"asset_url"`
	ImageURL      string    `json:"image_url" bson:"image_url"`
	ThumbnailURL  string    `json:"thumb_url" bson:"thumb_url"`
	AuthorId      string    `json:"author_id" bson:"author_id"`
	ReleaseStatus string    `json:"release" bson:"release"`
	UpdatedAt     time.Time `json:"updated_at" bson:"updated_at"`
	DetectedAt    time.Time `json:"detected_at" bson:"detected_at"`
}
type AvatarPostRequest struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	AssetURL      string `json:"asset_url"`
	ImageURL      string `json:"image_url"`
	ThumbnailURL  string `json:"thumb_url"`
	AuthorId      string `json:"author_id"`
	ReleaseStatus string `json:"release"`
}
