package session_review

import (
	model "github.com/philipnathan/pijar-backend/internal/session_review/model"
	"gorm.io/gorm"
)

type SessionReviewRepositoryInterface interface {
	CreateSessionReview(userID, sessionID, rating *uint, review *string) error
	GetUserReview(userID, sessionID *uint) (*model.SessionReview, error)
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
