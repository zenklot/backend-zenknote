package service

import (
	"errors"

	"github.com/zenklot/backend-zenknote/database"
	"github.com/zenklot/backend-zenknote/model"
)

func GetNotesByEmail(e string) (*[]model.Note, error) {
	note := []model.Note{}
	result := database.DB.Where(&model.Note{Email: e}).Find(&note)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("this email is not registered")
	}
	return &note, nil
}

func CreateNote(data *model.Note) (*model.Note, error) {
	result := database.DB.Create(data).Error
	if result != nil {
		return nil, errors.New(result.Error())
	}
	return data, nil
}

func GetNoteById(id string, email string) (*model.Note, error) {
	note := model.Note{}
	result := database.DB.Where(&model.Note{Email: email, Id: id}).Find(&note)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("please check your note id")
	}
	return &note, nil
}

func DeleteNoteById(id string, email string) error {
	note := model.Note{}
	result := database.DB.Where(&model.Note{Email: email, Id: id}).Delete(&note)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("please check your note id")
	}
	return nil
}
