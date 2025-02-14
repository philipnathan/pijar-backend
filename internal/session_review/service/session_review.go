package session_review

import (
	"context"

	participantService "github.com/philipnathan/pijar-backend/internal/mentor_session_participant/service"
	sessionService "github.com/philipnathan/pijar-backend/internal/session/service"
	custom_error "github.com/philipnathan/pijar-backend/internal/session_review/custom_error"
	model "github.com/philipnathan/pijar-backend/internal/session_review/model"
	repo "github.com/philipnathan/pijar-backend/internal/session_review/repository"
	userService "github.com/philipnathan/pijar-backend/internal/user/service"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

type SessionReviewServiceInterface interface {
	CreateSessionReview(ctx context.Context, userID, sessionID, rating *uint, review *string) error
	GetSessionReviews(sessionID *uint, page, pageSize *int) (*[]model.SessionReview, int, error)
}

type SessionReviewService struct {
	repo               repo.SessionReviewRepositoryInterface
	userService        userService.UserServiceInterface
	sessionService     sessionService.SessionService
	participantService participantService.MentorSessionParticipantServiceInterface
}

func NewSessionReviewService(
	repo repo.SessionReviewRepositoryInterface,
	userService userService.UserServiceInterface,
	sessionService sessionService.SessionService,
	participantService participantService.MentorSessionParticipantServiceInterface) SessionReviewServiceInterface {
	return &SessionReviewService{
		repo:               repo,
		userService:        userService,
		sessionService:     sessionService,
		participantService: participantService,
	}
}

func (s *SessionReviewService) CreateSessionReview(ctx context.Context, userID, sessionID, rating *uint, review *string) error {
	g, _ := errgroup.WithContext(ctx)
	// check if user exist
	g.Go(func() error {
		_, err := s.userService.GetUserDetails(*userID)
		if err != nil {
			return err
		}
		return nil
	})

	// check if session exist
	g.Go(func() error {
		_, err := s.sessionService.GetSessionByID(*sessionID)
		if err != nil {
			return err
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		return err
	}

	// check if learner is part of session
	_, err := s.participantService.GetLearnerEnrollment(userID, sessionID)
	if err != nil && err == gorm.ErrRecordNotFound {
		return custom_error.ErrLearnerNotEnrolled
	}
	if err != nil {
		return err
	}

	// check if user already reviewed
	rev, _ := s.repo.GetUserReview(userID, sessionID)
	if rev != nil {
		return custom_error.ErrUserAlreadyReviewed
	}

	//create session review
	err = s.repo.CreateSessionReview(userID, sessionID, rating, review)
	if err != nil {
		return err
	}

	return nil
}

func (s *SessionReviewService) GetSessionReviews(sessionID *uint, page, pageSize *int) (*[]model.SessionReview, int, error) {
	// check if session exist
	session, err := s.sessionService.GetSessionByID(*sessionID)
	if session == nil {
		return nil, 0, custom_error.ErrSessionNotFound
	}
	if err != nil {
		return nil, 0, err
	}

	return s.repo.GetSessionReviews(sessionID, page, pageSize)
}
