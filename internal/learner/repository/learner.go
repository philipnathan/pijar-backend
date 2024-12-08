package learner

import (
	model "github.com/philipnathan/pijar-backend/internal/learner/model"
	"gorm.io/gorm"
)

type LearnerRepositoryInterface interface {
	GetLearnerInterest(userID uint) ([]model.LearnerInterest, error)
	AddLearnerInterests(userID uint, interests []uint) error
}

type LearnerRepository struct {
	db *gorm.DB
}

func NewLearnerRepository(db *gorm.DB) LearnerRepositoryInterface {
	return &LearnerRepository{
		db: db,
	}
}

func (r *LearnerRepository) GetLearnerInterest(userID uint) ([]model.LearnerInterest, error) {
	var learnerInterest []model.LearnerInterest

	err := r.db.Where("user_id = ?", userID).Find(&learnerInterest).Error
	if err != nil {
		return nil, err
	}
	return learnerInterest, nil
}

func (r *LearnerRepository) AddLearnerInterests(userID uint, interests []uint) error {

	var learnerInterests []model.LearnerInterest
	for _, interestID := range interests {
		learnerInterests = append(learnerInterests, model.LearnerInterest{UserID: userID, CategoryID: interestID})
	}

	err := r.db.CreateInBatches(learnerInterests, 100).Error
	if err != nil {
		return err
	}

	return nil
}
