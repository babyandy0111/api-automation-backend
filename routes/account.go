package routes

import (
	"api-automation-backend/api"

	"api-automation-backend/pkg/reqcache"

	"github.com/gin-gonic/gin"
)

func AccountV1(r *gin.Engine, store reqcache.CacheStore) {
	v1 := r.Group("/v1")

	// Add Account
	v1.POST("/accounts", api.AddAccount)
}
