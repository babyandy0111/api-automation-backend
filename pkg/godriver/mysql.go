package godriver

import (
	"fmt"
	"log"
	"sync"
	"time"
	"xorm.io/xorm"
	xlog "xorm.io/xorm/log"
)

var (
	orm     *xorm.EngineGroup
	ormOnce sync.Once
)

// NewXorm return singleton xorm instance
func NewXorm(cfg *XormConfig) (*xorm.EngineGroup, error) {
	var err error
	ormOnce.Do(func() {
		err = newXorm(cfg)
	})
	return orm, err
}

type XormConfig struct {
	DSN      string
	SlaveDSN string
	Debug    string
}

func BuildDsn(user, password, host, port, dbname string) string {
	dsn := "%s:%s@(%s:%s)/%s?parseTime=true&charset=utf8mb4&collation=utf8mb4_unicode_ci"
	return fmt.Sprintf(dsn, user, password, host, port, dbname)
}

func newXorm(cfg *XormConfig) error {
	var err error

	// Connect to master
	master, err := xorm.NewEngine("mysql", cfg.DSN)

	if err != nil {
		log.Println("[MySQL] Connect to MySQL Master error", err)
		return err
	}

	log.Println("[MySQL] Connected to MySQL Master")

	// Connect to slave
	slave1, err := xorm.NewEngine("mysql", cfg.SlaveDSN)

	if err != nil {
		log.Println("[MySQL] Connect to MySQL Slave1 error", err)
		return err
	} else {
		log.Println("[MySQL] Connected to MySQL Slave1")
	}

	master.TZLocation, _ = time.LoadLocation("UTC")
	master.DatabaseTZ, _ = time.LoadLocation("UTC")
	slave1.TZLocation, _ = time.LoadLocation("UTC")
	slave1.DatabaseTZ, _ = time.LoadLocation("UTC")

	slaves := []*xorm.Engine{slave1}
	orm, err = xorm.NewEngineGroup(master, slaves, xorm.LeastConnPolicy())

	if err != nil {
		log.Println(err)
		return err
	}

	if cfg.Debug == "debug" {
		log.Println("[MySQL] XORM debug mode enabled")
		orm.ShowSQL(true)
		orm.Logger().SetLevel(xlog.LOG_DEBUG)
		pingErr := orm.Ping()
		if pingErr != nil {
			log.Println("[MySQL] " + pingErr.Error())
		}
	}

	return err
}
