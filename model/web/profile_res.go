package web

import "time"

type ProfileResponse struct {
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Notes     int       `json:"notes"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
