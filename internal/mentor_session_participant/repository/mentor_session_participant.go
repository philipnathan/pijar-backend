package mentor_session_participant

import (
	"gorm.io/gorm"

	model "github.com/philipnathan/pijar-backend/internal/mentor_session_participant/model"
)

type MentorSessionParticipantRepositoryInterface interface {
	CreateMentorSessionParticipant(userID, mentorSessionID *uint) error
	GetMentorSessionParticipant(userID, mentorSessionID *uint) (*model.MentorSessionParticipant, error)
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

func (r *MentorSessionParticipantRepository) GetMentorSessionParticipant(userID, mentorSessionID *uint) (*model.MentorSessionParticipant, error) {
	var participant model.MentorSessionParticipant
	if err := r.db.Where("user_id = ? AND mentor_session_id = ?", *userID, *mentorSessionID).First(&participant).Error; err != nil {
		return nil, err
	}

	return &participant, nil
}
