package middleware

import (
	"api-automation-backend/api"
	accRepo "api-automation-backend/internal/account/repository"
	"api-automation-backend/pkg/er"
	"api-automation-backend/pkg/token_lib"
	"strconv"

	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Bearer")

		if token == "" {
			authErr := er.NewAppErr(401, er.UnauthorizedError, "Token is required", nil)
			c.AbortWithStatusJSON(authErr.GetStatus(), authErr.GetMsg())
			return
		}

		claims, err := token_lib.ParseToken(token)

		if err != nil {
			authErr := er.NewAppErr(401, er.UnauthorizedError, "Token is not valid", nil)
			c.AbortWithStatusJSON(authErr.GetStatus(), authErr.GetMsg())
			return
		}

		// compare jwt's device uid <=> user's device uid
		env := api.GetEnv()
		ar := accRepo.NewMysqlAccountRepo(env.Orm)

		var jwtAccId, parseAccIdOk = claims["account_id"].(string)

		if !parseAccIdOk {
			parseJwtInfoErr := er.NewAppErr(401, er.UnauthorizedError, "Token is not valid", nil)
			c.AbortWithStatusJSON(parseJwtInfoErr.GetStatus(), parseJwtInfoErr.GetMsg())
			return
		}

		accId, _ := strconv.ParseInt(jwtAccId, 10, 64)
		acc, err := ar.FindOne(accId)
		if err != nil || acc == nil {
			parseJwtInfoErr := er.NewAppErr(401, er.UnauthorizedError, "Account is not found", nil)
			c.AbortWithStatusJSON(parseJwtInfoErr.GetStatus(), parseJwtInfoErr.GetMsg())
			return
		}

		if acc.IsDisable != 0 {
			authErr := er.NewAppErr(401, er.TokenExpiredError, "Token is expired", nil)
			c.AbortWithStatusJSON(authErr.GetStatus(), authErr.GetMsg())
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
