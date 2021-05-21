package api

import (
	"api-automation-backend/internal/account/repository"
	"api-automation-backend/internal/account/service"
	"api-automation-backend/models/apireq"
	"api-automation-backend/pkg/er"
	"api-automation-backend/pkg/valider"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// @Summary Add Account
// @Produce json
// @Accept json
// @Tags Account
// @Param Body body apireq.AddAccount true "Add Account Request"
// @Success 200 {object} models.Account
// @Failure 400 {object} er.AppErrorMsg "{"code":"400400","message":"Wrong parameter format or invalid"}"
// @Failure 500 {object} er.AppErrorMsg "{"code":"500000","message":"Database unknown error"}"
// @Router /v1/accounts [post]
func AddAccount(c *gin.Context) {
	req := apireq.AddAccount{}
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

	ar := repository.NewMysqlAccountRepo(env.Orm)
	as := service.NewAccountService(ar)

	res, err := as.AddAccount(&req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(200, res)
}
