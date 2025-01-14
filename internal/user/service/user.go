package user

import (
	"fmt"
	"reflect"

	custom_error "github.com/philipnathan/pijar-backend/internal/user/custom_error"
	dto "github.com/philipnathan/pijar-backend/internal/user/dto"
	model "github.com/philipnathan/pijar-backend/internal/user/model"
	repository "github.com/philipnathan/pijar-backend/internal/user/repository"
	"github.com/philipnathan/pijar-backend/utils"
	"gorm.io/gorm"
)

type UserServiceInterface interface {
	RegisterUserService(user *dto.RegisterUserDto) (string, string, error)
	LoginUserService(email, password string) (string, string, error)
	GetUserDetails(userID uint) (*model.User, error)
	DeleteUserService(userID uint) error
	UpdateUserPasswordService(userID uint, oldPassword, newPassword string) error
	UpdateUserDetailsService(userID uint, input interface{}) error
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

func (s *UserService) RegisterUserService(user *dto.RegisterUserDto) (string, string, error) {
	var err error

	if exist, err := s.isUserExist(user); err != nil {
		return "", "", err
	} else if exist {
		return "", "", err
	}

	user.Password, err = utils.HashPassword(user.Password)
	if err != nil {
		return "", "", err
	}

	var createdUser model.User
	createdUser, err = s.repo.CreateUser(user)
	if err != nil {
		return "", "", err
	}

	var access_token string
	if access_token, err = utils.GenerateJWT(createdUser.ID, createdUser.IsMentor); err != nil {
		return "", "", err
	}

	var refresh_token string
	if refresh_token, err = utils.GenerateRefreshToken(createdUser.ID, createdUser.IsMentor); err != nil {
		return "", "", err
	}

	return access_token, refresh_token, nil
}

func (s *UserService) isUserExist(user *dto.RegisterUserDto) (bool, error) {
	if exist, err := s.repo.FindUserByEmail(user.Email); err == nil || exist != nil {
		return true, custom_error.ErrEmailExist
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

func (s *UserService) UpdateUserDetailsService(userID uint, input interface{}) error {
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

	// Pastikan input adalah struct, bukan pointer
	v := reflect.ValueOf(input)
	if v.Kind() != reflect.Struct {
		return fmt.Errorf("input must be a struct")
	}

	// Loop untuk memeriksa field dalam input
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		// Skip jika field kosong (zero value)
		if field.IsZero() {
			continue
		}

		// Ambil field pada user berdasarkan nama
		userField := reflect.ValueOf(user).Elem().FieldByName(v.Type().Field(i).Name)

		// Jika field valid dan dapat di-set
		if userField.IsValid() && userField.CanSet() {

			// Jika field "BirthDate" maka lakukan unmarshall dahulu
			if v.Type().Field(i).Name == "BirthDate" {
				str := field.String()
				if str != "" {
					var customTime model.CustomTime
					err := customTime.UnmarshalJSON([]byte(str))
					if err != nil {
						return err
					}

					userField.Set(reflect.ValueOf(&customTime))
				}
				continue
			}

			// Set nilai field pada user
			if userField.Type() == field.Type() {
				userField.Set(field)
			} else {
				fmt.Printf("Type mismatch for field: %s\n", v.Type().Field(i).Name)
			}
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
