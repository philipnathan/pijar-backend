package session

import (
	learnerRepo "github.com/philipnathan/pijar-backend/internal/learner/repository"
	model "github.com/philipnathan/pijar-backend/internal/session/model"
	repository "github.com/philipnathan/pijar-backend/internal/session/repository"
)

type SessionService interface {
	GetUpcomingSessions(page, pageSize int) ([]model.MentorSession, int, error)
	GetLearnerHistorySession(userID uint) (*[]model.MentorSessionParticipant, error)
	GetSessionByLearnerInterests(userID uint, page, pageSize int) (*[]model.MentorSession, int, error)
	GetUpcommingSessionsByCategory(categoryID []uint, page, pageSize int) (*[]model.MentorSession, int, error)
	GetAllSessionsWithFilter(categoryID, mentorID uint, page, pageSize int) (*[]model.MentorSession, int, error)
	GetSessionByID(sessionID uint) (*model.MentorSession, error)
	GetDetailSessionByID(sessionID uint) (*model.MentorSession, error)
}

type sessionService struct {
	repo                 repository.SessionRepository
	learnerInterestsRepo learnerRepo.LearnerRepositoryInterface
}

func NewSessionService(repo repository.SessionRepository, learnerInterestsRepo learnerRepo.LearnerRepositoryInterface) SessionService {
	return &sessionService{repo: repo, learnerInterestsRepo: learnerInterestsRepo}
}

func (s *sessionService) GetUpcomingSessions(page, pageSize int) ([]model.MentorSession, int, error) {
	return s.repo.GetUpcomingSessions(page, pageSize)
}

func (s *sessionService) GetLearnerHistorySession(userID uint) (*[]model.MentorSessionParticipant, error) {
	// Fetch sessions from the repository
	session, err := s.repo.GetLearnerHistorySession(&userID)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (s *sessionService) GetSessionByLearnerInterests(userID uint, page, pageSize int) (*[]model.MentorSession, int, error) {
	userInterests, err := s.learnerInterestsRepo.GetLearnerInterest(userID)
	if err != nil {
		return nil, 0, err
	}

	// convert user interest to slice of uint
	var interests []uint
	for _, interest := range userInterests {
		interests = append(interests, interest.CategoryID)
	}

	// if user doesn't have any interest
	if interests == nil {
		sessions, total, err := s.repo.GetUpcomingSessions(page, pageSize)
		if err != nil {
			return nil, 0, err
		}
		return &sessions, total, nil
	}

	// if user has interest, return sessions with that interest
	sessions, total, err := s.repo.GetUpcommingSessionsByCategory(interests, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	return sessions, total, nil
}

func (s *sessionService) GetUpcommingSessionsByCategory(categoryID []uint, page, pageSize int) (*[]model.MentorSession, int, error) {
	sessions, total, err := s.repo.GetUpcommingSessionsByCategory(categoryID, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	return sessions, total, nil
}

func (s *sessionService) GetAllSessionsWithFilter(categoryID, mentorID uint, page, pageSize int) (*[]model.MentorSession, int, error) {
	sessions, total, err := s.repo.GetAllSessionsWithFilter(categoryID, mentorID, page, pageSize)

	if err != nil {
		return nil, 0, err
	}

	return sessions, total, nil
}

func (s *sessionService) GetSessionByID(sessionID uint) (*model.MentorSession, error) {
	session, err := s.repo.GetSessionByID(sessionID)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (s *sessionService) GetDetailSessionByID(sessionID uint) (*model.MentorSession, error) {
	session, err := s.repo.GetSessionDetailByID(sessionID)
	if err != nil {
		return nil, err
	}
	return session, nil
}
