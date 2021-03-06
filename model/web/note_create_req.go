package web

type NoteCreateRequest struct {
	Id    string `json:"id"`
	Title string `json:"title" validate:"required"`
	Tags  string `json:"tags"`
	Note  string `json:"note"`
	Email string `json:"email"`
}
