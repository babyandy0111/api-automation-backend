package routes

import (
	"api-automation-backend/api"
	"time"

	"api-automation-backend/pkg/reqcache"
	"github.com/gin-gonic/gin"
)

func ApiTokenV1(r *gin.Engine, store reqcache.CacheStore) {
	v1 := r.Group("/v1")

	v1.POST("/api-token", reqcache.CachePage(store, time.Second*1, func(c *gin.Context) {
		api.GetApiToken(c)
	}))
}
