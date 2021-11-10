package messages

type MessageId uint16

const (
	NoneId MessageId = iota
	Cmd_LogOut

	DeviceId
	FriendRequestId
	GroupInviteId
	SessionDirectId
	Event_AccountPasswordChangedId
	Event_AccountEmailChangedId
	Event_AccountBannedId
	Event_GroupAddedId
	Event_GroupDeletedId
	Event_GroupNameChangedId
	Event_GroupImageChangedId
	Event_GroupMemberAddedId
	Event_GroupMemberRemovedId
	Event_GroupMemberNicknameChangedId
	Event_RateLimitId
	Event_UserDeletedId
	Event_UserNameChangedId
	Event_UserProfilePictureChangedId
)

type Message struct {
	Id   MessageId   `json:"m"`
	Data interface{} `json:"d"`
}
