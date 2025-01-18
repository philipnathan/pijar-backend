package search

import (
	repository "github.com/philipnathan/pijar-backend/internal/search/repository"

	categoryModel "github.com/philipnathan/pijar-backend/internal/category/model"
	sessionModel "github.com/philipnathan/pijar-backend/internal/session/model"
	userModel "github.com/philipnathan/pijar-backend/internal/user/model"
)

type SearchServiceInterface interface {
	Search(keyword *string, page, pageSize *int) (*[]sessionModel.MentorSession, *[]userModel.User, *[]categoryModel.Category, int, error)
}

type SearchService struct {
	repo repository.SearchRepositoryInterface
}

func NewSearchService(repo repository.SearchRepositoryInterface) SearchServiceInterface {
	return &SearchService{
		repo: repo,
	}
}

func (s *SearchService) Search(keyword *string, page, pageSize *int) (*[]sessionModel.MentorSession, *[]userModel.User, *[]categoryModel.Category, int, error) {
	sessions, total, err := s.repo.SearchSessions(keyword, page, pageSize)
	if err != nil {
		return nil, nil, nil, 0, err
	}

	mentors, err := s.repo.SearchMentors(keyword)
	if err != nil {
		return nil, nil, nil, 0, err
	}

	categories, err := s.repo.SearchCategories(keyword)
	if err != nil {
		return nil, nil, nil, 0, err
	}

	return sessions, mentors, categories, total, nil
}
