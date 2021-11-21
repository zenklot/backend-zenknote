package web

type ProfileRequest struct {
	Name        string `json:"name" validate:"required"`
	OldPassword string `json:"old_password"`
	Password    string `json:"password" validate:"required,min=6"`
}
