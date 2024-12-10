package learner

type CreateLearnerBioDto struct {
	Bio         string `json:"bio" example:"My bio" binding:"omitempty"`
	Occupation  string `json:"occupation" example:"Software Engineer" binding:"omitempty"`
	Institution string `json:"institution" example:"Google" binding:"omitempty"`
}

type CreateLearnerBioResponseDto struct {
	Message string `json:"message" example:"bio added successfully"`
}
