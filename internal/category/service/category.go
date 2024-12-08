package category

import (
	model "github.com/philipnathan/pijar-backend/internal/category/model"
	repository "github.com/philipnathan/pijar-backend/internal/category/repository"
)

type CategoryServiceInterface interface{
	GetAllCategoriesService() ([]model.Category, error)
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