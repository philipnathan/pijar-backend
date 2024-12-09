package learner

import (
	model "github.com/philipnathan/pijar-backend/internal/learner/model"
	repo "github.com/philipnathan/pijar-backend/internal/learner/repository"
)

type LearnerServiceInterface interface {
	GetLearnerInterests(userID uint) ([]model.LearnerInterest, error)
	AddLearnerInterests(userID uint, interests []uint) error
	DeleteLearnerInterests(userID uint, interests []uint) error
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
	learnerInterests, err := s.repo.GetLearnerInterest(userID)
	if err != nil {
		return err
	}

	existingInterests := make(map[uint]bool)
	for _, data := range learnerInterests {
		existingInterests[data.Category.ID] = true
	}

	var newInterests []uint
	for _, categoryID := range interests {
		if !existingInterests[categoryID] {
			newInterests = append(newInterests, categoryID)
		}
	}

	if err := s.repo.AddLearnerInterests(userID, newInterests); err != nil {
		return err
	}

	return nil
}

func (s *LearnerService) DeleteLearnerInterests(userID uint, interests []uint) error {
	learnerInterests, err := s.repo.GetLearnerInterest(userID)
	if err != nil {
		return err
	}

	existingInterests := make(map[uint]bool)
	for _, data := range learnerInterests {
		existingInterests[data.Category.ID] = true
	}

	var deleteInterests []uint
	for _, categoryID := range interests {
		if existingInterests[categoryID] {
			deleteInterests = append(deleteInterests, categoryID)
		}
	}

	if err := s.repo.DeleteLearnerInterests(userID, deleteInterests); err != nil {
		return err
	}

	return nil
}
