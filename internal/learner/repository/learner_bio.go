package learner

import "gorm.io/gorm"

type LearnerBioRepositoryInterface interface{}

type LearnerBioRepository struct {
	db *gorm.DB
}

func NewLearnerBioRepository(db *gorm.DB) LearnerBioRepositoryInterface {
	return &LearnerBioRepository{
		db: db,
	}
}
