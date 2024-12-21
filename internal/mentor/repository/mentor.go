package mentor

import (
	model "github.com/philipnathan/pijar-backend/internal/user/model"
	"gorm.io/gorm"
)

type MentorRepositoryInterface interface {
	GetMentorDetails(MentorID *uint) (*model.User, error)
	GetMentorByMentorExpertisesCategory(interests []uint, page, pageSize int) ([]model.User, int, error)
	GetAllMentors(page, pageSize int) (*[]model.User, int, error)
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
		Preload("MentorExperiences", func(db *gorm.DB) *gorm.DB { return db.Order("COALESCE(end_date, NOW()) DESC") }).
		Preload("MentorExpertises").
		Preload("MentorExpertises.Category").
		Where("id = ?", MentorID).
		First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *MentorRepository) GetMentorByMentorExpertisesCategory(interests []uint, page, pageSize int) ([]model.User, int, error) {
	var users []model.User
	var total int64

	countQuery := r.db.Model(&model.User{}).
		Joins("JOIN mentor_expertises ON mentor_expertises.user_id = users.id").
		Where("mentor_expertises.category_id IN ?", interests).
		Distinct("users.id")

	if err := countQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.Preload("MentorExpertises").
		Preload("MentorExperiences", func(db *gorm.DB) *gorm.DB { return db.Order("COALESCE(end_date, NOW()) DESC") }).
		Preload("MentorExpertises.Category").
		Distinct("users.*").
		Joins("JOIN mentor_expertises ON mentor_expertises.user_id = users.id").
		Where("mentor_expertises.category_id IN ?", interests).
		Order("created_at DESC").
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&users).Error

	if err != nil {
		return nil, 0, err
	}

	return users, int(total), nil
}

func (r *MentorRepository) GetAllMentors(page, pageSize int) (*[]model.User, int, error) {
	var users []model.User
	var total int64

	countQuery := r.db.Model(&model.User{}).
		Joins("JOIN mentor_expertises ON mentor_expertises.user_id = users.id").
		Distinct("users.id")

	if err := countQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.Preload("MentorExpertises").
		Preload("MentorExperiences", func(db *gorm.DB) *gorm.DB { return db.Order("COALESCE(end_date, NOW()) DESC") }).
		Preload("MentorExpertises.Category").
		Distinct("users.*").
		Joins("JOIN mentor_expertises ON mentor_expertises.user_id = users.id").
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&users).Error

	if err != nil {
		return nil, 0, err
	}

	return &users, int(total), nil
}
