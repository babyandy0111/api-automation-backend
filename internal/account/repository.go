package account

import (
	"api-automation-backend/models"
)

type Repository interface {
	Exist(acc *models.Account) (bool, error)
	Insert(acc *models.Account) (*models.Account, error)
	FindOne(id int64) (*models.Account, error)
}
