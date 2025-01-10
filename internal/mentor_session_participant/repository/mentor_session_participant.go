package mentor_session_participant

import (
	"gorm.io/gorm"

	model "github.com/philipnathan/pijar-backend/internal/mentor_session_participant/model"
)

type MentorSessionParticipantRepositoryInterface interface {
	CreateMentorSessionParticipant(userID, mentorSessionID *uint) error
}

type MentorSessionParticipantRepository struct {
	db *gorm.DB
}

func NewMentorSessionParticipantRepository(db *gorm.DB) MentorSessionParticipantRepositoryInterface {
	return &MentorSessionParticipantRepository{
		db: db,
	}
}

func (r *MentorSessionParticipantRepository) CreateMentorSessionParticipant(userID, MentorSessionID *uint) error {
	participant := model.MentorSessionParticipant{
		UserID:          *userID,
		MentorSessionID: *MentorSessionID,
		Status:          "registered",
	}

	return r.db.Create(&participant).Error
}
