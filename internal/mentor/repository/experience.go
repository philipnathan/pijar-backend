package mentor

import (
	model "github.com/philipnathan/pijar-backend/internal/mentor/model"
	"gorm.io/gorm"
)

type MentorExperiencesRepositoryInterface interface {
	SaveMentorExperience(mentor *model.MentorExperiences) error
}

type MentorExperiencesRepository struct {
	db *gorm.DB
}

func NewMentorExperienceRepository(db *gorm.DB) MentorExperiencesRepositoryInterface {
	return &MentorExperiencesRepository{
		db: db,
	}
}

func (r *MentorExperiencesRepository) SaveMentorExperience(mentor *model.MentorExperiences) error {
	err := r.db.Create(mentor).Error
	if err != nil {
		return err
	}
	return nil
}
