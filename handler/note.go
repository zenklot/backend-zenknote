package handler

import (
	"fmt"
	"strings"

	"github.com/chilts/sid"
	"github.com/gofiber/fiber/v2"
	"github.com/zenklot/backend-zenknote/helper"
	"github.com/zenklot/backend-zenknote/model"
	"github.com/zenklot/backend-zenknote/model/web"
	"github.com/zenklot/backend-zenknote/service"
)

func GetNotes(c *fiber.Ctx) error {
	email := c.Locals("email")
	var noteData *[]model.Note
	noteData, err := service.GetNotesByEmail(email.(string))
	if err != nil {
		return helper.SendResponse(c, fiber.StatusNoContent, nil)
	}
	response := []web.NotesResponse{}
	var resData web.NotesResponse

	// *** Optional Use :
	// myNote := web.NoteRes{response}
	for _, d := range *noteData {
		resData.Id = d.Id
		titleText := []rune(d.Title)
		if len(titleText) >= 50 {
			resData.Tittle = string(titleText[0:50]) + "..."
		} else {
			resData.Tittle = d.Title
		}
		resData.Tags = d.Tags
		noteText := []rune(d.Note)
		if len(noteText) >= 100 {
			resData.Note = string(noteText[0:100]) + "..."
		} else {
			resData.Note = d.Note
		}
		resData.CreatedAt = d.CreatedAt
		resData.UpdatedAt = d.UpdatedAt
		// *** comment after this line for use optional method
		response = append(response, resData)
		// *** uncomment if you want using this optional method
		// myNote.AddNote(resData)

	}
	// if using AddNote, response must change to myNote.Notes
	return helper.SendResponse(c, fiber.StatusOK, response)
}

func PostNote(c *fiber.Ctx) error {
	email := c.Locals("email")
	input := web.NoteCreateRequest{}
	err := c.BodyParser(&input)
	if err != nil {
		return helper.SendErrorResponse(c, fiber.StatusUnprocessableEntity, nil)
	}
	notValid := helper.ValidReq(c, input)
	if notValid != nil {
		return helper.SendErrorResponse(c, fiber.StatusBadRequest, notValid)
	}

	tags := strings.Split(input.Tags, ",")
	for i, t := range tags {
		tags[i] = strings.ToLower(strings.Replace(t, " ", "", -1))
	}
	tag := strings.Join(tags, ",")
	note := model.Note{
		Email: email.(string),
		Id:    sid.IdBase64(),
		Title: input.Title,
		Note:  input.Note,
		Tags:  tag,
	}

	notes, err := service.CreateNote(&note)
	if err != nil {
		return helper.SendErrorResponse(c, fiber.StatusConflict, helper.StringToSlice(err.Error()))
	}

	response := web.NoteCreateResponse{
		Id: notes.Id,
	}

	return helper.SendResponse(c, fiber.StatusOK, response)
}

func GetNoteById(c *fiber.Ctx) error {
	email := c.Locals("email")
	ids := c.Params("id")
	id := web.NoteRequest{
		Id: ids,
	}
	notValid := helper.ValidReq(c, id)
	if notValid != nil {
		return helper.SendErrorResponse(c, fiber.StatusBadRequest, notValid)
	}

	note, err := service.GetNoteById(id.Id, email.(string))
	if err != nil {
		return helper.SendErrorResponse(c, fiber.StatusBadRequest, helper.StringToSlice(err.Error()))
	}
	response := web.NotesResponse{
		Id:        note.Id,
		Tittle:    note.Title,
		Tags:      note.Tags,
		Note:      note.Note,
		CreatedAt: note.CreatedAt,
		UpdatedAt: note.UpdatedAt,
	}
	return helper.SendResponse(c, fiber.StatusOK, response)
}

func DeleteNote(c *fiber.Ctx) error {
	email := c.Locals("email")
	id := c.Params("id")
	idReq := web.NoteRequest{Id: id}
	notValid := helper.ValidReq(c, idReq)
	if notValid != nil {
		return helper.SendErrorResponse(c, fiber.StatusBadRequest, notValid)
	}

	err := service.DeleteNoteById(idReq.Id, email.(string))
	if err != nil {
		return helper.SendErrorResponse(c, fiber.StatusBadRequest, helper.StringToSlice(err.Error()))
	}

	fmt.Println(err)
	return helper.SendResponse(c, fiber.StatusNoContent, nil)
}

func PutNote(c *fiber.Ctx) error {
	email := c.Locals("email")
	id := c.Params("id")

	input := web.NoteUpdateRequest{}
	err := c.BodyParser(&input)
	if err != nil {
		return helper.SendErrorResponse(c, fiber.StatusUnprocessableEntity, nil)
	}

	notValid := helper.ValidReq(c, input)
	if notValid != nil {
		return helper.SendErrorResponse(c, fiber.StatusBadRequest, notValid)
	}

	tags := strings.Split(input.Tags, ",")
	for i, t := range tags {
		tags[i] = strings.ToLower(strings.Replace(t, " ", "", -1))
	}
	tag := strings.Join(tags, ",")

	newNote := model.Note{
		Title: input.Title,
		Tags:  tag,
		Note:  input.Note,
	}

	upNote, err := service.UpdateNote(&newNote, id, email.(string))
	if err != nil {
		return helper.SendErrorResponse(c, fiber.StatusInternalServerError, nil)
	}

	response := web.NoteUpdateResponse{
		Id:        id,
		Title:     upNote.Title,
		Tags:      upNote.Tags,
		Note:      upNote.Note,
		UpdatedAt: upNote.UpdatedAt,
	}

	return helper.SendResponse(c, fiber.StatusOK, response)

}
