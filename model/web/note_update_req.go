package web

type NoteUpdateRequest struct {
	Title string `json:"title" validate:"required"`
	Tags  string `json:"tags"`
	Note  string `json:"note"`
}
