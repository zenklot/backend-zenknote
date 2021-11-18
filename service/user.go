package service

import (
	"errors"

	"github.com/zenklot/backend-zenknote/database"
	"github.com/zenklot/backend-zenknote/model"
)

func GetUserByEmail(e string) (*model.User, error) {
	db := database.DB
	var user model.User
	result := db.Where(&model.User{Email: e}).Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("this email is not registered")
	}
	return &user, nil
}
