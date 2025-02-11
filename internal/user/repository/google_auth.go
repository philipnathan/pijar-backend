package user

import (
	model "github.com/philipnathan/pijar-backend/internal/user/model"
	"gorm.io/gorm"
)

type GoogleAuthRepoInterface interface {
	FindUserByEmail(email *string) (*model.User, error)
	CreateLearner(email, password *string, fullname *string) (*model.User, error)
	CreateMentor(email, password *string, fullname *string) (*model.User, error)
	SetIsLearnerToTrue(email *string) (*model.User, error)
	SetIsMentorToTrue(email *string) (*model.User, error)
}

type GoogleAuthRepo struct {
	db *gorm.DB
}

func NewGoogleAuthRepo(db *gorm.DB) GoogleAuthRepoInterface {
	return &GoogleAuthRepo{
		db: db,
	}
}

func (r *GoogleAuthRepo) FindUserByEmail(email *string) (*model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *GoogleAuthRepo) CreateLearner(email, password *string, fullname *string) (*model.User, error) {
	var user model.User
	err := r.db.Create(&model.User{
		Email:        *email,
		Password:     *password,
		Fullname:     *fullname,
		IsLearner:    true,
		AuthProvider: "google",
	}).Error

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *GoogleAuthRepo) CreateMentor(email, password *string, fullname *string) (*model.User, error) {
	var user model.User
	status := true
	err := r.db.Create(&model.User{
		Email:        *email,
		Password:     *password,
		Fullname:     *fullname,
		IsMentor:     &status,
		AuthProvider: "google",
	}).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *GoogleAuthRepo) SetIsLearnerToTrue(email *string) (*model.User, error) {
	var user model.User
	err := r.db.Model(&user).Where("email = ?", *email).Update("is_learner", true).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *GoogleAuthRepo) SetIsMentorToTrue(email *string) (*model.User, error) {
	var user model.User
	err := r.db.Model(&user).Where("email = ?", *email).Update("is_mentor", true).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
