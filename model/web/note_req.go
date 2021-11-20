package web

type NoteRequest struct {
	Id string `json:"id" validate:"required"`
}
