package mentor_session_participant

import (
	custom_error "github.com/philipnathan/pijar-backend/internal/mentor_session_participant/custom_error"
	repo "github.com/philipnathan/pijar-backend/internal/mentor_session_participant/repository"
	session "github.com/philipnathan/pijar-backend/internal/session/service"
	userService "github.com/philipnathan/pijar-backend/internal/user/service"
)

type MentorSessionParticipantServiceInterface interface {
	CreateMentorSessionParticipant(userID, mentorSessionID *uint) error
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

func (s *MentorSessionParticipantService) CreateMentorSessionParticipant(userID, mentorSessionID *uint) error {
	// check if user exist
	user, err := s.userService.GetUserDetails(*userID)
	if err != nil {
		return err
	}
	if user == nil {
		return custom_error.ErrUserNotFound
	}
	if user.DeletedAt.Valid {
		return custom_error.ErrUserNotFound
	}

	// check if session exist
	session, err := s.sessionService.GetSessionByID(*mentorSessionID)
	if err != nil {
		return err
	}
	if session == nil {
		return custom_error.ErrSessionNotFound
	}

	// fmt.Printf("Session: %v\n", session.Schedule)

	// if session.Schedule.Before(time.Now().UTC()) {
	// 	return custom_error.ErrSessionAlreadyFinished
	// }

	// check if user already registered
	_, err = s.repo.GetMentorSessionParticipant(userID, mentorSessionID)
	if err == nil {
		return custom_error.ErrUserAlreadyRegistered
	}

	err = s.repo.CreateMentorSessionParticipant(userID, mentorSessionID)
	if err != nil {
		return err
	}

	return nil
}
