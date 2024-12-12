package mentor

import (
	model "github.com/philipnathan/pijar-backend/internal/mentor/model"
	"gorm.io/gorm"
)

type MentorExpertiseRepositoryInterface interface {
	SaveMentorExpertise(mentor *model.MentorExpertises) error
}

type MentorExpertiseRepository struct {
	db *gorm.DB
}

func NewMentorExpertiseRepository(db *gorm.DB) MentorExpertiseRepositoryInterface {
	return &MentorExpertiseRepository{
		db: db,
	}
}

func (r *MentorExpertiseRepository) SaveMentorExpertise(mentor *model.MentorExpertises) error {
	err := r.db.Create(mentor).Error
	if err != nil {
		return err
	}

	return nil
}
