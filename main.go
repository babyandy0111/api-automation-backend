package main

import (
	"api-automation-backend/api"
	"api-automation-backend/config"
	"api-automation-backend/pkg/logr"
	"api-automation-backend/pkg/valider"
	"api-automation-backend/routes"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
	_ "go.uber.org/automaxprocs"
	"go.uber.org/zap"
)

var port string

// Interrupt handler.
var errChan = make(chan error, 1)

// @title Api Automation Backend
// @version 1.0
// @description Api Automation Backend
// @termsOfService https://github.com/babyandy0111/api-automation-backend
// @license.name MIT
// @license.url
func main() {
	// 從 command line 讀取 -port 的值，並指定給 port 這個 pointer。若不指定的話，預設使用 8080
	flag.StringVar(&port, "port", "8080", "Initial port number")
	flag.Parse()

	config.InitEnv()

	initialize()

	r := routes.Init()

	// Start gin server
	go func() {
		logr.L.Info("Api-Automation-Backend server start", zap.String("port", port))
		errChan <- r.Run(":" + port)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	<-errChan

	api.CloseXorm()
	api.CloseRedis()

	logr.L.Warn("Api-Automation-Backend server shutdown")
}

func initialize() {
	// init logger
	logr.InitGlobalLogger("")

	// validation init
	valider.Init()

	_ = api.InitXorm()
	_ = api.InitRedis()
}
