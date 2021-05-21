package repository

import (
	"api-automation-backend/config"
	"api-automation-backend/driver"
	"api-automation-backend/models"
	"log"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestAccountRepoExist(t *testing.T) {
	orm, _ := driver.NewXorm()
	repo := NewMysqlAccountRepo(orm)

	acc := models.Account{Account: "admin"}
	exist, err := repo.Exist(&acc)

	assert.Nil(t, err)
	assert.True(t, exist)
}

func TestAccountRepoInsert(t *testing.T) {
	orm, _ := driver.NewXorm()
	repo := NewMysqlAccountRepo(orm)

	acc := models.Account{
		Account:   "testmysql",
		Password:  "12345678",
		IsDisable: 0,
	}
	res, err := repo.Insert(&acc)
	assert.Nil(t, err)

	// TearDown
	_, _ = orm.ID(res.Id).Delete(&models.Account{})
}

func TestAccountRepoFindOne(t *testing.T) {
	orm, _ := driver.NewXorm()
	repo := NewMysqlAccountRepo(orm)

	_, err := repo.FindOne(1)

	assert.Nil(t, err)
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
