package database

import (
	"fmt"

	"github.com/zenklot/backend-zenknote/helper"
	"github.com/zenklot/backend-zenknote/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() {

	var err error
	DB, err = gorm.Open(postgres.Open(helper.Config("DB_URI")))
	helper.ErrPanic(err)

	fmt.Println("Connection Opened to Database")
	DB.AutoMigrate(&model.User{})
	fmt.Println("Database Migrated")
}
