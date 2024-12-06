package services

import (
	"github.com/philipnathan/pijar-backend/custom_error"
	"github.com/philipnathan/pijar-backend/internal/models"
	"github.com/philipnathan/pijar-backend/internal/repositories"
	"github.com/philipnathan/pijar-backend/utils"
	"gorm.io/gorm"
)

type UserService struct {
	repo repositories.UserRepository
}


func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) RegisterUserService(user *models.User) (error) {
	var err error

	if exist, err := s.isUserExist(user); err != nil  {
		return err } else if exist {
			return err
		}

	user.Password, err = utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	err = s.repo.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) isUserExist(user *models.User) (bool, error) {
	if exist, err := s.repo.FindUserByEmail(user.Email); err == nil || exist != nil {
		return true, custom_error.ErrEmailExist
	}

	if  exist, err := s.repo.FindByPhoneNumber(user.PhoneNumber); err == nil || exist != nil {
		return true, custom_error.ErrPhoneNumberExist
	}

	return false, nil
}

func (s *UserService) LoginUserService(email, password string) (string, string, error) {
	user, err := s.repo.FindUserByEmail(email)
	if err != nil {
		return "", "", custom_error.ErrLogin
	}

	if err := utils.ComparePassword(user.Password, password); err != nil {
		return "", "", custom_error.ErrLogin
	}

	var access_token string
	if access_token, err = utils.GenerateJWT(user.ID, user.IsMentor); err != nil {
		return "", "", err
	}

	var refresh_token string
	if refresh_token, err = utils.GenerateRefreshToken(user.ID, user.IsMentor); err != nil {
		return "", "", err
	}



	return access_token, refresh_token,nil
}

func (s *UserService) GetUserDetails(userID uint) (*models.User, error) {
	user, err := s.repo.FindByUserId(userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, custom_error.ErrUserNotFound
		}
		return nil, err
	}

	return user, nil
}