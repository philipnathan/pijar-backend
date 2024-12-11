package mentor

import (
	service "github.com/philipnathan/pijar-backend/internal/mentor/service"
)

type MentorExpertiseHandler struct {
	service service.MentorExpertiseServiceInterface
}

func NewMentorExpertiseHandler(repo service.MentorExpertiseServiceInterface) *MentorExpertiseHandler {
	return &MentorExpertiseHandler{
		service: repo,
	}
}
