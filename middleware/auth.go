package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zenklot/backend-zenknote/helper"
)

func Protected(key string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:     []byte(key),
		ErrorHandler:   jwtError,
		SuccessHandler: jwtSuccess,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return helper.SendErrorResponse(c, fiber.StatusBadRequest, helper.StringToSlice("missing or malformed JWT"))
	}
	return helper.SendErrorResponse(c, fiber.StatusUnauthorized, helper.StringToSlice("missing or malformed JWT"))
}

func jwtSuccess(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	var email string
	defer func() {
		if r := recover(); r != nil {
			email = ""
		}
	}()
	email = claims["email"].(string)
	c.Locals("email", email)
	return c.Next()
}
