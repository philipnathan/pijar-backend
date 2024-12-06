package repositories

import (
	"github.com/philipnathan/pijar-backend/custom_error"
	"github.com/philipnathan/pijar-backend/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	FindUserByEmail(email string) (*models.User, error)
	FindByPhoneNumber(phoneNumber string) (*models.User, error)
	FindByUserId(id uint) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) CreateUser(user *models.User) (error) {
	return r.db.Create(user).Error
}

func (r *userRepository) FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByPhoneNumber(phoneNumber string) (*models.User, error) {
	var user models.User
	err := r.db.Where("phone_number = ?", phoneNumber).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByUserId(id uint) (*models.User, error) {
	var user models.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, custom_error.ErrUserNotFound
		} else {
			return nil, err
		}
	}
	return &user, nil
}