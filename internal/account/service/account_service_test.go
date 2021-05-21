package service

import (
	"api-automation-backend/config"
	"api-automation-backend/driver"
	"api-automation-backend/internal/account/repository"
	"api-automation-backend/models"
	"api-automation-backend/models/apireq"
	"api-automation-backend/pkg/er"
	"log"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestAccountServiceAddAccount(t *testing.T) {
	// Arrange
	orm, _ := driver.NewXorm()
	repo := repository.NewMysqlAccountRepo(orm)
	serv := NewAccountService(repo)

	// TEST account exist error
	req := apireq.AddAccount{
		Account:  "admin",
		Password: "12345678",
		Name:     "name",
	}
	_, err := serv.AddAccount(&req)
	existErr := err.(*er.AppError)
	assert.Equal(t, "400400", existErr.Code)
	assert.Equal(t, 400, existErr.StatusCode)

	// TEST Success
	req.Account = "admin_1"
	res, err := serv.AddAccount(&req)
	assert.Nil(t, err)

	// TearDown
	_, _ = orm.ID(res.Id).Delete(&models.Account{})
}

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}

func setUp() {
	remoteBranch := os.Getenv("REMOTE_BRANCH")
	if remoteBranch == "" {
		// load env
		err := godotenv.Load(config.GetBasePath() + "/.env")
		if err != nil {
			log.Panicln(err)
		}
	}
}
