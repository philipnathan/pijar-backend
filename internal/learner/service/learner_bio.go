package learner

import (
	dto "github.com/philipnathan/pijar-backend/internal/learner/dto"
	model "github.com/philipnathan/pijar-backend/internal/learner/model"
	repo "github.com/philipnathan/pijar-backend/internal/learner/repository"
)

type LearnerBioServiceInterface interface {
	CreateLearnerBio(UserID uint, input *dto.CreateLearnerBioDto) error
	GetLearnerBio(UserID uint) (*model.LearnerBio, error)
}

type LearnerBioService struct {
	repo repo.LearnerBioRepositoryInterface
}

func NewLearnerBioService(repo repo.LearnerBioRepositoryInterface) LearnerBioServiceInterface {
	return &LearnerBioService{
		repo: repo,
	}
}

func (s *LearnerBioService) CreateLearnerBio(UserID uint, input *dto.CreateLearnerBioDto) error {
	return s.repo.CreateLearnerBio(UserID, input)
}

func (s *LearnerBioService) GetLearnerBio(UserID uint) (*model.LearnerBio, error) {
	bio, err := s.repo.GetLearnerBio(UserID)
	if err != nil {
		return nil, err
	}

	return bio, nil
}
