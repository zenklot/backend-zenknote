package service

import (
	"errors"

	"github.com/zenklot/backend-zenknote/database"
	"github.com/zenklot/backend-zenknote/model"
)

func GetUserByEmail(e string) (*model.User, error) {

	var user model.User
	result := database.DB.Where(&model.User{Email: e}).Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("this email is not registered")
	}
	return &user, nil
}

func CreateUser(data *model.User) (*model.User, error) {
	result := database.DB.Create(data).Error
	if result != nil {
		return nil, errors.New("email address already registered")
	}
	return data, nil
}

func UpdateUser(data *model.User) (*model.User, error) {
	result := database.DB.Model(&data).Updates(data)
	if result.RowsAffected != 1 {
		return nil, result.Error
	}
	return data, nil
}
