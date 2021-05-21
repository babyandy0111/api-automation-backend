package routes

import (
	"api-automation-backend/config"
	_ "api-automation-backend/docs"
	"api-automation-backend/middleware"
	"api-automation-backend/pkg/packs"
	_ "embed"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	golibMiddleware "github.com/indochat/golib/middleware"
	"github.com/indochat/golib/reqcache"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Init() *gin.Engine {
	r := gin.New()

	// gin 檔案上傳 body 限制
	r.MaxMultipartMemory = 64 << 20 // 8 MiB

	// request cache
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	addr := fmt.Sprintf("%s:%s", redisHost, redisPort)
	store := reqcache.NewRedisCache(addr, "", time.Second)

	environment := os.Getenv("ENVIRONMENT")

	r.Use(golibMiddleware.LogRequest())
	r.Use(middleware.ErrorResponse())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if environment == config.EnvDevelopment || environment == config.EnvLocalhost {
		r.GET("/test/coverage", func(c *gin.Context) {
			file, _ := packs.PackTestCovertFile()
			c.Data(http.StatusOK, "text/html; charset=utf-8", file)
		})
	}

	corsConf := cors.DefaultConfig()
	corsConf.AllowAllOrigins = true
	corsConf.AllowCredentials = true
	corsConf.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"}
	corsConf.AllowHeaders = []string{"Origin", "X-Requested-With", "Content-Type", "Accept", "Authorization", "Bearer", "Accept-Language"}
	r.Use(cors.New(corsConf))

	AccountV1(r, store)
	ApiTokenV1(r, store)

	return r
}
