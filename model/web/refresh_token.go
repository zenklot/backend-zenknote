package web

type RefreshToken struct {
	Email         string `json:"email" validate:"email"`
	Refresh_Token string `json:"refresh_token" validate:"jwt"`
}
