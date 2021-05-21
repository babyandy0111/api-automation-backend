package service

import (
	"api-automation-backend/internal/account"
	"api-automation-backend/models"
	"api-automation-backend/models/apireq"
	"api-automation-backend/pkg/er"
	"api-automation-backend/pkg/helper"
)

type AccountService struct {
	repo account.Repository
}

func NewAccountService(r account.Repository) account.Service {
	return &AccountService{repo: r}
}

func (s *AccountService) AddAccount(req *apireq.AddAccount) (*models.Account, error) {
	var err error

	// check account exist
	accExist, err := s.repo.Exist(&models.Account{Account: req.Account})
	if err != nil {
		err = er.NewAppErr(500, er.UnknownError, "check account exist error.", err)
		return nil, err
	}

	if accExist {
		err = er.NewAppErr(400, er.ErrorParamInvalid, "account is used.", nil)
		return nil, err
	}

	acc := models.Account{
		Account:   req.Account,
		Password:  helper.ScryptStr(req.Password), // password to sha1
		Name:      req.Name,
		IsDisable: 0, // 預設啟用
	}

	res, err := s.repo.Insert(&acc)
	if err != nil {
		err = er.NewAppErr(500, er.DBInsertError, "insert account error.", err)
		return nil, err
	}

	res.Password = "********"

	return res, nil
}
