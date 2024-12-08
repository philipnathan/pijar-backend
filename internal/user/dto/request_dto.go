package user

type RegisterUserDto struct {
	Email    string `json:"email" example:"test@example.com"`
	Password string `json:"password" example:"password123"`
	Fullname string `json:"fullname" example:"John Doe"`
}
