package follow

import (
	"context"

	repo "github.com/philipnathan/pijar-backend/internal/follow/repository"
	userService "github.com/philipnathan/pijar-backend/internal/user/service"
	"golang.org/x/sync/errgroup"
)

type FollowServiceInterface interface {
	FollowUnfollow(ctx context.Context, followerID, followingID *uint) error
	IsFollowing(followerID, followingID *uint) (bool, error)
}

type FollowService struct {
	repo        repo.FollowRepositoryInterface
	userService userService.UserServiceInterface
}

func NewFollowService(repo repo.FollowRepositoryInterface, userService userService.UserServiceInterface) FollowServiceInterface {
	return &FollowService{
		repo:        repo,
		userService: userService,
	}
}

func (s *FollowService) FollowUnfollow(ctx context.Context, followerID, followingID *uint) error {
	g, _ := errgroup.WithContext(ctx)

	g.Go(func() error {
		_, err := s.userService.GetUserDetails(*followerID)
		if err != nil {
			return err
		}

		return nil
	})

	g.Go(func() error {
		_, err := s.userService.GetUserDetails(*followingID)
		if err != nil {
			return err
		}

		return nil
	})

	if err := g.Wait(); err != nil {
		return err
	}

	status, err := s.repo.IsFollowing(followerID, followingID)
	if err != nil {
		return err
	}

	if status {
		return s.repo.Unfollow(followerID, followingID)
	} else {
		return s.repo.Follow(followerID, followingID)
	}
}

func (s *FollowService) IsFollowing(followerID, followingID *uint) (bool, error) {
	return s.repo.IsFollowing(followerID, followingID)
}
