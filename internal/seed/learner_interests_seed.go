package seed

import (
	model "github.com/philipnathan/pijar-backend/internal/learner/model"
	"gorm.io/gorm"
)

func SeedLearnerInterests(db *gorm.DB) error {
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	learnerInterests := []model.LearnerInterest{
		{
			UserID:     2,
			CategoryID: 1,
		},
		{
			UserID:     4,
			CategoryID: 4,
		},
		{
			UserID:     6,
			CategoryID: 4,
		},
		{
			UserID:     6,
			CategoryID: 3,
		},
		{
			UserID:     8,
			CategoryID: 1,
		},
		{
			UserID:     8,
			CategoryID: 2,
		},
	}

	var count int64
	if err := db.Model(&model.LearnerInterest{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	for _, learnerInterest := range learnerInterests {
		if err := db.Model(&model.LearnerInterest{}).Create(&learnerInterest).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
