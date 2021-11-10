package messages

type GroupMember struct {
	UserId   string `json:"usr_id"`
	NickName string `json:"nickname"`
}
type GroupAdded struct {
	GroupId string        `json:"grp_id"`
	Name    string        `json:"name"`
	ImageId string        `json:"img_id"`
	Members []GroupMember `json:"members"`
}
type GroupDeleted struct {
	GroupId string `json:"grp_id"`
}
type GroupNameChanged struct {
	GroupId string `json:"grp_id"`
	Name    string `json:"name"`
}
type GroupImageChanged struct {
	GroupId string `json:"grp_id"`
	ImageId string `json:"img_id"`
}
type GroupMemberAdded struct {
	UserId string `json:"usr_id"`
}
type GroupMemberRemoved struct {
	UserId string `json:"usr_id"`
}
type GroupMemberNicknameChanged struct {
	UserId   string `json:"usr_id"`
	NickName string `json:"nickname"`
}
