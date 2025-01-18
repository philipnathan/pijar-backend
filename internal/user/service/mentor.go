package user

import (
	custom_error "github.com/philipnathan/pijar-backend/internal/user/custom_error"
	model "github.com/philipnathan/pijar-backend/internal/user/model"
	repository "github.com/philipnathan/pijar-backend/internal/user/repository"
	utils "github.com/philipnathan/pijar-backend/utils"
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

	isExist, err := s.repo.IsUserExist(email)
	if err != nil {
		return "", "", err
	}

	if isExist {
		mentor, err = s.repo.FindUserByEmail(*email)
		if err != nil {
			return "", "", err
		}

		if *mentor.IsMentor {
			return "", "", custom_error.ErrAlreadyMentor
		}

		err = utils.ComparePassword(mentor.Password, *password)
		if err != nil {
			return "", "", custom_error.ErrWrongPasswordAndLearnerRegistered
		}

		mentor, err = s.repo.SetIsMentorToTrue(email)
		if err != nil {
			return "", "", err
		}
	} else {
		hashedPassword, err := utils.HashPassword(*password)
		if err != nil {
			return "", "", err
		}
		mentor, err = s.repo.CreateNewMentor(email, &hashedPassword, fullname)
		if err != nil {
			return "", "", err
		}
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
