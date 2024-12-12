package seed

import (
	userModel "github.com/philipnathan/pijar-backend/internal/user/model"
	userRepo "github.com/philipnathan/pijar-backend/internal/user/repository"
	utils "github.com/philipnathan/pijar-backend/utils"
	"gorm.io/gorm"
)

func SeedUser(db *gorm.DB) error {
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	mentor01PW, _ := utils.HashPassword("mentor01")
	learner01PW, _ := utils.HashPassword("learner01")
	learner02PW, _ := utils.HashPassword("learner02")

	users := []userModel.User{
		{
			Email:    "mentor01@example.com",
			Password: mentor01PW,
			Fullname: "Mentor01",
		},
		{
			Email:    "learner01@example.com",
			Password: learner01PW,
			Fullname: "learner01",
		},
		{
			Email:    "learner02@example.com",
			Password: learner02PW,
			Fullname: "learner02",
		},
	}

	var count int64
	if err := db.Model(&userModel.User{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	for _, user := range users {
		if err := userRepo.NewUserRepository(db).SaveUser(&user); err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
