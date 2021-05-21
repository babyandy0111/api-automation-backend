package account

import (
	"api-automation-backend/models"
	"api-automation-backend/models/apireq"
)

type Service interface {
	AddAccount(req *apireq.AddAccount) (*models.Account, error)
}
