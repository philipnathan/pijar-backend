package mentor

import (
	repo "github.com/philipnathan/pijar-backend/internal/mentor/repository"
	userModel "github.com/philipnathan/pijar-backend/internal/user/model"
)

type MentorServiceInterface interface {
	GetMentorDetails(MentorID uint) (*userModel.User, error)
}

type MentorService struct {
	repo repo.MentorRepositoryInterface
}

func NewMentorService(repo repo.MentorRepositoryInterface) MentorServiceInterface {
	return &MentorService{
		repo: repo,
	}
}

func (s *MentorService) GetMentorDetails(MentorID uint) (*userModel.User, error) {
	mentor, err := s.repo.GetMentorDetails(&MentorID)
	if err != nil {
		return nil, err
	}

	return mentor, nil
}
