package apireq

type GetToken struct {
	AccountId int64  `json:"account_id" validate:"required,numeric" example:"1"`
	Password  string `json:"password" validate:"required,min=8,max=16" example:"12345678"`
}
