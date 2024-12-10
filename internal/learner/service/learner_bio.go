package learner

import (
	repo "github.com/philipnathan/pijar-backend/internal/learner/repository"
)

type LearnerBioServiceInterface interface{}

type LearnerBioService struct {
	repo repo.LearnerBioRepositoryInterface
}

func NewLearnerBioService(repo repo.LearnerBioRepositoryInterface) LearnerBioServiceInterface {
	return &LearnerBioService{
		repo: repo,
	}
}
