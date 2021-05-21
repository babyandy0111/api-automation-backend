package repository

import (
	"api-automation-backend/internal/account"
	"api-automation-backend/models"

	"xorm.io/xorm"
)

type AccountRepo struct {
	orm *xorm.EngineGroup
}

func NewMysqlAccountRepo(orm *xorm.EngineGroup) account.Repository {
	return &AccountRepo{orm: orm}
}

func (r *AccountRepo) Exist(acc *models.Account) (bool, error) {
	has, err := r.orm.Exist(acc)
	return has, err
}

func (r *AccountRepo) Insert(acc *models.Account) (*models.Account, error) {
	_, err := r.orm.Insert(acc)
	if err != nil {
		return nil, err
	}

	return acc, nil
}

func (r *AccountRepo) FindOne(id int64) (*models.Account, error) {
	acc := models.Account{}

	has, err := r.orm.ID(id).Get(&acc)
	if err != nil {
		return nil, err
	}

	if !has {
		return nil, nil
	}

	return &acc, err
}
