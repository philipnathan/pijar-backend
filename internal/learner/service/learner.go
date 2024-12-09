package learner

import (
	model "github.com/philipnathan/pijar-backend/internal/learner/model"
	repo "github.com/philipnathan/pijar-backend/internal/learner/repository"
)

type LearnerServiceInterface interface {
	GetLearnerInterests(userID uint) ([]model.LearnerInterest, error)
	AddLearnerInterests(userID uint, interests []uint) error
}

type LearnerService struct {
	repo repo.LearnerRepositoryInterface
}

func NewLearnerService(repo repo.LearnerRepositoryInterface) LearnerServiceInterface {
	return &LearnerService{
		repo: repo,
	}
}

func (s *LearnerService) GetLearnerInterests(userID uint) ([]model.LearnerInterest, error) {
	var learnerInterest []model.LearnerInterest
	learnerInterest, err := s.repo.GetLearnerInterest(userID)
	if err != nil {
		return nil, err
	}
	return learnerInterest, nil
}

func (s *LearnerService) AddLearnerInterests(userID uint, interests []uint) error {
	if err := s.repo.AddLearnerInterests(userID, interests); err != nil {
		return err
	}

	return nil
}
