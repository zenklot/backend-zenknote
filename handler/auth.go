package handler

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zenklot/backend-zenknote/database"
	"github.com/zenklot/backend-zenknote/helper"
	"github.com/zenklot/backend-zenknote/model"
	"github.com/zenklot/backend-zenknote/model/web"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func getUserByEmail(e string) (*model.User, error) {
	db := database.DB
	var user model.User
	err := db.Where(&model.User{Email: e}).Find(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil

}

func CheckPasswordHash(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func PostRegister(c *fiber.Ctx) error {

	db := database.DB
	reqUser := new(web.UserCreateRequest)
	err := c.BodyParser(reqUser)
	if err != nil {
		return helper.SendErrorResponse(c, fiber.StatusUnprocessableEntity, nil)
	}

	noValid := helper.ValidUserRegister(reqUser, c)
	if noValid != nil {
		return helper.SendErrorResponse(c, fiber.StatusBadRequest, noValid)
	}

	hash, _ := hashPassword(reqUser.Password)
	user := new(model.User)
	user.Password = hash
	user.Name = reqUser.Name
	user.Email = reqUser.Email

	result := db.Create(&user).Error
	if result != nil {
		return helper.SendErrorResponse(c, fiber.StatusConflict, helper.StringToSlice("Email address has been registered!"))
	}

	newUser := &web.UserCreateResponse{
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return helper.SendResponse(c, fiber.StatusCreated, newUser)

}

func PostLogin(c *fiber.Ctx) error {
	type LoginInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var input LoginInput
	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on login request", "data": err})
	}

	email := input.Email
	pass := input.Password

	smail, err := getUserByEmail(email)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Error on email", "data": err})
	}

	if smail == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "User not found", "data": err})
	}

	if !CheckPasswordHash(smail.Password, pass) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Salah password", "data": err})
	}
	// er := bcrypt.CompareHashAndPassword([]byte(smail.Password), []byte(pass))
	// fmt.Println(er)
	// if er != nil {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid password", "data": err})
	// }

	type TokenClaims struct {
		Email string `json:"email"`
		jwt.RegisteredClaims
	}

	// Create the Claims
	claims := TokenClaims{
		smail.Email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 30)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(helper.Config("KUNCI_RAHASIA")))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid password jwt", "data": err})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Success login", "data": tokenString})
}
