package learner

import (
	dto "github.com/philipnathan/pijar-backend/internal/learner/dto"
	model "github.com/philipnathan/pijar-backend/internal/learner/model"
	"gorm.io/gorm"
)

type LearnerBioRepositoryInterface interface {
	CreateLearnerBio(UserID uint, input *dto.CreateLearnerBioDto) error
	GetLearnerBio(UserID uint) (*model.LearnerBio, error)
}

type LearnerBioRepository struct {
	db *gorm.DB
}

func NewLearnerBioRepository(db *gorm.DB) LearnerBioRepositoryInterface {
	return &LearnerBioRepository{
		db: db,
	}
}

func (r *LearnerBioRepository) CreateLearnerBio(UserID uint, input *dto.CreateLearnerBioDto) error {
	return r.db.Create(&model.LearnerBio{UserID: UserID, Bio: input.Bio, Occupation: input.Occupation, Institution: input.Institution}).Error
}

func (r *LearnerBioRepository) GetLearnerBio(UserID uint) (*model.LearnerBio, error) {
	var bio model.LearnerBio
	err := r.db.Where("user_id = ?", UserID).First(&bio).Error
	return &bio, err
}
