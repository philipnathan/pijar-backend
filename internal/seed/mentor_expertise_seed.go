package seed

import (
	model "github.com/philipnathan/pijar-backend/internal/mentor/model"
	repo "github.com/philipnathan/pijar-backend/internal/mentor/repository"
	"gorm.io/gorm"
)

func SeedMentorExpertise(db *gorm.DB) error {
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	mentors := []model.MentorExpertises{
		{
			UserID:      1,
			Expertise:   "Software Engineer",
			Category_id: 1,
		},
		{
			UserID:      1,
			Expertise:   "Software Engineer",
			Category_id: 2,
		},
	}

	var count int64
	if err := db.Model(&model.MentorExpertises{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	for _, mentor := range mentors {
		if err := repo.NewMentorExpertiseRepository(db).SaveMentorExpertise(&mentor); err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
