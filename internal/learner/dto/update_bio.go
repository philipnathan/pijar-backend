package learner

type UpdateLearnerBioDto struct {
	Bio         *string `json:"bio" example:"My bio" binding:"omitempty"`
	Occupation  *string `json:"occupation" example:"Software Engineer" binding:"omitempty"`
	Institution *string `json:"institution" example:"Google" binding:"omitempty"`
}

type UpdateLearnerBioResponseDto struct {
	Message string `json:"message" example:"bio updated successfully"`
}
