package learner

import (
	model "github.com/philipnathan/pijar-backend/internal/learner/model"
	"gorm.io/gorm"
)

type LearnerRepositoryInterface interface {
	GetLearnerInterest(userID uint) ([]model.LearnerInterest, error)
	AddLearnerInterests(userID uint, interests []uint) error
	DeleteLearnerInterests(userID uint, interests []uint) error
}

type LearnerRepository struct {
	db *gorm.DB
}

func NewLearnerRepository(db *gorm.DB) *LearnerRepository {
	return &LearnerRepository{
		db: db,
	}
}

func (r *LearnerRepository) GetLearnerInterest(userID uint) ([]model.LearnerInterest, error) {
	var learnerInterests []model.LearnerInterest

	if err := r.db.Preload("Category").Where("user_id = ?", userID).Find(&learnerInterests).Error; err != nil {
		return nil, err
	}

	return learnerInterests, nil
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

func (r *LearnerRepository) DeleteLearnerInterests(userID uint, interests []uint) error {
	var learnerInterests []model.LearnerInterest
	for _, interestID := range interests {
		learnerInterests = append(learnerInterests, model.LearnerInterest{UserID: userID, CategoryID: interestID})
	}

	err := r.db.Where("user_id = ? AND category_id IN ?", userID, interests).Delete(&model.LearnerInterest{}).Error
	if err != nil {
		return err
	}

	return nil
}
