package database

import (
	"fmt"

	model "github.com/philipnathan/pijar-backend/internal/user/model"
	"gorm.io/gorm"
)

func MigrateDatabase(db *gorm.DB) {
	err := db.AutoMigrate(&model.User{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Database migrated successfully")
}