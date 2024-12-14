package service

import (
	"time"
	model "github.com/philipnathan/pijar-backend/internal/session/model"
	repository "github.com/philipnathan/pijar-backend/internal/session/repository"
	dto "github.com/philipnathan/pijar-backend/internal/session/dto" // Import the dto package
)

type SessionService struct {
	SessionRepo repository.SessionRepository
}

func MapMentorSessionToResponse(session model.MentorSession, isRegistered bool) model.MentorSessionResponse {
	day := session.Schedule.Weekday().String()
	timeFormatted := session.Schedule.Format("03:04 PM") 

	return model.MentorSessionResponse{
		Day:                day,
		Time:               timeFormatted,
		MentorSessionTitle: session.Title,
		ShortDescription:   session.ShortDescription,
		Schedule:           session.Schedule.Format("2006-01-02"), 
		Registered:         isRegistered,
	}
}

func (s *SessionService) GetUpcomingSessions(userID uint) (*dto.GetUpcomingSessionResponse, error) {
	// Fetch upcoming sessions from the repository
	sessions, err := s.SessionRepo.FetchUpcomingSessions(userID)
	if err != nil {
		return nil, err
	}

	// Map the repository result to the DTO
	var sessionDetails []dto.SessionDetail
	for _, session := range sessions {
		day := session.Schedule.Weekday().String()
		time := session.Schedule.Format("03:04 PM") 
		sessionDetails = append(sessionDetails, dto.SessionDetail{
			Day:                day,
			Time:               time,
			MentorSessionTitle: session.Title,
			ShortDescription:   session.ShortDescription,
			Schedule:           session.Schedule.Format("2006-01-02"),
			Registered:         session.IsRegistered,
		})
	}

	return &dto.GetUpcomingSessionResponse{
		Sessions: sessionDetails,
	}, nil
}