package mentor

type GetMentorBioResponseDto struct {
	Message string `json:"message" example:"bio fetched successfully"`
	Bio     string `json:"bio" example:"My bio"`
}
