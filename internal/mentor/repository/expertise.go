package mentor

import "gorm.io/gorm"

type MentorExpertiseRepositoryInterface interface{}

type MentorExpertiseRepository struct {
	db *gorm.DB
}

func NewMentorExpertiseRepository(db *gorm.DB) MentorExpertiseRepositoryInterface {
	return &MentorExpertiseRepository{
		db: db,
	}
}
