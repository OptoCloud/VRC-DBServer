package messages

import (
	"time"
)

type AccountPasswordChanged struct {
	Salt      string `json:"salt"`
	CpuLimit  uint64 `json:"cpu_limit"`
	RamLimit  uint64 `json:"ram_limit"`
	Algorithm uint32 `json:"algorithm"`
}
type AccountEmailChanged struct {
	Username string `json:"username"`
}
type AccountBanned struct {
	BanExpiresAt time.Time `json:"ban_expires_at"`
}
