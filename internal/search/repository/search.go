package repository

import (
	categoryModel "github.com/philipnathan/pijar-backend/internal/category/model"
	sessionModel "github.com/philipnathan/pijar-backend/internal/session/model"
	userModel "github.com/philipnathan/pijar-backend/internal/user/model"
	"gorm.io/gorm"
)

type SearchRepositoryInterface interface {
	SearchSessions(keyword *string, page, pageSize *int) (*[]sessionModel.MentorSession, int, error)
	SearchMentors(keyword *string) (*[]userModel.User, error)
	SearchCategories(keyword *string) (*[]categoryModel.Category, error)
}

type SearchRepository struct {
	db *gorm.DB
}

func NewSearchRepository(db *gorm.DB) SearchRepositoryInterface {
	return &SearchRepository{
		db: db,
	}
}

func (r *SearchRepository) SearchSessions(keyword *string, page, pageSize *int) (*[]sessionModel.MentorSession, int, error) {
	var sessions []sessionModel.MentorSession
	var total int64

	countQuery := r.db.Model(&sessionModel.MentorSession{}).Where("LOWER(title) LIKE ? OR LOWER(short_description) LIKE ?", "%"+*keyword+"%", "%"+*keyword+"%")

	if err := countQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.
		Where("LOWER(title) LIKE ? OR LOWER(short_description) LIKE ?", "%"+*keyword+"%", "%"+*keyword+"%").
		Preload("User").
		Preload("Category").
		Preload("SessionReviews").
		Offset((*page - 1) * *pageSize).
		Limit(*pageSize).
		Find(&sessions).Error

	if err != nil {
		return nil, 0, err
	}

	return &sessions, int(total), nil
}

func (r *SearchRepository) SearchMentors(keyword *string) (*[]userModel.User, error) {
	var mentors []userModel.User
	err := r.db.
		Where("LOWER(fullname) LIKE ? AND is_mentor = ?", "%"+*keyword+"%", true).
		Preload("MentorExperiences").
		Find(&mentors).Error
	if err != nil {
		return nil, err
	}

	return &mentors, nil
}

func (r *SearchRepository) SearchCategories(keyword *string) (*[]categoryModel.Category, error) {
	var categories []categoryModel.Category
	err := r.db.Where("LOWER(category_name) LIKE ? ", "%"+*keyword+"%").Find(&categories).Error
	if err != nil {
		return nil, err
	}

	return &categories, nil
}
