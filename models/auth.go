package models

type AccountKeyGetRequest struct {
	RegistrationKey string `json:"regkey"`
}
type RespAuthLogout struct {
	Success bool `json:"success"`
}
