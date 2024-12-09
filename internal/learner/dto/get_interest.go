package learner

type InterestOnlyDto struct {
	CategoryID   uint   `json:"category_id" example:"1"`
	CategoryName string `json:"category_name" example:"Development"`
}

type GetLearnerInterestResponseDto struct {
	Message string            `json:"message" example:"interests retrieved successfully"`
	Data    []InterestOnlyDto `json:"data"`
}
