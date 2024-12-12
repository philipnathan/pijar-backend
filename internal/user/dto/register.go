package user

type RegisterUserDto struct {
	Email    string `json:"email" example:"test@example.com" binding:"required"`
	Password string `json:"password" example:"password123" binding:"required"`
	Fullname string `json:"fullname" example:"John Doe" binding:"required"`
}

type RegisterUserResponseDto struct {
	Message      string `json:"message" example:"user registered successfully"`
	AccessToken  string `json:"access_token" example:"eyAsgh435789"`
	RefreshToken string `json:"refresh_token" example:"eyAsgh435789"`
}
