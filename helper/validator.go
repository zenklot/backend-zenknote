package helper

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ValidReq(c *fiber.Ctx, data interface{}) []string {
	validate := validator.New()
	err := validate.Struct(data)
	if err != nil {
		var er []string
		for _, v := range err.(validator.ValidationErrors) {
			er = append(er, "Field validation for '"+v.Field()+"' on the '"+v.ActualTag()+"' tag.")
		}
		return er
	}
	return nil
}
