package web

type UserLoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
