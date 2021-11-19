package database

import (
	"fmt"

	"github.com/zenklot/backend-zenknote/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() {

	var err error
	DB, err = gorm.Open(postgres.Open(helper.Config("DB_URI")))
	helper.ErrPanic(err)

	fmt.Println("Connection Opened to Database")
	// DB.AutoMigrate(&model.User{}, &model.Note{})
	// fmt.Println("Database Migrated")
}
