package api

import (
	accRepo "api-automation-backend/internal/account/repository"
	"api-automation-backend/internal/token_auth/service"
	"api-automation-backend/models/apireq"
	"api-automation-backend/models/apires"
	"api-automation-backend/pkg/er"
	"api-automation-backend/pkg/valider"

	"github.com/gin-gonic/gin"
)

// @Summary Get JWT Token
// @Produce json
// @Accept json
// @Tags Token
// @Param token body apireq.GetToken true "Get JWT Token Request"
// @Success 200 {object} apires.TokenSuccess
// @Failure 400 {object} er.AppErrorMsg "{"code":"400400","msg":"Wrong parameter format or invalid"}"
// @Failure 401 {object} er.AppErrorMsg "{"code":"400401","msg":"Unauthorized"}"
// @Router /v1/api-token [post]
func GetApiToken(c *gin.Context) {
	req := apireq.GetToken{}
	err := c.Bind(&req)
	if err != nil {
		bindErr := er.NewAppErr(400, er.ErrorParamInvalid, err.Error(), err)
		_ = c.Error(bindErr)
		return
	}

	// 參數驗證
	err = valider.Validate.Struct(req)
	if err != nil {
		err = er.NewAppErr(400, er.ErrorParamInvalid, err.Error(), err)
		_ = c.Error(err)
		return
	}

	ar := accRepo.NewMysqlAccountRepo(env.Orm)
	ts := service.NewTokenService(ar)

	// start auth and generate token
	token, expireAt, err := ts.GenJwtToken(req.AccountId, req.Password)
	if err != nil {
		_ = c.Error(err)
		return
	}

	res := apires.TokenSuccess{
		Token:    token,
		ExpireAt: expireAt,
	}
	c.JSON(200, res)
}
