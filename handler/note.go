package handler

import (
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
	// fmt.Println(email)
	input := web.NoteCreateRequest{}
	err := c.BodyParser(&input)
	if err != nil {
		return helper.SendErrorResponse(c, fiber.StatusUnprocessableEntity, nil)
	}
	notValid := helper.ValidReq(c, input)
	if notValid != nil {
		return helper.SendErrorResponse(c, fiber.StatusBadRequest, notValid)
	}

	tags := strings.Split(input.Tag, ",")
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
