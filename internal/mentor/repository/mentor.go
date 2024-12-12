package mentor

import (
	model "github.com/philipnathan/pijar-backend/internal/user/model"
	"gorm.io/gorm"
)

type MentorRepositoryInterface interface {
	GetMentorDetails(MentorID *uint) (*model.User, error)
}

type MentorRepository struct {
	db *gorm.DB
}

func NewMentorRepository(db *gorm.DB) *MentorRepository {
	return &MentorRepository{
		db: db,
	}
}

func (r *MentorRepository) GetMentorDetails(MentorID *uint) (*model.User, error) {
	var user model.User
	err := r.db.Preload("MentorBio").
		Preload("MentorExperiences").
		Preload("MentorExpertises").
		Preload("MentorExpertises.Category").
		Where("id = ?", MentorID).
		First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}
