package seeds

import (
	"api-automation-backend/models"
	"api-automation-backend/pkg/helper"

	"github.com/brianvoe/gofakeit/v4"
	"xorm.io/xorm"
)

func CreateAccount(engine *xorm.Engine, account, password, name string, isDisable int) error {
	acc := models.Account{
		Account:   account,
		Password:  helper.ScryptStr(password),
		Name:      name,
		IsDisable: isDisable,
	}
	_, err := engine.Insert(&acc)
	return err
}

func AllAccount() []Seed {
	return []Seed{
		{
			Name: "Create Account 1",
			Run: func(engine *xorm.Engine) error {
				err := CreateAccount(engine, "admin", "12345678", "administrator", 0)
				return err
			},
		},
		{
			Name: "Create Account 2",
			Run: func(engine *xorm.Engine) error {
				err := CreateAccount(engine, gofakeit.Email(), "12345678", gofakeit.Name(), 0)
				return err
			},
		},
		{
			Name: "Create Account 3 with is disable = 1",
			Run: func(engine *xorm.Engine) error {
				err := CreateAccount(engine, gofakeit.Email(), "12345678", gofakeit.Name(), 1)
				return err
			},
		},
	}
}
