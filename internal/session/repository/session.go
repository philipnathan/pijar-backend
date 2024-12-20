package session

import (
	time "time"

	model "github.com/philipnathan/pijar-backend/internal/session/model"
	"gorm.io/gorm"
)

type SessionRepository interface {
	GetSessions(userID uint) ([]model.MentorSession, error)
	FetchUpcomingSessions(userID uint) ([]model.MentorSession, error)
	GetUpcomingSessions() ([]model.MentorSession, error)
	GetLearnerHistorySession(userID *uint) (*[]model.MentorSessionParticipant, error)
}

type sessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) SessionRepository {
	return &sessionRepository{db: db}
}

func (r *sessionRepository) GetSessions(userID uint) ([]model.MentorSession, error) {
	var sessions []model.MentorSession
	err := r.db.Preload("User").Where("user_id = ?", userID).Find(&sessions).Error
	if err != nil {
		return nil, err
	}
	return sessions, nil
}

func (r *sessionRepository) FetchUpcomingSessions(userID uint) ([]model.MentorSession, error) {
	var sessions []model.MentorSession
	err := r.db.Preload("User").Where("user_id = ? AND schedule > ?", userID, time.Now()).Find(&sessions).Error

	if err != nil {
		return nil, err
	}
	return sessions, nil
}

func (r *sessionRepository) GetUpcomingSessions() ([]model.MentorSession, error) {
	var sessions []model.MentorSession
	err := r.db.Where("schedule > ?", time.Now()).Find(&sessions).Error
	if err != nil {
		return nil, err
	}
	return sessions, nil
}
func (r *sessionRepository) GetLearnerHistorySession(userID *uint) (*[]model.MentorSessionParticipant, error) {
	var sessions []model.MentorSessionParticipant
	err := r.db.Preload("MentorSession").Preload("MentorSession.User").Where("user_id = ?", *userID).Find(&sessions).Error
	if err != nil {
		return nil, err
	}
	return &sessions, nil
}
