package service

import (
	"api-automation-backend/config"
	"api-automation-backend/driver"
	accRepo "api-automation-backend/internal/account/repository"
	"api-automation-backend/pkg/er"
	"log"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestTokenServiceGenJwtToken(t *testing.T) {
	orm, _ := driver.NewXorm()

	ar := accRepo.NewMysqlAccountRepo(orm)
	serv := NewTokenService(ar)

	// TEST account not found
	_, _, err := serv.GenJwtToken(999, "12345678")
	notFoundErr := err.(*er.AppError)
	assert.Equal(t, "400401", notFoundErr.Code)
	assert.Equal(t, 401, notFoundErr.StatusCode)

	// TEST account is disable error
	_, _, err = serv.GenJwtToken(3, "12345678")
	disableErr := err.(*er.AppError)
	assert.Equal(t, "400401", disableErr.Code)
	assert.Equal(t, 401, disableErr.StatusCode)

	// TEST password error
	_, _, err = serv.GenJwtToken(1, "11223344")
	pwdErr := err.(*er.AppError)
	assert.Equal(t, "400401", pwdErr.Code)
	assert.Equal(t, 401, pwdErr.StatusCode)

	// TEST get jwt token success
	token, expiredTime, err := serv.GenJwtToken(1, "12345678")
	assert.Nil(t, err)
	assert.NotNil(t, token)
	assert.NotNil(t, expiredTime)
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
