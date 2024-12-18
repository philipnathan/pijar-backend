package search

import (
    dto "github.com/philipnathan/pijar-backend/internal/search/dto"
    repository "github.com/philipnathan/pijar-backend/internal/search/repository"
)

type SearchServiceInterface interface {
    Search(keyword string) (*dto.SearchResponse, error)
}

type SearchService struct {
    repo repository.SearchRepositoryInterface
}

func NewSearchService(repo repository.SearchRepositoryInterface) SearchServiceInterface {
    return &SearchService{
        repo: repo,
    }
}

func (s *SearchService) Search(keyword string) (*dto.SearchResponse, error) {
    sessions, err := s.repo.SearchSessions(keyword)
    if err != nil {
        return nil, err
    }

    mentors, err := s.repo.SearchMentors(keyword)
    if err != nil {
        return nil, err
    }

    categories, err := s.repo.SearchCategories(keyword)
    if err != nil {
        return nil, err
    }

    return &dto.SearchResponse{
        Sessions:  sessions,
        Mentors:   mentors,
        Topics: categories,
    }, nil
}