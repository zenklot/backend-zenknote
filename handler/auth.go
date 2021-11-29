package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zenklot/backend-zenknote/helper"
	"github.com/zenklot/backend-zenknote/model"
	"github.com/zenklot/backend-zenknote/model/web"
	"github.com/zenklot/backend-zenknote/service"
)

func PostRegister(c *fiber.Ctx) error {
	reqUser := new(web.UserCreateRequest)
	err := c.BodyParser(reqUser)
	if err != nil {
		return helper.SendErrorResponse(c, fiber.StatusUnprocessableEntity, nil)
	}

	notValid := helper.ValidReq(c, reqUser)
	if notValid != nil {
		return helper.SendErrorResponse(c, fiber.StatusBadRequest, notValid)
	}

	hash, _ := helper.HashPassword(reqUser.Password)
	user := new(model.User)
	user.Password = hash
	user.Name = reqUser.Name
	user.Email = reqUser.Email

	user, err = service.CreateUser(user)
	if err != nil {
		return helper.SendErrorResponse(c, fiber.StatusConflict, helper.StringToSlice(err.Error()))
	}

	response := &web.UserCreateResponse{
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return helper.SendResponse(c, fiber.StatusCreated, response)

}

func PostLogin(c *fiber.Ctx) error {
	input := web.UserLoginRequest{}
	err := c.BodyParser(&input)
	if err != nil {
		return helper.SendErrorResponse(c, fiber.StatusUnprocessableEntity, nil)
	}

	notValid := helper.ValidReq(c, input)
	if notValid != nil {
		return helper.SendErrorResponse(c, fiber.StatusBadRequest, notValid)
	}

	email := input.Email
	pass := input.Password

	userData, err := service.GetUserByEmail(email)
	if err != nil {
		return helper.SendErrorResponse(c, fiber.StatusUnauthorized, helper.StringToSlice(err.Error()))
	}

	if !helper.CheckPasswordHash(userData.Password, pass) {
		return helper.SendErrorResponse(c, fiber.StatusUnauthorized, helper.StringToSlice("Password Wrong!"))
	}

	tokenClaims := web.TokenClaims{
		Email: userData.Email,
	}

	tokenClaims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Minute * 30))
	tokenString, err := helper.CreateJWT(tokenClaims, helper.Config("KUNCI_RAHASIA"))
	if err != nil {
		return helper.SendErrorResponse(c, fiber.StatusInternalServerError, nil)
	}
	refTokenClaims := web.RefreshTokenClaims{}
	refTokenClaims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Hour * 24))
	refTokenString, err := helper.CreateJWT(refTokenClaims, helper.Config("KUNCI_REFRESH"))
	if err != nil {
		return helper.SendErrorResponse(c, fiber.StatusInternalServerError, nil)
	}
	response := web.UserLoginResponse{
		Token:        tokenString,
		RefreshToken: refTokenString,
	}

	return helper.SendResponse(c, fiber.StatusOK, response)
}

func PostForgetPassword(c *fiber.Ctx) error {
	input := web.UserForgetRequest{}

	err := c.BodyParser(&input)
	if err != nil {
		return helper.SendErrorResponse(c, fiber.StatusUnprocessableEntity, nil)
	}
	notValid := helper.ValidReq(c, input)
	if notValid != nil {
		return helper.SendErrorResponse(c, fiber.StatusBadRequest, notValid)
	}

	email := input.Email
	userData, err := service.GetUserByEmail(email)
	if err != nil {
		return helper.SendErrorResponse(c, fiber.StatusUnauthorized, helper.StringToSlice(err.Error()))
	}

	tokenClaim := web.TokenClaims{
		Email: userData.Email,
	}
	tokenClaim.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Minute * 30))

	token, err := helper.CreateJWT(tokenClaim, helper.Config("KUNCI_RAHASIA"))
	if err != nil {
		return helper.SendErrorResponse(c, fiber.StatusInternalServerError, nil)
	}
	tokenString := jwt.EncodeSegment([]byte(token))
	err = helper.SendMail(email, "Forget Password", "To Renew Your Password Open This Link : <br > http://zenk-note.netlify.app/repassword/"+tokenString)
	if err != nil {
		return helper.SendErrorResponse(c, fiber.StatusBadGateway, helper.StringToSlice(err.Error()))
	}
	return helper.SendResponse(c, fiber.StatusOK, "Forget Password has been sent to "+email)
}

func PutForgetPassword(c *fiber.Ctx) error {
	// db := database.DB
	inputToken := c.Query("token")
	if inputToken == "" {
		return helper.SendErrorResponse(c, fiber.StatusUnprocessableEntity, nil)
	}

	claims, err := helper.ValidateJWT(inputToken, helper.Config("KUNCI_RAHASIA"))
	if err != nil {
		return helper.SendErrorResponse(c, fiber.StatusBadRequest, helper.StringToSlice(err.Error()))
	}
	if claims == nil {
		return helper.SendErrorResponse(c, fiber.StatusBadRequest, helper.StringToSlice("Please Check your token, Token invalid"))
	}
	email := claims.Email
	userData, err := service.GetUserByEmail(email)
	if err != nil {
		return helper.SendErrorResponse(c, fiber.StatusUnauthorized, helper.StringToSlice(err.Error()))
	}

	inputPassword := web.UserRenewPassword{}
	err = c.BodyParser(&inputPassword)
	if err != nil {
		return helper.SendErrorResponse(c, fiber.StatusUnprocessableEntity, nil)
	}
	notValid := helper.ValidReq(c, inputPassword)
	if notValid != nil {
		return helper.SendErrorResponse(c, fiber.StatusBadRequest, notValid)
	}
	passHash, _ := helper.HashPassword(inputPassword.Password)
	// result := db.Model(&userData).Update("password", passHash)
	// if result.RowsAffected != 1 {
	// 	return helper.SendErrorResponse(c, fiber.StatusInternalServerError, nil)
	// }
	userData.Password = passHash
	_, err = service.UpdateUser(userData)
	if err != nil {
		return helper.SendErrorResponse(c, fiber.StatusInternalServerError, nil)
	}
	return helper.SendResponse(c, fiber.StatusOK, email)
}

func PostRefresh(c *fiber.Ctx) error {
	input := web.RefreshToken{}
	err := c.BodyParser(&input)
	if err != nil {
		return helper.SendErrorResponse(c, fiber.StatusUnprocessableEntity, nil)
	}
	notValid := helper.ValidReq(c, input)
	if notValid != nil {
		return helper.SendErrorResponse(c, fiber.StatusBadRequest, notValid)
	}

	inputToken := input.Refresh_Token
	claims, err := helper.ValidateJWT(inputToken, helper.Config("KUNCI_REFRESH"))
	if err != nil {
		return helper.SendErrorResponse(c, fiber.StatusBadRequest, helper.StringToSlice(err.Error()))
	}
	if claims == nil {
		return helper.SendErrorResponse(c, fiber.StatusBadRequest, helper.StringToSlice("Please Check your token, Token invalid"))
	}
	email := input.Email
	userData, err := service.GetUserByEmail(email)
	if err != nil {
		return helper.SendErrorResponse(c, fiber.StatusUnauthorized, helper.StringToSlice(err.Error()))
	}

	tokenClaims := web.TokenClaims{
		Email: userData.Email,
	}
	tokenClaims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Minute * 30))
	tokenString, err := helper.CreateJWT(tokenClaims, helper.Config("KUNCI_RAHASIA"))
	if err != nil {
		return helper.SendErrorResponse(c, fiber.StatusInternalServerError, nil)
	}
	refTokenClaims := web.RefreshTokenClaims{}
	refTokenClaims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Hour * 24))
	refTokenString, err := helper.CreateJWT(refTokenClaims, helper.Config("KUNCI_REFRESH"))
	if err != nil {
		return helper.SendErrorResponse(c, fiber.StatusInternalServerError, nil)
	}
	response := web.UserLoginResponse{
		Token:        tokenString,
		RefreshToken: refTokenString,
	}

	return helper.SendResponse(c, fiber.StatusOK, response)
}
