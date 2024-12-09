package user

type ChangePasswordDto struct {
	OldPassword string `json:"old_password" example:"oldPassword123" binding:"required"`
	NewPassword string `json:"new_password" example:"newPassword123" binding:"required"`
}

type ChangePasswordResponseDto struct {
	Message string `json:"message" example:"password changed successfully"`
}
