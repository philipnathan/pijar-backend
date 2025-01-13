package mentor_session_participant

import (
	"gorm.io/gorm"

	model "github.com/philipnathan/pijar-backend/internal/session/model"
)

type MentorSessionParticipantStatus string

const (
	Registered         MentorSessionParticipantStatus = "registered"
	Confirmed          MentorSessionParticipantStatus = "confirmed"
	CancelledByMentor  MentorSessionParticipantStatus = "cancelled_by_mentor"
	CancelledByLearner MentorSessionParticipantStatus = "cancelled_by_learner"
	Complete           MentorSessionParticipantStatus = "complete"
)

type MentorSessionParticipant struct {
	gorm.Model      `json:"-"`
	UserID          uint                           `gorm:"not null;index" json:"user_id"`
	MentorSessionID uint                           `gorm:"not null;index" json:"mentor_session_id"`
	Status          MentorSessionParticipantStatus `gorm:"type:varchar(20);default:'registered';not null" json:"status"`

	MentorSession model.MentorSession `gorm:"foreignKey:MentorSessionID" json:"mentor_session"`
}
