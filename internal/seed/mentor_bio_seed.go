package seed

import (
	model "github.com/philipnathan/pijar-backend/internal/mentor/model"
	repo "github.com/philipnathan/pijar-backend/internal/mentor/repository"
	"gorm.io/gorm"
)

func SeedMentorBio(db *gorm.DB) error {
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	mentors := []model.MentorBiographies{
		{UserID: 1, Bio: "I am a software engineer."},
	}

	var count int64
	if err := db.Model(&model.MentorBiographies{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	for _, mentor := range mentors {
		if err := repo.NewMentorBioRepository(db).SaveMentorBio(&mentor); err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
