package messages

type FriendRequest struct {
	Id       string `json:"req_id"`
	FromName string `json:"from_name"`
}
