package user

import (
	"os"

	"github.com/gin-gonic/gin"
	custom_error "github.com/philipnathan/pijar-backend/internal/user/custom_error"
	model "github.com/philipnathan/pijar-backend/internal/user/model"
	repo "github.com/philipnathan/pijar-backend/internal/user/repository"
	"github.com/philipnathan/pijar-backend/utils"
	"gorm.io/gorm"
)

type GoogleAuthServiceInterface interface {
	GoogleRegister(c *gin.Context, email, fullname *string, entity string) (string, string, error)
	GoogleLogin(email, entity *string) (string, string, error)
}

type GoogleAuthService struct {
	repo repo.GoogleAuthRepoInterface
}

func NewGoogleAuthService(repo repo.GoogleAuthRepoInterface) GoogleAuthServiceInterface {
	return &GoogleAuthService{
		repo: repo,
	}
}

func (s *GoogleAuthService) GoogleRegister(c *gin.Context, email, fullname *string, entity string) (string, string, error) {
	var user *model.User
	var err error

	defaultPassword := os.Getenv("SECRET_PASSWORD")
	hashedDefaultPassword, err := utils.HashPassword(defaultPassword)
	if err != nil {
		return "", "", err
	}

	user, err = s.repo.FindUserByEmail(email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return "", "", err
	}

	if err == gorm.ErrRecordNotFound {
		if entity == "learner" {
			user, err = s.repo.CreateLearner(
				email,
				&hashedDefaultPassword,
				fullname)
		} else if entity == "mentor" {
			user, err = s.repo.CreateMentor(
				email,
				&hashedDefaultPassword,
				fullname)
		}
	} else {
		if entity == "learner" {
			if user.IsLearner {
				return "", "", custom_error.ErrAlreadyLearner
			}
			user, err = s.repo.SetIsLearnerToTrue(email)
		} else if entity == "mentor" {
			if *user.IsMentor {
				return "", "", custom_error.ErrAlreadyMentor
			}
			user, err = s.repo.SetIsMentorToTrue(email)
		}
	}

	if err != nil {
		return "", "", err
	}

	var access_token string
	access_token, err = utils.GenerateJWT(user.ID, user.IsMentor)
	if err != nil {
		return "", "", err
	}

	var refresh_token string
	refresh_token, err = utils.GenerateRefreshToken(user.ID, user.IsMentor)
	if err != nil {
		return "", "", err
	}

	return access_token, refresh_token, nil
}

func (s *GoogleAuthService) GoogleLogin(email, entity *string) (string, string, error) {
	user, err := s.repo.FindUserByEmail(email)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", "", custom_error.ErrUserNotFound
		}

		return "", "", err
	}

	if user.AuthProvider != "google" {
		return "", "", custom_error.ErrNotUsingGoogle
	}

	if *entity == "learner" && !user.IsLearner {
		return "", "", custom_error.ErrUserNotFound
	} else if *entity == "mentor" && !*user.IsMentor {
		return "", "", custom_error.ErrUserNotFound
	}

	var access_token string
	access_token, err = utils.GenerateJWT(user.ID, user.IsMentor)
	if err != nil {
		return "", "", err
	}

	var refresh_token string
	refresh_token, err = utils.GenerateRefreshToken(user.ID, user.IsMentor)
	if err != nil {
		return "", "", err
	}

	return access_token, refresh_token, nil
}
