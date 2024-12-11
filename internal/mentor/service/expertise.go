package mentor

import (
	repo "github.com/philipnathan/pijar-backend/internal/mentor/repository"
)

type MentorExpertiseServiceInterface interface{}

type MentorExpertiseService struct {
	repo repo.MentorExpertiseRepositoryInterface
}

func NewMentorExpertiseService(repo repo.MentorExpertiseRepositoryInterface) MentorExpertiseServiceInterface {
	return &MentorExpertiseService{
		repo: repo,
	}
}
