package user

import (
	custom_error "github.com/philipnathan/pijar-backend/internal/user/custom_error"
	model "github.com/philipnathan/pijar-backend/internal/user/model"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	CreateUser(email, password, fullname *string) (*model.User, error)
	FindUserByEmail(email string) (*model.User, error)
	FindByPhoneNumber(phoneNumber string) (*model.User, error)
	FindByUserId(id uint) (*model.User, error)
	DeleteUserById(id uint) error
	UpdateUserPassword(user *model.User) error
	SaveUser(user *model.User) error
	GetUserByID(userID uint) (*model.User, error)
	SetIsLearnerToTrue(email *string) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) CreateUser(email, password, fullname *string) (*model.User, error) {

	user := &model.User{
		Email:     *email,
		Password:  *password,
		Fullname:  *fullname,
		IsLearner: true,
	}

	err := r.db.Create(user).Error
	if err != nil {
		return &model.User{}, err
	}

	return user, nil
}

func (r *userRepository) FindUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByPhoneNumber(phoneNumber string) (*model.User, error) {
	var user model.User
	err := r.db.Where("phone_number = ?", phoneNumber).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByUserId(id uint) (*model.User, error) {
	var user model.User
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

func (r *userRepository) DeleteUserById(id uint) error {
	return r.db.Where("id = ?", id).Delete(&model.User{}).Error
}

func (r *userRepository) UpdateUserPassword(user *model.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) SaveUser(user *model.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) GetUserByID(userID uint) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, userID).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) SetIsLearnerToTrue(email *string) (*model.User, error) {
	var user model.User
	err := r.db.
		Model(&user).
		Where("email = ?", *email).
		Update("is_learner", true).
		Error
	if err != nil {
		return nil, err
	}

	err = r.db.Where("email = ?", *email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
