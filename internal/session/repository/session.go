package session

import (
	time "time"

	model "github.com/philipnathan/pijar-backend/internal/session/model"
	"gorm.io/gorm"
)

type SessionRepository interface {
	GetUpcomingSessions(page, pageSize int) (*[]model.MentorSession, int, error)
	GetLearnerHistorySession(userID *uint) (*[]model.MentorSessionParticipant, error)
	GetUpcommingSessionsByCategory(categoryID []uint, page, pageSize int) (*[]model.MentorSession, int, error)
	GetAllSessionsWithFilter(categoryID, mentorID uint, page, pageSize int, rating, schedule string) (*[]model.MentorSession, int, error)
	GetSessionByID(sessionID uint) (*model.MentorSession, error)
	GetSessionDetailByID(sessionID uint) (*model.MentorSession, error)
}

type sessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) SessionRepository {
	return &sessionRepository{db: db}
}

func (r *sessionRepository) GetUpcomingSessions(page, pageSize int) (*[]model.MentorSession, int, error) {
	var total int64
	countQuery := r.db.Model(&model.MentorSession{}).Where("schedule > ?", time.Now())
	if err := countQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var sessions []model.MentorSession
	err := r.db.Where("schedule > ?", time.Now()).
		Preload("User").
		Order("schedule ASC").
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&sessions).Error
	if err != nil {
		return nil, 0, err
	}
	return &sessions, int(total), nil
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
		Preload("User").
		Order("schedule ASC").
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&sessions).Error
	if err != nil {
		return nil, 0, err
	}

	return &sessions, int(total), nil
}

func (r *sessionRepository) GetAllSessionsWithFilter(categoryID, mentorID uint, page, pageSize int, rating, schedule string) (*[]model.MentorSession, int, error) {
	var sessions []model.MentorSession
	var total int64

	countQuery := r.db.Model(&model.MentorSession{})
	query := r.db

	if mentorID > 0 {
		countQuery = countQuery.Where("user_id = ?", mentorID)
		query = query.Where("user_id = ?", mentorID)
	}
	if categoryID > 0 {
		countQuery = countQuery.Where("category_id = ?", categoryID)
		query = query.Where("category_id = ?", categoryID)
	}

	if schedule != "" {
		if schedule == "newest" {
			query = query.Order("schedule DESC")
		} else if schedule == "oldest" {
			query = query.Order("schedule ASC")
		}
	}

	if err := countQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.
		Preload("User").
		Preload("SessionReviews").
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&sessions).Error

	if err != nil {
		return nil, 0, err
	}

	return &sessions, int(total), nil
}

func (r *sessionRepository) GetSessionByID(sessionID uint) (*model.MentorSession, error) {
	err := r.db.Where("id = ?", sessionID).First(&model.MentorSession{}).Error

	if err != nil {
		return nil, err
	}

	return &model.MentorSession{}, nil
}

func (r *sessionRepository) GetSessionDetailByID(sessionID uint) (*model.MentorSession, error) {
	// Difference between GetSessionDetailByID & Get SessionByID is that GetSessionDetailByID will include user info
	var session model.MentorSession
	err := r.db.Preload("User").
		Preload("SessionReviews").
		Where("id = ?", sessionID).
		First(&session).Error

	if err != nil {
		return nil, err
	}

	return &session, nil
}
