package user

type GetUserResponseDto struct {
	ID          uint    `json:"id" example:"1"`
	Email       string  `json:"email" example:"test@example.com"`
	Fullname    string  `json:"fullname" example:"John Doe"`
	BirthDate   string  `json:"birth_date" example:"1990-01-01"`
	PhoneNumber string  `json:"phone_number" example:"123456789"`
	IsMentor    *bool   `json:"is_mentor" example:"true"`
	ImageURL    *string `json:"image_url" example:"https://example.com/image.jpg"`
}
