package mentor

import (
	model "github.com/philipnathan/pijar-backend/internal/mentor/model"
	repository "github.com/philipnathan/pijar-backend/internal/mentor/repository"
	"gorm.io/gorm"

	custom_error "github.com/philipnathan/pijar-backend/internal/mentor/custom_error"
)

type MentorBioServiceInterface interface {
	GetMentorBio(userID *uint) (*model.MentorBiographies, error)
}

type MentorBioService struct {
	repo repository.MentorBioRepositoryInterface
}

func NewMentorBioService(repo repository.MentorBioRepositoryInterface) repository.MentorBioRepositoryInterface {
	return &MentorBioService{
		repo: repo,
	}
}

func (s *MentorBioService) GetMentorBio(userID *uint) (*model.MentorBiographies, error) {
	bio, err := s.repo.GetMentorBio(userID)
	if err == gorm.ErrRecordNotFound {
		return nil, custom_error.ErrMentorBioNotFound
	} else if err != gorm.ErrRecordNotFound && err != nil {
		return nil, err
	}

	return bio, nil
}
