package user

import (
	model "github.com/philipnathan/pijar-backend/internal/user/model"
	"gorm.io/gorm"
)

type MentorUserRepositoryInterface interface {
	CreateNewMentor(email, password, fullname *string) (*model.User, error)
	SetIsMentorToTrue(email *string) (*model.User, error)
	IsUserExist(email *string) (bool, error)
	FindUserByEmail(email string) (*model.User, error)
}

type MentorUserRepository struct {
	db *gorm.DB
}

func NewMentorUserRepository(db *gorm.DB) *MentorUserRepository {
	return &MentorUserRepository{
		db: db,
	}
}

func (r *MentorUserRepository) CreateNewMentor(email, password, fullname *string) (*model.User, error) {
	var mentor model.User
	isActive := true
	err := r.db.Create(&model.User{
		Email:    *email,
		Password: *password,
		Fullname: *fullname,
		IsMentor: &isActive,
	}).Error

	if err != nil {
		return nil, err
	}

	return &mentor, nil
}

func (r *MentorUserRepository) SetIsMentorToTrue(email *string) (*model.User, error) {
	var mentor model.User
	err := r.db.
		Model(&mentor).
		Where("email = ?", *email).
		Update("is_mentor", true).
		Error
	if err != nil {
		return nil, err
	}

	err = r.db.Where("email = ?", *email).First(&mentor).Error
	if err != nil {
		return nil, err
	}

	return &mentor, nil
}

func (r *MentorUserRepository) IsUserExist(email *string) (bool, error) {
	var user model.User
	err := r.db.Where("email = ?", *email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (r *MentorUserRepository) FindUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
