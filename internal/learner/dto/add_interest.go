package learner

type AddLearnerInterestsDto struct {
	CategoryID []uint `json:"category_id" binding:"required" example:"1,2,3"`
}

type AddLearnerInterestsResponseDto struct {
	Message string `json:"message" example:"interests added successfully"`
}
