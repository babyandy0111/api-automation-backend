package config

import (
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"time"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

const (
	EnvProduction          = "production"
	EnvDevelopment         = "development"
	EnvLocalhost           = "localhost"
	AdminUserId            = int64(-1)
	RedisDefaultExpireTime = time.Second * 60 * 60 * 24 * 30 // 預設一個月
)

var EnvShortName = map[string]string{
	"production":  "prod",
	"development": "dev",
	"localhost":   "local",
}

func GetEnvironment() string {
	return os.Getenv("ENVIRONMENT")
}

var (
	_, b, _, _ = runtime.Caller(0)
	basePath   = filepath.Dir(b)
)

func GetBasePath() string {
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}

func GetRedisDefaultExpireTime() time.Duration {
	max := 12
	min := 1
	rand.Seed(time.Now().UnixNano())
	// n will be between 1 and 12
	n := rand.Intn(max-min) + 1
	expireTime := RedisDefaultExpireTime + (time.Hour * time.Duration(n))

	return expireTime
}

func GetShortEnv() string {
	switch GetEnvironment() {
	case EnvLocalhost:
		return "local"
	case EnvDevelopment:
		return "dev"
	case EnvProduction:
		return "prod"
	default:
		return ""
	}
}

func InitEnv() {
	remoteBranch := os.Getenv("REMOTE_BRANCH")

	if remoteBranch == "" {
		// load env from .env file
		path := GetBasePath() + "/.env"
		err := godotenv.Load(path)

		if err != nil {
			log.Panicln(err)
		}
	}
}

func GetDBName() string {
	switch GetEnvironment() {
	case EnvLocalhost:
		return os.Getenv("DB_NAME")
	case EnvDevelopment:
		return "dev_indochat"
	case EnvProduction:
		return "prod_indochat"
	default:
		return ""
	}
}
