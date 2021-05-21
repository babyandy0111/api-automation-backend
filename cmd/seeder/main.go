package main

import (
	"api-automation-backend/config"
	"api-automation-backend/pkg/seeds"
	"fmt"
	"os"
	"time"

	"api-automation-backend/pkg/logr"
	"github.com/brianvoe/gofakeit/v4"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"xorm.io/xorm"
)

func main() {
	var err error
	remoteBranch := os.Getenv("REMOTE_BRANCH")

	logger := logr.NewLogger("")
	if remoteBranch == "" {
		// load env
		err = godotenv.Load()

		if err != nil {
			logger.Debug(err.Error())
		}
	}

	dsn := "%s:%s@(%s:%s)/%s?parseTime=true"

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := config.GetDBName()
	engine, err := xorm.NewEngine("mysql", fmt.Sprintf(dsn, dbUser, dbPassword, dbHost, dbPort, dbName))

	engine.TZLocation, _ = time.LoadLocation("UTC")
	engine.DatabaseTZ, _ = time.LoadLocation("UTC")

	if err != nil {
		logger.Error(err.Error())
	}

	gofakeit.Seed(time.Now().Unix())

	// Create Accounts	----------------------------------------------------------------------------
	accSeeds := seeds.AllAccount()
	run(engine, accSeeds)
}

func run(engine *xorm.Engine, channelSeeds []seeds.Seed) {
	logger := logr.NewLogger("")
	for _, seed := range channelSeeds {
		logger.Info(seed.Name)
		err := seed.Run(engine)
		if err != nil {
			logger.Error(seed.Name + " Failed")
			logger.Error(err.Error())
		}
	}
}
