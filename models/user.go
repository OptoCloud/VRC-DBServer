package models

import "time"

type UserDbRecord struct {
	Id            string    `json:"id" bson:"_id"`
	Name          string    `json:"name" bson:"name"`
	PreviousNames []string  `json:"prev_names" bson:"prev_names"`
	DeveloperType string    `json:"dev_type" bson:"dev_type"`
	TrustLevel    string    `json:"trust_lvl" bson:"trust_lvl"`
	SystemTags    []string  `json:"system_tags" bson:"system_tags"`
	AdminTags     []string  `json:"admin_tags" bson:"admin_tags"`
	MiscTags      []string  `json:"misc_tags" bson:"misc_tags"`
	HasVrcPlus    bool      `json:"has_vrcp" bson:"has_vrcp"`
	WornAvatarIds []string  `json:"avtr_ids" bson:"avtr_ids"`
	UpdatedAt     time.Time `json:"updated_at" bson:"updated_at"`
	DetectedAt    time.Time `json:"detected_at" bson:"detected_at"`
}
type UserPostRequest struct {
	Id             string   `json:"id"`
	DisplayName    string   `json:"displayname"`
	DeveloperType  string   `json:"devtype"`
	IconURL        string   `json:"icon_url"`
	AvatarImageURL string   `json:"avtr_img_url"`
	WorldId        string   `json:"world_id"`
	Tags           []string `json:"tags"`
}
