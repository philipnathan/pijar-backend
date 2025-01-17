package follow

import (
	custom_error "github.com/philipnathan/pijar-backend/internal/follow/custom_error"
	repo "github.com/philipnathan/pijar-backend/internal/follow/repository"
	userService "github.com/philipnathan/pijar-backend/internal/user/service"
)

type FollowServiceInterface interface {
	FollowUnfollow(followerID, followingID *uint) error
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

func (s *FollowService) FollowUnfollow(followerID, followingID *uint) error {
	// check if both user are exist
	learner, err := s.userService.GetUserDetails(*followerID)
	if err != nil {
		return err
	}
	if !learner.IsLearner {
		return custom_error.ErrNotLearner
	}

	mentor, err := s.userService.GetUserDetails(*followingID)
	if err != nil {
		return err
	}
	if mentor.IsMentor == nil || !*mentor.IsMentor {
		return custom_error.ErrNotMentor
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
