package session

import (
	time "time"

	model "github.com/philipnathan/pijar-backend/internal/session/model"
	"gorm.io/gorm"
)

type SessionRepository interface {
	GetSessions(userID uint) ([]model.MentorSession, error)
	GetUpcomingSessions(page, pageSize int) ([]model.MentorSession, int, error)
	GetLearnerHistorySession(userID *uint) (*[]model.MentorSessionParticipant, error)
	GetUpcommingSessionsByCategory(categoryID []uint, page, pageSize int) (*[]model.MentorSession, int, error)
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

func (r *sessionRepository) GetUpcomingSessions(page, pageSize int) ([]model.MentorSession, int, error) {
	var total int64
	countQuery := r.db.Model(&model.MentorSession{}).Where("schedule > ?", time.Now())
	if err := countQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var sessions []model.MentorSession
	err := r.db.Where("schedule > ?", time.Now()).
		Order("schedule ASC").
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&sessions).Error
	if err != nil {
		return nil, 0, err
	}
	return sessions, int(total), nil
}

func (r *sessionRepository) GetLearnerHistorySession(userID *uint) (*[]model.MentorSessionParticipant, error) {
	var sessions []model.MentorSessionParticipant
	err := r.db.Preload("MentorSession").Preload("MentorSession.User").Where("user_id = ?", *userID).Find(&sessions).Error
	if err != nil {
		return nil, err
	}
	return &sessions, nil
}

func (r *sessionRepository) GetUpcommingSessionsByCategory(categoryID []uint, page, pageSize int) (*[]model.MentorSession, int, error) {
	var sessions []model.MentorSession
	var total int64

	countQuery := r.db.Model(&model.MentorSession{}).
		Where("category_id IN ?", categoryID).
		Where("schedule > ?", time.Now())

	if err := countQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.Where("category_id IN ? AND schedule > ?", categoryID, time.Now()).
		Order("schedule ASC").
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&sessions).Error
	if err != nil {
		return nil, 0, err
	}

	return &sessions, int(total), nil
}
