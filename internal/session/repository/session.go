package service

import (
	model "github.com/philipnathan/pijar-backend/internal/session/model"
	repository "github.com/philipnathan/pijar-backend/internal/session/repository"
	"gorm.io/gorm"
)

type SessionService interface {
	GetSessions(userID uint) ([]model.MentorSessionResponse, error)
}

type sessionService struct {
	repo repository.SessionRepository
}

func NewSessionService(repo repository.SessionRepository) SessionService {
	return &sessionService{repo: repo}
}

func (s *sessionService) GetSessions(userID uint) ([]model.MentorSessionResponse, error) {
	// Fetch sessions from the repository
	sessions, err := s.repo.GetSessions(userID)
	if err != nil {
		return nil, err
	}

	// Map to API response struct
	var responses []model.MentorSessionResponse
	for _, session := range sessions {
		isRegistered := session.Registered
		response := MapMentorSessionToResponse(session, isRegistered)
		responses = append(responses, response)
	}

	return responses, nil
}
