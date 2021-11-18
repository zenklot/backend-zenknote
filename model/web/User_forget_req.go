package web

type UserForgetRequest struct {
	Email string `json:"email" validate:"required,email"`
}
