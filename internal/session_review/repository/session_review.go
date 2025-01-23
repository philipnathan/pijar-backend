package session_review

import (
	model "github.com/philipnathan/pijar-backend/internal/session_review/model"
	"gorm.io/gorm"
)

type SessionReviewRepositoryInterface interface {
	CreateSessionReview(userID, sessionID, rating *uint, review *string) error
	GetUserReview(userID, sessionID *uint) (*model.SessionReview, error)
	GetSessionReviews(sessionID *uint, page, pageSize *int) (*[]model.SessionReview, int, error)
}

type SessionReviewRepository struct {
	db *gorm.DB
}

func NewSessionReviewRepository(db *gorm.DB) SessionReviewRepositoryInterface {
	return &SessionReviewRepository{
		db: db,
	}
}

func (r *SessionReviewRepository) CreateSessionReview(userID, sessionID, rating *uint, review *string) error {
	rev := model.SessionReview{
		UserID:    *userID,
		SessionID: *sessionID,
		Review:    review,
		Rating:    *rating,
	}

	return r.db.Create(&rev).Error
}

func (r *SessionReviewRepository) GetUserReview(userID, sessionID *uint) (*model.SessionReview, error) {
	var rev model.SessionReview
	err := r.db.Where("user_id = ? AND session_id = ?", userID, sessionID).First(&rev).Error

	if err != nil {
		return nil, err
	}

	return &rev, nil
}

func (r *SessionReviewRepository) GetSessionReviews(sessionID *uint, page, pageSize *int) (*[]model.SessionReview, int, error) {
	var revs []model.SessionReview
	var total int64

	countQuery := r.db.Model(&model.SessionReview{}).Where("session_id = ?", sessionID)
	if err := countQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.Where("session_id = ?", sessionID).
		Preload("User").
		Offset((*page - 1) * *pageSize).
		Limit(*pageSize).
		Find(&revs).Error

	if err != nil {
		return nil, 0, err
	}

	return &revs, int(total), nil
}
