package web

type NoteCreateRequest struct {
	Id    string `json:"id"`
	Title string `json:"title" validate:"required"`
	Tag   string `json:"tag"`
	Note  string `json:"note" validate:"required"`
	Email string `json:"email"`
}
