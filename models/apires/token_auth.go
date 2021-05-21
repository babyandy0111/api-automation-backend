package apires

import "time"

type TokenSuccess struct {
	Token    string    `json:"token"`
	ExpireAt time.Time `json:"expire_at"`
}
