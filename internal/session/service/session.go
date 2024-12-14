package service

import (
    model "github.com/philipnathan/pijar-backend/internal/session/model"
    repository "github.com/philipnathan/pijar-backend/internal/session/repository"
)

type SessionService interface {
    GetSessions(userID uint) ([]model.MentorSession, error)
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