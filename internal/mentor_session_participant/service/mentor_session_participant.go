package mentor_session_participant

import (
	"context"

	custom_error "github.com/philipnathan/pijar-backend/internal/mentor_session_participant/custom_error"
	model "github.com/philipnathan/pijar-backend/internal/mentor_session_participant/model"
	repo "github.com/philipnathan/pijar-backend/internal/mentor_session_participant/repository"
	session "github.com/philipnathan/pijar-backend/internal/session/service"
	userService "github.com/philipnathan/pijar-backend/internal/user/service"
	"golang.org/x/sync/errgroup"
)

type MentorSessionParticipantServiceInterface interface {
	CreateMentorSessionParticipant(ctx context.Context, userID, mentorSessionID *uint) error
	GetLearnerEnrollments(userID *uint, page, pageSize *int) (*[]model.MentorSessionParticipant, int, error)
	GetLearnerEnrollment(userID, mentorSessionID *uint) (*model.MentorSessionParticipant, error)
}

type MentorSessionParticipantService struct {
	repo           repo.MentorSessionParticipantRepositoryInterface
	userService    userService.UserServiceInterface
	sessionService session.SessionService
}

func NewMentorSessionParticipantService(
	repo repo.MentorSessionParticipantRepositoryInterface, userService userService.UserServiceInterface,
	sessionService session.SessionService) MentorSessionParticipantServiceInterface {
	return &MentorSessionParticipantService{
		repo:           repo,
		userService:    userService,
		sessionService: sessionService,
	}
}

func (s *MentorSessionParticipantService) CreateMentorSessionParticipant(ctx context.Context, userID, mentorSessionID *uint) error {

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		// check if user exist
		_, err := s.userService.GetUserDetails(*userID)
		if err != nil {
			return err
		}

		return nil
	})

	g.Go(func() error {
		// check if session exist
		_, err := s.sessionService.GetSessionByID(*mentorSessionID)
		if err != nil {
			return err
		}

		return nil
	})

	g.Go(func() error {
		// check if user already registered
		_, err := s.repo.GetMentorSessionParticipant(userID, mentorSessionID)
		if err == nil {
			return custom_error.ErrUserAlreadyRegistered
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		return err
	}

	err := s.repo.CreateMentorSessionParticipant(userID, mentorSessionID)
	if err != nil {
		return err
	}

	return nil
}

func (s *MentorSessionParticipantService) GetLearnerEnrollments(userID *uint, page, pageSize *int) (*[]model.MentorSessionParticipant, int, error) {
	// check if user exist
	user, err := s.userService.GetUserDetails(*userID)
	if err != nil {
		return nil, 0, err
	}
	if user == nil {
		return nil, 0, custom_error.ErrUserNotFound
	}
	if user.DeletedAt.Valid {
		return nil, 0, custom_error.ErrUserNotFound
	}

	data, total, err := s.repo.GetLearnerEnrollments(userID, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	return data, total, nil
}

func (s *MentorSessionParticipantService) GetLearnerEnrollment(userID, mentorSessionID *uint) (*model.MentorSessionParticipant, error) {
	// check if user exist
	user, err := s.userService.GetUserDetails(*userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, custom_error.ErrUserNotFound
	}
	if user.DeletedAt.Valid {
		return nil, custom_error.ErrUserNotFound
	}

	data, err := s.repo.GetLearnerEnrollment(userID, mentorSessionID)
	if err != nil {
		return nil, err
	}

	return data, nil
}
