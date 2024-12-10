package learner

import (
	service "github.com/philipnathan/pijar-backend/internal/learner/service"
)

type LearnerBioHandler struct {
	service service.LearnerBioServiceInterface
}

func NewLearnerBioHandler(service service.LearnerBioServiceInterface) *LearnerBioHandler {
	return &LearnerBioHandler{
		service: service,
	}
}
