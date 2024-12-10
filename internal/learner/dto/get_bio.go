package learner

import (
	model "github.com/philipnathan/pijar-backend/internal/learner/model"
)

type GetLearnerBioResponseDto struct {
	Message string            `json:"message" example:"bio fetched successfully"`
	Bio     *model.LearnerBio `json:"bio"`
}
