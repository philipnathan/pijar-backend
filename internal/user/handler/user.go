package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	custom_error "github.com/philipnathan/pijar-backend/internal/user/custom_error"
	model "github.com/philipnathan/pijar-backend/internal/user/model"
	service "github.com/philipnathan/pijar-backend/internal/user/service"
)

type userResponse struct {
	ID          uint    `json:"id"`
	Email       string  `json:"email"`
	Fullname    string  `json:"fullname"`
	BirthDate   string  `json:"birth_date"`
	PhoneNumber string  `json:"phone_number"`
	IsMentor    *bool   `json:"is_mentor"`
	ImageURL    *string `json:"image_url"`
}

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

	err := h.service.RegisterUserService(&user)

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

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
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

	c.JSON(http.StatusOK, dto.LoginUserResponseDto{
		Message:      "user logged in successfully",
		AccessToken:  access_token,
		RefreshToken: refresh_token})
}

// @Summary	Get user details
// @Schemes
// @Description	Get user details
// @Tags			User
// @Security		Bearer
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
		IsMentor:    user.IsMentor,
		ImageURL:    user.ImageURL,
	}

	c.JSON(http.StatusOK, userResponse)
}

func (h *UserHandler) DeleteUserHandler(c *gin.Context) {
	userID, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	id, ok := userID.(float64)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	err := h.service.DeleteUserService(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func (h *UserHandler) UpdateUserPasswordHandler(c *gin.Context) {
	userID, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	id, ok := userID.(float64)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var input struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.UpdateUserPasswordService(uint(id), input.OldPassword, input.NewPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
}

func (h *UserHandler) UpdateUserDetailsHandler(c *gin.Context) {
	userID, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	id, ok := userID.(float64)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var input struct {
		Fullname    string           `json:"fullname"`
		BirthDate   model.CustomTime `json:"birth_date"`
		PhoneNumber string           `json:"phone_number"`
		ImageURL    string           `json:"image_url"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.UpdateUserDetailsService(uint(id), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User details updated successfully"})
}