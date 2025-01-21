package user

import (
	custom_error "github.com/philipnathan/pijar-backend/internal/user/custom_error"
	model "github.com/philipnathan/pijar-backend/internal/user/model"
	repository "github.com/philipnathan/pijar-backend/internal/user/repository"
	utils "github.com/philipnathan/pijar-backend/utils"
	"gorm.io/gorm"
)

type MentorUserServiceInterface interface {
	RegisterMentor(email, password, fullname *string) (string, string, error)
}

type MentorUserService struct {
	repo repository.MentorUserRepositoryInterface
}

func NewMentorUserService(repo repository.MentorUserRepositoryInterface) MentorUserServiceInterface {
	return &MentorUserService{
		repo: repo,
	}
}

func (s *MentorUserService) RegisterMentor(email, password, fullname *string) (string, string, error) {
	var mentor *model.User
	var err error

	mentor, err = s.repo.FindUserByEmail(*email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return "", "", err
	}

	if mentor != nil {
		if mentor.IsMentor != nil && *mentor.IsMentor {
			return "", "", custom_error.ErrAlreadyMentor
		}
		return "", "", custom_error.ErrChangeDetails
	}

	hashedPassword, err := utils.HashPassword(*password)
	if err != nil {
		return "", "", err
	}
	mentor, err = s.repo.CreateNewMentor(email, &hashedPassword, fullname)
	if err != nil {
		return "", "", err
	}

	var access_token string
	if access_token, err = utils.GenerateJWT(mentor.ID, mentor.IsMentor); err != nil {
		return "", "", err
	}

	var refresh_token string
	if refresh_token, err = utils.GenerateRefreshToken(mentor.ID, mentor.IsMentor); err != nil {
		return "", "", err
	}

	return access_token, refresh_token, nil
}
