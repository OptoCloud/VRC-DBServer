package messages

type UserDeleted struct {
	UserId string `json:"usr_id"`
}
type UserNameChanged struct {
	UserId string `json:"usr_id"`
	Name   string `json:"username"`
}
type UserProfilePictureChanged struct {
	UserId  string `json:"usr_id"`
	ImageId string `json:"img_id"`
}
