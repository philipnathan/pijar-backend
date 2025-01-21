package user

import (
	"time"

	custom_error "github.com/philipnathan/pijar-backend/internal/user/custom_error"
	dto "github.com/philipnathan/pijar-backend/internal/user/dto"
	model "github.com/philipnathan/pijar-backend/internal/user/model"
	repository "github.com/philipnathan/pijar-backend/internal/user/repository"
	"github.com/philipnathan/pijar-backend/utils"
	"gorm.io/gorm"
)

type UserServiceInterface interface {
	RegisterUserService(email, password, fullname *string) (string, string, error)
	LoginUserService(email, password string) (string, string, error)
	GetUserDetails(userID uint) (*model.User, error)
	DeleteUserService(userID uint) error
	UpdateUserPasswordService(userID uint, oldPassword, newPassword string) error
	UpdateUserDetailsService(userID uint, input dto.UpdateUserDetailsDto) error
	GetUserProfile(userID uint) (*dto.UserProfileResponse, error)
}

type UserService struct {
	repo repository.UserRepositoryInterface
}

func NewUserService(repo repository.UserRepositoryInterface) UserServiceInterface {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) RegisterUserService(email, password, fullname *string) (string, string, error) {
	var user *model.User
	var err error

	user, err = s.repo.FindUserByEmail(*email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return "", "", err
	}

	if user != nil {
		if user.IsLearner {
			return "", "", custom_error.ErrAlreadyLearner
		}
		return "", "", custom_error.ErrChangeDetails
	}

	hashedPassword, err := utils.HashPassword(*password)
	if err != nil {
		return "", "", err
	}
	user, err = s.repo.CreateUser(email, &hashedPassword, fullname)
	if err != nil {
		return "", "", err
	}

	var access_token string
	if access_token, err = utils.GenerateJWT(user.ID, user.IsMentor); err != nil {
		return "", "", err
	}

	var refresh_token string
	if refresh_token, err = utils.GenerateRefreshToken(user.ID, user.IsMentor); err != nil {
		return "", "", err
	}

	return access_token, refresh_token, nil
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

	return access_token, refresh_token, nil
}

func (s *UserService) GetUserDetails(userID uint) (*model.User, error) {
	user, err := s.repo.FindByUserId(userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, custom_error.ErrUserNotFound
		}
		return nil, err
	}

	return user, nil
}

func (s *UserService) DeleteUserService(userID uint) error {
	return s.repo.DeleteUserById(userID)
}

func (s *UserService) UpdateUserPasswordService(userID uint, oldPassword, newPassword string) error {

	user, err := s.repo.FindByUserId(userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return custom_error.ErrUserNotFound
		}
		return err
	}

	if err := utils.ComparePassword(user.Password, oldPassword); err != nil {
		return custom_error.ErrWrongPassword
	}
	if oldPassword == newPassword {
		return custom_error.ErrSamePassword
	}

	user.Password, err = utils.HashPassword(newPassword)
	if err != nil {
		return err
	}

	if err := s.repo.UpdateUserPassword(user); err != nil {
		return err
	}

	return nil
}

func (s *UserService) UpdateUserDetailsService(userID uint, input dto.UpdateUserDetailsDto) error {
	var user, err = s.repo.FindByUserId(userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return custom_error.ErrUserNotFound
		}
		return err
	}

	if user == nil {
		return custom_error.ErrUserNotFound
	}

	if input.Fullname != "" {
		user.Fullname = input.Fullname
	}

	if input.BirthDate != "" {
		var customTime model.CustomTime
		customTime.Time, err = time.Parse("2006-01-02", input.BirthDate)
		if err != nil {
			return err
		}

		user.BirthDate = &customTime
	}

	if input.PhoneNumber != "" {
		user.PhoneNumber = &input.PhoneNumber
	}

	if input.ImageURL != "" {
		user.ImageURL = &input.ImageURL
	}

	if input.IsMentor != nil {
		if *input.IsMentor {
			user.IsMentor = input.IsMentor
		} else {
			return custom_error.ErrStatusCannotBeFalse
		}
	}

	if input.IsLearner != nil {
		if *input.IsLearner {
			user.IsLearner = *input.IsLearner
		} else {
			return custom_error.ErrStatusCannotBeFalse
		}
	}

	if err := s.repo.SaveUser(user); err != nil {
		return err
	}

	return nil
}

func (s *UserService) GetUserProfile(userID uint) (*dto.UserProfileResponse, error) {
	user, err := s.repo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	return &dto.UserProfileResponse{
		ID:       user.ID,
		Fullname: user.Fullname,
		Email:    user.Email,
		ImageURL: *user.ImageURL,
	}, nil
}
