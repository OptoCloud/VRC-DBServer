package messages

type Invite struct {
	Id           string        `json:"invite_id"`
	GroupName    string        `json:"group_name"`
	GroupImageId string        `json:"group_image_id"`
	Members      []GroupMember `json:"members"`
}
