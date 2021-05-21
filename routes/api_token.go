package routes

import (
	"api-automation-backend/api"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/indochat/golib/reqcache"
)

func ApiTokenV1(r *gin.Engine, store reqcache.CacheStore) {
	v1 := r.Group("/v1")

	v1.POST("/api-token", reqcache.CachePage(store, time.Second*1, func(c *gin.Context) {
		api.GetApiToken(c)
	}))
}
