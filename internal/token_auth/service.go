package token_auth

import (
	"time"
)

type Service interface {
	GenJwtToken(accId int64, password string) (string, time.Time, error)
}
