package apireq

type AddAccount struct {
	Account  string `json:"account" validate:"required,max=64" example:"admin"`
	Password string `json:"password" validate:"required,min=8,max=16" example:"12345678"`
	Name     string `json:"name" validate:"required,max=64" example:"administrator"`
}
