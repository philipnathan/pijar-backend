package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	custom_error "github.com/philipnathan/pijar-backend/internal/user/custom_error"
	dto "github.com/philipnathan/pijar-backend/internal/user/dto"
	service "github.com/philipnathan/pijar-backend/internal/user/service"
	"github.com/philipnathan/pijar-backend/utils"
)

type UserHandler struct {
	service service.UserServiceInterface
}

func NewUserHandler(service service.UserServiceInterface) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

// @Summary	Register new user
// @Schemes
// @Description	Register new user
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			user	body		RegisterUserDto	true	"User"
// @Success		200		{object}	RegisterUserResponseDto
// @Failure		400		{object}	Error	"Invalid request body"
// @Failure		500		{object}	Error	"Internal server error"
// @Router			/users/register [post]
func (h *UserHandler) RegisterUser(c *gin.Context) {
	var user dto.RegisterUserDto

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	access_token, refresh_token, err := h.service.RegisterUserService(&user.Email, &user.Password, &user.Fullname)

	if err != nil {
		switch err {
		case custom_error.ErrEmailExist, custom_error.ErrPhoneNumberExist:
			c.JSON(http.StatusBadRequest, custom_error.Error{Error: err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, custom_error.Error{Error: err.Error()})
			return
		}
	}

	utils.SetCookie(c, access_token, refresh_token)

	response := dto.RegisterUserResponseDto{
		Message: "user registered successfully",
	}

	c.JSON(http.StatusOK, response)
}

// @Summary	Login user
// @Schemes
// @Description	Login user
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			user	body		LoginUserDto	true	"Login Information"
// @Success		200		{object}	LoginUserResponseDto
// @Failure		400		{object}	Error	"Invalid request body"
// @Failure		500		{object}	Error	"Internal server error"
// @Router			/users/login [post]
func (h *UserHandler) LoginUser(c *gin.Context) {
	var input dto.LoginUserDto

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, custom_error.Error{Error: err.Error()})
		return
	}

	var access_token, refresh_token string

	access_token, refresh_token, err := h.service.LoginUserService(input.Email, input.Password)
	if err != nil {
		switch err {
		case custom_error.ErrLogin:
			c.JSON(http.StatusUnauthorized, custom_error.Error{Error: err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, custom_error.Error{Error: err.Error()})
			return
		}
	}

	utils.SetCookie(c, access_token, refresh_token)

	c.JSON(http.StatusOK, dto.LoginUserResponseDto{
		Message: "user logged in successfully"})
}

// @Summary	Get user details
// @Schemes
// @Description	Get user details
// @Tags			User
// @Accept			json
// @Produce		json
// @Success		200	{object}	GetUserResponseDto
// @Failure		401	{object}	Error	"Unauthorized"
// @Failure		500	{object}	Error	"Internal server error"
// @Router			/users/me [get]
func (h *UserHandler) GetUser(c *gin.Context) {
	userID, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusUnauthorized, custom_error.Error{Error: "Unauthorized"})
		return
	}

	id, ok := userID.(float64)
	if !ok {
		c.JSON(http.StatusUnauthorized, custom_error.Error{Error: "Unauthorized"})
		return
	}

	user, err := h.service.GetUserDetails(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, custom_error.Error{Error: err.Error()})
		return
	}

	var formattedBirthDate string
	if user.BirthDate != nil {
		formattedBirthDate = user.BirthDate.Format("2006-01-02")
	} else {
		formattedBirthDate = ""
	}

	var formattedPhoneNumber string
	if user.PhoneNumber != nil {
		formattedPhoneNumber = *user.PhoneNumber
	} else {
		formattedPhoneNumber = ""
	}

	userResponse := dto.GetUserResponseDto{
		ID:          user.ID,
		Email:       user.Email,
		Fullname:    user.Fullname,
		BirthDate:   formattedBirthDate,
		PhoneNumber: formattedPhoneNumber,
		IsLearner:   user.IsLearner,
		IsMentor:    user.IsMentor,
		ImageURL:    user.ImageURL,
	}

	c.JSON(http.StatusOK, userResponse)
}

// @Summary	Delete user
// @Schemes
// @Description	Delete user
// @Tags			User
// @Produce		json
// @Success		200	{object}	DeleteUserResponseDto
// @Failure		401	{object}	Error	"Unauthorized"
// @Failure		500	{object}	Error	"Internal server error"
// @Router			/users/me [delete]
func (h *UserHandler) DeleteUserHandler(c *gin.Context) {
	userID, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusUnauthorized, custom_error.Error{Error: "Unauthorized"})
		return
	}

	id, ok := userID.(float64)
	if !ok {
		c.JSON(http.StatusUnauthorized, custom_error.Error{Error: "Unauthorized"})
		return
	}

	err := h.service.DeleteUserService(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, custom_error.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.DeleteUserResponseDto{Message: "user deleted successfully"})
}

// @Summary	Update user password
// @Schemes
// @Description	Update user password
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			password	body		ChangePasswordDto	true	"User"
// @Success		200			{object}	ChangePasswordResponseDto
// @Failure		400			{object}	Error	"Invalid request body"
// @Failure		500			{object}	Error	"Internal server error"
// @Router			/users/me/password [patch]
func (h *UserHandler) UpdateUserPasswordHandler(c *gin.Context) {
	userID, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusUnauthorized, custom_error.Error{Error: "Unauthorized"})
		return
	}

	id, ok := userID.(float64)
	if !ok {
		c.JSON(http.StatusUnauthorized, custom_error.Error{Error: "Unauthorized"})
		return
	}

	var input dto.ChangePasswordDto

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, custom_error.Error{Error: err.Error()})
		return
	}

	err := h.service.UpdateUserPasswordService(uint(id), input.OldPassword, input.NewPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, custom_error.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ChangePasswordResponseDto{Message: "password changed successfully"})
}

// @Summary	Update user details
// @Schemes
// @Description	Update user details
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			user	body		UpdateUserDetailsDto	true	"User"
// @Success		200		{object}	UpdateUserDetailsResponseDto
// @Failure		400		{object}	Error	"Invalid request body"
// @Failure		500		{object}	Error	"Internal server error"
// @Router			/users/me/details [patch]
func (h *UserHandler) UpdateUserDetailsHandler(c *gin.Context) {
	userID, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusUnauthorized, custom_error.Error{Error: "Unauthorized"})
		return
	}

	id, ok := userID.(float64)
	if !ok {
		c.JSON(http.StatusUnauthorized, custom_error.Error{Error: "Unauthorized"})
		return
	}

	var input dto.UpdateUserDetailsDto

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, custom_error.Error{Error: err.Error()})
		return
	}

	err := h.service.UpdateUserDetailsService(uint(id), input)
	if err != nil {
		switch err {
		case custom_error.ErrUserNotFound, custom_error.ErrStatusCannotBeFalse:
			c.JSON(http.StatusNotFound, custom_error.Error{Error: err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, custom_error.Error{Error: err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, dto.UpdateUserDetailsResponseDto{Message: "user details updated successfully"})
}

// @Summary		Get user profile
// @Description	Get the profile of the logged-in user
// @Tags			User
// @Produce		json
// @Success		200	{object}	user.UserProfileResponse
// @Failure		401	{object}	Error	"Unauthorized"
// @Failure		500	{object}	Error	"Internal server error"
// @Router			/users/me/profile [get]
func (h *UserHandler) GetUserProfile(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	profile, err := h.service.GetUserProfile(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, profile)
}

// @Summary		Logout user
// @Description	Logout the user
// @Tags			User
// @Produce		json
// @Success		200	{object} object{message=string}	"User logged out successfully"
// @Router			/users/logout [post]
func (h *UserHandler) UserLogout(c *gin.Context) {
	utils.DeleteCookie(c)

	c.JSON(http.StatusOK, gin.H{"message": "user logged out successfully"})
}
