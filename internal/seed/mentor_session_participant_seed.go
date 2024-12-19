package seed

import (
	model "github.com/philipnathan/pijar-backend/internal/session/model"
	"gorm.io/gorm"
)

func MentorSessionParticipant(db *gorm.DB) error {
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	mentorSessionParticipants := []model.MentorSessionParticipant{
		{
			UserID:          2,
			MentorSessionID: 1,
			Status:          "registered",
		},
		{
			UserID:          4,
			MentorSessionID: 4,
			Status:          "registered",
		},
	}

	var count int64
	if err := db.Model(&model.MentorSessionParticipant{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	for _, mentorSessionParticipant := range mentorSessionParticipants {
		if err := db.Create(&mentorSessionParticipant).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
