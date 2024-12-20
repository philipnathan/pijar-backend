package session

import (
	model "github.com/philipnathan/pijar-backend/internal/session/model"
	repository "github.com/philipnathan/pijar-backend/internal/session/repository"
)

type SessionService interface {
	GetSessions(userID uint) ([]model.MentorSession, error)
	GetUpcomingSessions() ([]model.MentorSession, error)
	GetLearnerHistorySession(userID uint) (*[]model.MentorSessionParticipant, error)
}

type sessionService struct {
	repo repository.SessionRepository
}

func NewSessionService(repo repository.SessionRepository) SessionService {
	return &sessionService{repo: repo}
}

func (s *sessionService) GetSessions(userID uint) ([]model.MentorSession, error) {
	// Fetch sessions from the repository
	sessions, err := s.repo.GetSessions(userID)
	if err != nil {
		return nil, err
	}

	return sessions, nil
}

func (s *sessionService) GetUpcomingSessions() ([]model.MentorSession, error) {
	return s.repo.GetUpcomingSessions()
}

func (s *sessionService) GetLearnerHistorySession(userID uint) (*[]model.MentorSessionParticipant, error) {
	// Fetch sessions from the repository
	session, err := s.repo.GetLearnerHistorySession(&userID)
	if err != nil {
		return nil, err
	}

	return session, nil
}
