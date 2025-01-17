package seed

import (
	model "github.com/philipnathan/pijar-backend/internal/mentor_session_participant/model"
	"gorm.io/gorm"
)

func MentorSessionParticipant(db *gorm.DB) error {
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	mentorSessionParticipants := []model.MentorSessionParticipant{
		{
			UserID:          1,
			MentorSessionID: 5,
			Status:          "complete",
		},
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
		{
			UserID:          6,
			MentorSessionID: 4,
			Status:          "registered",
		},
		{
			UserID:          8,
			MentorSessionID: 4,
			Status:          "registered",
		},
		{
			UserID:          2,
			MentorSessionID: 7,
			Status:          "registered",
		},
		{
			UserID:          2,
			MentorSessionID: 8,
			Status:          "registered",
		},
		{
			UserID:          2,
			MentorSessionID: 9,
			Status:          "registered",
		},
		{
			UserID:          2,
			MentorSessionID: 10,
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
