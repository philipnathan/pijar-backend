package category

import (
	dto "github.com/philipnathan/pijar-backend/internal/category/dto"
	model "github.com/philipnathan/pijar-backend/internal/category/model"
	repository "github.com/philipnathan/pijar-backend/internal/category/repository"
)

type CategoryServiceInterface interface {
	GetAllCategoriesService() ([]model.Category, error)
	GetFeaturedCategoriesService() ([]dto.FeaturedCategoryResponseDto, error)
}

type CategoryService struct {
	repo repository.CategoryRepositoryInterface
}

func NewCategoryService(repo repository.CategoryRepositoryInterface) CategoryServiceInterface {
	return &CategoryService{
		repo: repo,
	}
}

func (s *CategoryService) GetAllCategoriesService() ([]model.Category, error) {
	return s.repo.GetAllCategories()
}

func (s *CategoryService) GetFeaturedCategoriesService() ([]dto.FeaturedCategoryResponseDto, error) {
	return s.repo.GetFeaturedCategories()
}
