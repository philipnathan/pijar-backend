package user

type RegisterMentorDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Fullname string `json:"fullname" validate:"required"`
}

type RegisterMentorResponse struct {
	Message      string `json:"message" example:"mentor registered successfully"`
	AccessToken  string `json:"access_token" example:"eyAsgh435789"`
	RefreshToken string `json:"refresh_token" example:"eyAsgh435789"`
}
