package category

import (
	model "github.com/philipnathan/pijar-backend/internal/category/model"
	dto "github.com/philipnathan/pijar-backend/internal/category/dto"
	"gorm.io/gorm"
)

type CategoryRepositoryInterface interface {
	GetAllCategories() ([]model.Category, error)
	SaveCategory(category *model.Category) error
	GetFeaturedCategories() ([]dto.FeaturedCategoryResponseDto, error)
}

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepositoryInterface {
	return &CategoryRepository{
		db: db,
	}
}

func (r *CategoryRepository) GetAllCategories() ([]model.Category, error) {
	var categories []model.Category
	err := r.db.Preload("SubCategories").Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *CategoryRepository) SaveCategory(category *model.Category) error {
	err := r.db.Create(category).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *CategoryRepository) GetFeaturedCategories() ([]dto.FeaturedCategoryResponseDto, error) {
    var categories []dto.FeaturedCategoryResponseDto
    err := r.db.Model(&model.Category{}).Select("category_name, image_url").Scan(&categories).Error
    return categories, err
}