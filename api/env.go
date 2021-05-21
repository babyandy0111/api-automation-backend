package api

import (
	"api-automation-backend/driver"
	"log"

	"github.com/go-redis/redis/v7"
	"xorm.io/xorm"
)

type Env struct {
	Orm    *xorm.EngineGroup
	Redis  *redis.ClusterClient
	Engine *xorm.Engine
}

var env = &Env{}

func GetEnv() *Env {
	return env
}

func InitXorm() *xorm.EngineGroup {
	var err error
	env.Orm, err = driver.NewXorm()
	if err != nil {
		log.Println(err)
	}

	return env.Orm
}

func InitRedis() *redis.ClusterClient {
	var err error
	env.Redis, err = driver.NewRedis()
	if err != nil {
		log.Println(err)
	}

	return env.Redis
}

func InitEngine() *xorm.Engine {
	var err error
	env.Engine, err = driver.NewEngine()
	if err != nil {
		log.Println(err)
	}

	return env.Engine
}

func CloseXorm() error {
	if env.Orm != nil {
		return env.Orm.Close()
	}
	return nil
}

func CloseRedis() error {
	if env.Redis != nil {
		return env.Redis.Close()
	}
	return nil
}
