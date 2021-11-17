package helper

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/zenklot/backend-zenknote/model/web"
)

func ValidUserRegister(user *web.UserCreateRequest, c *fiber.Ctx) []string {
	validate := validator.New()
	e := validate.Struct(user)
	if e != nil {
		var er []string
		for _, v := range e.(validator.ValidationErrors) {
			er = append(er, "Field validation for '"+v.Field()+"' on the '"+v.ActualTag()+"' tag.")
		}
		return er
	}
	return nil
}
