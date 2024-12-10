package learner

import (
	"reflect"

	custom_error "github.com/philipnathan/pijar-backend/internal/learner/custom_error"
	dto "github.com/philipnathan/pijar-backend/internal/learner/dto"
	model "github.com/philipnathan/pijar-backend/internal/learner/model"
	repo "github.com/philipnathan/pijar-backend/internal/learner/repository"
	"gorm.io/gorm"
)

type LearnerBioServiceInterface interface {
	CreateLearnerBio(UserID uint, input *dto.CreateLearnerBioDto) error
	GetLearnerBio(UserID uint) (*model.LearnerBio, error)
	UpdateLearnerBio(UserID uint, input *dto.UpdateLearnerBioDto) error
}

type LearnerBioService struct {
	repo repo.LearnerBioRepositoryInterface
}

func NewLearnerBioService(repo repo.LearnerBioRepositoryInterface) LearnerBioServiceInterface {
	return &LearnerBioService{
		repo: repo,
	}
}

func (s *LearnerBioService) CreateLearnerBio(UserID uint, input *dto.CreateLearnerBioDto) error {
	_, err := s.repo.GetLearnerBio(UserID)
	if err != gorm.ErrRecordNotFound {
		return custom_error.ErrLearnerBioAlreadyExist
	}
	return s.repo.CreateLearnerBio(UserID, input)
}

func (s *LearnerBioService) GetLearnerBio(UserID uint) (*model.LearnerBio, error) {
	bio, err := s.repo.GetLearnerBio(UserID)
	if err != nil {
		return nil, err
	}

	return bio, nil
}

func (s *LearnerBioService) UpdateLearnerBio(UserID uint, input *dto.UpdateLearnerBioDto) error {
	learnerBio, _ := s.repo.GetLearnerBio(UserID)
	if learnerBio == nil {
		return custom_error.ErrLearnerBioNotFound
	}

	val := reflect.ValueOf(input).Elem()
	valType := val.Type()
	learnerBioVal := reflect.ValueOf(learnerBio).Elem()

	for i := 0; i < val.NumField(); i++ {
		valField := val.Field(i)
		valFieldName := valType.Field(i).Name
		oldField := learnerBioVal.FieldByName(valFieldName)

		if valField.Kind() == reflect.Ptr && !valField.IsNil() && oldField.IsValid() && oldField.CanSet() {
			oldField.Set(valField.Elem())
		}
	}

	return s.repo.SaveBio(learnerBio)
}
