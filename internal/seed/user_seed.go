package seed

import (
	model "github.com/philipnathan/pijar-backend/internal/user/model"
	repo "github.com/philipnathan/pijar-backend/internal/user/repository"
	"gorm.io/gorm"
)

func SeedUser(db *gorm.DB) error {
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	users := []model.User{
		{
			Email:    "mentor01@example.com",
			Password: "mentor01",
			Fullname: "Mentor01",
		},
		{
			Email:    "learner01@example.com",
			Password: "learner01",
			Fullname: "learner01",
		},
		{
			Email:    "learner02@example.com",
			Password: "learner02",
			Fullname: "learner02",
		},
	}

	var count int64
	if err := db.Model(&model.User{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	for _, user := range users {
		if err := repo.NewUserRepository(db).SaveUser(&user); err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
