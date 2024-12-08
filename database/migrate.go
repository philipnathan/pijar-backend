package database

import (
	"fmt"
	"os/user"

	category "github.com/philipnathan/pijar-backend/internal/category/model"
	user "github.com/philipnathan/pijar-backend/internal/user/model"
	"gorm.io/gorm"
)

func MigrateDatabase(db *gorm.DB) {
	err := db.AutoMigrate(&user.User{}, &category.Category{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Database migrated successfully")
}
