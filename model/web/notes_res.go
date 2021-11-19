package web

import "time"

type NotesResponse struct {
	Id        string    `json:"id"`
	Tittle    string    `json:"title"`
	Tags      string    `json:"tags"`
	Note      string    `json:"note"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// optional
type NoteRes struct {
	Notes []NotesResponse
}

func (myNote *NoteRes) AddNote(notes NotesResponse) []NotesResponse {
	myNote.Notes = append(myNote.Notes, notes)
	return myNote.Notes
}
