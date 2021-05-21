package routes

import (
	"api-automation-backend/api"

	"github.com/gin-gonic/gin"
	"github.com/indochat/golib/reqcache"
)

func AccountV1(r *gin.Engine, store reqcache.CacheStore) {
	v1 := r.Group("/v1")

	// Add Account
	v1.POST("/accounts", api.AddAccount)
}
