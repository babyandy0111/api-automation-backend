package service

import (
	"api-automation-backend/internal/account"
	"api-automation-backend/internal/token_auth"
	"api-automation-backend/pkg/er"
	"api-automation-backend/pkg/helper"
	"api-automation-backend/pkg/token_lib"
	"strconv"
	"time"
)

type TokenService struct {
	accRepo account.Repository
}

func NewTokenService(ar account.Repository) token_auth.Service {
	return &TokenService{accRepo: ar}
}

func (s *TokenService) GenJwtToken(accId int64, passwrod string) (string, time.Time, error) {
	acc, err := s.accRepo.FindOne(accId)
	if err != nil || acc == nil {
		notFoundErr := er.NewAppErr(401, er.UnauthorizedError, "", err)
		return "", time.Time{}, notFoundErr
	}

	if acc.IsDisable != 0 {
		authErr := er.NewAppErr(401, er.UnauthorizedError, "", nil)
		return "", time.Time{}, authErr
	}

	passwrod = helper.ScryptStr(passwrod)
	if acc.Password != passwrod {
		authErr := er.NewAppErr(401, er.UnauthorizedError, "", nil)
		return "", time.Time{}, authErr
	}

	// start generate jwt token
	accIdStr := strconv.FormatInt(acc.Id, 10)

	token, expireAt, err := token_lib.GenToken(accIdStr, acc.Account)
	if err != nil {
		unknownErr := er.NewAppErr(500, er.UnknownError, "", err)
		return "", time.Time{}, unknownErr
	}

	return token, expireAt, nil
}
