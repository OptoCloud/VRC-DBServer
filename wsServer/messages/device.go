package messages

type Device struct {
	Id           string `json:"dev_id"`
	Name         string `json:"name"`
	OwnerId      string `json:"owner_id"`
	SerialNumber string `json:"serial_num"`
}
