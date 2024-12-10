package learner

import (
	dto "github.com/philipnathan/pijar-backend/internal/learner/dto"
	repo "github.com/philipnathan/pijar-backend/internal/learner/repository"
)

type LearnerBioServiceInterface interface {
	CreateLearnerBio(UserID uint, input *dto.CreateLearnerBioDto) error
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
