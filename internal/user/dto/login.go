package user

type LoginUserDto struct {
	Email    string `json:"email" example:"test@example.com" binding:"required"`
	Password string `json:"password" example:"password123" binding:"required"`
}

type LoginUserResponseDto struct {
	Message      string `json:"message" example:"user logged in successfully"`
	AccessToken  string `json:"access_token" example:"eyAsgh435789"`
	RefreshToken string `json:"refresh_token" example:"eyAsgh435789"`
}
