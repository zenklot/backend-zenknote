package web

type UserRenewPassword struct {
	Password string `json:"password" validate:"required,min=6"`
}
