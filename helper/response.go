package helper

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/zenklot/backend-zenknote/model/web"
)

func SendBadRequest(c *fiber.Ctx, err []string) error {
	respon := web.ErrorResponse{
		Code:    fiber.StatusBadRequest,
		Message: utils.StatusMessage(fiber.StatusBadRequest),
		Error:   err,
	}

	return c.Status(fiber.StatusBadRequest).JSON(respon)
}

func SendConflict(c *fiber.Ctx, err []string) error {
	respon := web.ErrorResponse{
		Code:    fiber.StatusConflict,
		Message: utils.StatusMessage(fiber.StatusConflict),
		Error:   err,
	}

	return c.Status(fiber.StatusConflict).JSON(respon)
}

func SendResponse(c *fiber.Ctx, code int, data interface{}) error {
	response := web.SuccessResponse{
		Code:    code,
		Message: utils.StatusMessage(code),
		Data:    data,
	}

	return c.Status(code).JSON(response)
}

func SendErrorResponse(c *fiber.Ctx, code int, err []string) error {
	response := web.ErrorResponse{
		Code:    code,
		Message: utils.StatusMessage(code),
		Error:   err,
	}

	return c.Status(code).JSON(response)
}
