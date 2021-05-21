package logr

import (
	"go.uber.org/zap"
	"log"
	"os"
)

var L *zap.Logger

const (
	EnvProduction  = "production"
	EnvDevelopment = "development"
	EnvLocalhost   = "localhost"
)

// InitGlobalLogger initialize a global logger (logr.L)
func InitGlobalLogger(environment string) {
	var err error
	env := os.Getenv("ENVIRONMENT")

	if environment != "" {
		env = environment
	}

	if env == EnvProduction {
		// 設定為線上 log level
		L, err = zap.NewProduction()
	} else if env == EnvLocalhost {
		L, err = zap.NewDevelopment()
	} else {
		config := zap.NewProductionConfig()
		config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
		L, err = config.Build()
	}

	defer func() {
		_ = L.Sync()
	}()

	if err != nil {
		log.Panicln("Init zap log failed...")
	}
}

func NewLogger(environment string) *zap.Logger {
	var err error
	var logger *zap.Logger
	env := os.Getenv("ENVIRONMENT")

	if environment != "" {
		env = environment
	}

	if env != EnvLocalhost {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}

	defer func() {
		_ = logger.Sync()
	}()

	if err != nil {
		log.Println("Init zap log failed...")
	}

	return logger
}
