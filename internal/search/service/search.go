package search

import (
	"context"

	repository "github.com/philipnathan/pijar-backend/internal/search/repository"
	"golang.org/x/sync/errgroup"

	categoryModel "github.com/philipnathan/pijar-backend/internal/category/model"
	sessionModel "github.com/philipnathan/pijar-backend/internal/session/model"
	userModel "github.com/philipnathan/pijar-backend/internal/user/model"
)

type SearchServiceInterface interface {
	Search(ctx context.Context, keyword *string, page, pageSize *int) (*[]sessionModel.MentorSession, *[]userModel.User, *[]categoryModel.Category, int, error)
}

type SearchService struct {
	repo repository.SearchRepositoryInterface
}

func NewSearchService(repo repository.SearchRepositoryInterface) SearchServiceInterface {
	return &SearchService{
		repo: repo,
	}
}

func (s *SearchService) Search(ctx context.Context, keyword *string, page, pageSize *int) (*[]sessionModel.MentorSession, *[]userModel.User, *[]categoryModel.Category, int, error) {
	var se []sessionModel.MentorSession
	var t int
	var m []userModel.User
	var c []categoryModel.Category
	g, _ := errgroup.WithContext(ctx)

	g.Go(func() error {
		sessions, total, err := s.repo.SearchSessions(keyword, page, pageSize)
		if err != nil {
			return err
		}

		se = *sessions
		t = total

		return nil
	})

	g.Go(func() error {
		mentors, err := s.repo.SearchMentors(keyword)
		if err != nil {
			return err
		}

		m = *mentors

		return nil
	})

	g.Go(func() error {
		categories, err := s.repo.SearchCategories(keyword)
		if err != nil {
			return err
		}

		c = *categories

		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, nil, nil, 0, err
	}

	return &se, &m, &c, t, nil
}
