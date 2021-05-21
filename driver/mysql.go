package driver

import (
	"api-automation-backend/config"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/indochat/godriver"
	log "github.com/sirupsen/logrus"
	"xorm.io/xorm"
	xlog "xorm.io/xorm/log"
)

var (
	engine     *xorm.Engine
	engineOnce sync.Once
)

// NewXorm return singleton xorm instance
func NewXorm() (*xorm.EngineGroup, error) {
	cfg := godriver.XormConfig{}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := config.GetDBName()

	dbSlave1User := os.Getenv("DB_SLAVE1_USER")
	dbSlave1Password := os.Getenv("DB_SLAVE1_PASSWORD")
	dbSlave1Host := os.Getenv("DB_SLAVE1_HOST")
	dbSlave1Port := os.Getenv("DB_SLAVE1_PORT")
	dbSlave1Name := config.GetDBName()

	cfg.DSN = godriver.BuildDsn(dbUser, dbPassword, dbHost, dbPort, dbName)
	cfg.SlaveDSN = godriver.BuildDsn(dbSlave1User, dbSlave1Password, dbSlave1Host, dbSlave1Port, dbSlave1Name)
	cfg.Debug = os.Getenv("XORM_MODE")

	return godriver.NewXorm(&cfg)
}

// NewEngine return singleton xorm engine instance
// Note: 這個單一連線是 for oauth library 用的，目前尚未改寫成支援 xorm engine group
func NewEngine() (*xorm.Engine, error) {
	var err error
	engineOnce.Do(func() {
		err = newEngine()
	})
	return engine, err
}

func newEngine() error {
	var err error
	dsn := "%s:%s@(%s:%s)/%s?parseTime=true&charset=utf8mb4&collation=utf8mb4_unicode_ci"

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := config.GetDBName()
	engine, err = xorm.NewEngine("mysql", fmt.Sprintf(dsn, dbUser, dbPassword, dbHost, dbPort, dbName))

	if err != nil {
		log.Println("[MySQL Engine] Connect to MySQL Master error...", err)
		return err
	} else {
		log.Println("[MySQL Engine] [MySQL] Connected to MySQL Master...", err)
	}

	engine.TZLocation, _ = time.LoadLocation("UTC")
	engine.DatabaseTZ, _ = time.LoadLocation("UTC")

	if err != nil {
		log.Println(err)
		return err
	}

	if os.Getenv("XORM_MODE") == "debug" {
		log.WithFields(log.Fields{}).Info("[MySQL Engine] XORM debug mode enabled")
		engine.ShowSQL(true)
		engine.Logger().SetLevel(xlog.LOG_DEBUG)
		pingErr := engine.Ping()
		if pingErr != nil {
			log.WithFields(log.Fields{}).Error(pingErr.Error())
		}
	}

	return err
}
