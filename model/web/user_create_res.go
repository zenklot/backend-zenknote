package web

import "time"

type UserCreateResponse struct {
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"creted_at"`
	UpdatedAt time.Time `json:"update_at"`
}
