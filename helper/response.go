package helper

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/zenklot/backend-zenknote/model/web"
)

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
