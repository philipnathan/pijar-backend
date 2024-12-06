package database

import (
	"fmt"

	"github.com/philipnathan/pijar-backend/internal/models"
	"gorm.io/gorm"
)

func MigrateDatabase(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Database migrated successfully")
}