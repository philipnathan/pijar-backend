package learner

type DeleteLearnerInterestsDto struct {
	CategoryID []uint `json:"category_id" binding:"required" example:"1,2,3"`
}

type DeleteLearnerInterestsResponseDto struct {
	Message string `json:"message" example:"interests deleted successfully"`
}
