package user

type UpdateUserDetailsDto struct {
	Fullname    string  `json:"fullname" example:"John Doe" binding:"omitempty"`
	BirthDate   string  `json:"birth_date" example:"1990-02-01" binding:"omitempty"`
	PhoneNumber *string `json:"phone_number" example:"123456789" binding:"omitempty"`
	ImageURL    *string `json:"image_url" example:"https://example.com/image.jpg" binding:"omitempty"`
}

type UpdateUserDetailsResponseDto struct {
	Message string `json:"message" example:"user details updated successfully"`
}
