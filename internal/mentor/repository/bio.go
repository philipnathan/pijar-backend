package mentor

import (
	model "github.com/philipnathan/pijar-backend/internal/mentor/model"
	"gorm.io/gorm"
)

type MentorBioRepositoryInterface interface {
	GetMentorBio(userID *uint) (*model.MentorBiographies, error)
	SaveMentorBio(mentor *model.MentorBiographies) error
}

type MentorBioRepository struct {
	db *gorm.DB
}

func NewMentorBioRepository(db *gorm.DB) MentorBioRepositoryInterface {
	return &MentorBioRepository{
		db: db,
	}
}

func (r *MentorBioRepository) GetMentorBio(userID *uint) (*model.MentorBiographies, error) {
	var bio model.MentorBiographies
	err := r.db.Where("user_id = ?", userID).First(&bio).Error
	return &bio, err
}

func (r *MentorBioRepository) SaveMentorBio(mentor *model.MentorBiographies) error {
	err := r.db.Create(mentor).Error
	if err != nil {
		return err
	}
	return nil
}
