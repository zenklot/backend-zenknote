package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zenklot/backend-zenknote/helper"
	"github.com/zenklot/backend-zenknote/model/web"
	"github.com/zenklot/backend-zenknote/service"
)

func GetProfile(c *fiber.Ctx) error {
	email := c.Locals("email")

	userData, err := service.GetUserByEmail(email.(string))
	if err != nil {
		return helper.SendErrorResponse(c, fiber.StatusInternalServerError, nil)
	}

	userNotes, _ := service.GetNotesByEmail(email.(string))
	var notes int
	if userNotes == nil {
		notes = 0	
	}else{
		notes = len(*userNotes)	
	}

	response := web.ProfileResponse{
		Email:     userData.Email,
		Name:      userData.Name,
		Notes:     notes,
		CreatedAt: userData.CreatedAt,
		UpdatedAt: userData.UpdatedAt,
	}
	return helper.SendResponse(c, fiber.StatusOK, response)
}

func PutProfile(c *fiber.Ctx) error {
	passHash := ""
	email := c.Locals("email")

	input := web.ProfileRequest{}

	err := c.BodyParser(&input)
	if err != nil {
		return helper.SendErrorResponse(c, fiber.StatusInternalServerError, nil)
	}

	if input.Password == "" {
		input.Password = "_12qwdef4ULT"
	}
	notValid := helper.ValidReq(c, &input)
	if notValid != nil {
		return helper.SendErrorResponse(c, fiber.StatusBadRequest, notValid)
	}

	user, err := service.GetUserByEmail(email.(string))
	if err != nil {
		return helper.SendErrorResponse(c, fiber.StatusBadRequest, helper.StringToSlice(err.Error()))
	}

	if input.OldPassword != "" {
		passwordOld := helper.CheckPasswordHash(user.Password, input.OldPassword)
		if !passwordOld {
			return helper.SendErrorResponse(c, fiber.StatusBadRequest, helper.StringToSlice("your current password is wrong"))
		}
	}

	if input.Password != "" && input.OldPassword != "" && input.Password != "_12qwdef4ULT" {
		passHash, err = helper.HashPassword(input.Password)
		if err != nil {
			return helper.SendErrorResponse(c, fiber.StatusInternalServerError, nil)
		}
		user.Password = passHash
	}

	user.Name = input.Name

	userUpdate, err := service.UpdateUser(user)
	if err != nil {
		return helper.SendErrorResponse(c, fiber.StatusInternalServerError, nil)
	}

	userNotes, err := service.GetNotesByEmail(email.(string))
	if err != nil {
		return helper.SendErrorResponse(c, fiber.StatusBadRequest, helper.StringToSlice(err.Error()))
	}
	notes := len(*userNotes)

	response := web.ProfileResponse{
		Email:     userUpdate.Email,
		Name:      userUpdate.Name,
		Notes:     notes,
		CreatedAt: userUpdate.CreatedAt,
		UpdatedAt: userUpdate.UpdatedAt,
	}

	return helper.SendResponse(c, fiber.StatusOK, response)

}
