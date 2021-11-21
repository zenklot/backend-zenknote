package web

import "time"

type NoteUpdateResponse struct {
	Id        string    `json:"id"`
	Title     string    `json:"title"`
	Tags      string    `json:"tags"`
	Note      string    `json:"note"`
	UpdatedAt time.Time `json:"updated_ad"`
}
