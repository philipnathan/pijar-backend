package learner

type AddLearnerInterestsDto struct {
	CategoryID []uint `json:"category_id" binding:"required" example:"[1,2,3]"`
}

type AddLearnerInterestsResponseDto struct {
	CategoryID   uint   `json:"category_id" example:"1"`
	CategoryName string `json:"category_name" example:"Development"`
}
