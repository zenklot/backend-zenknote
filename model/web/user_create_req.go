package web

type UserCreateRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required,min=4"`
	Password string `json:"password" validate:"required,min=6"`
}
