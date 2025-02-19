package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	custom_error "github.com/philipnathan/pijar-backend/internal/user/custom_error"
	dto "github.com/philipnathan/pijar-backend/internal/user/dto"
	service "github.com/philipnathan/pijar-backend/internal/user/service"
	"github.com/philipnathan/pijar-backend/utils"
)

type MentorUserHandlerInterface interface {
	RegisterMentor(c *gin.Context)
}

type MentorUserHandler struct {
	service service.MentorUserServiceInterface
}

func NewMentorUserHandler(service service.MentorUserServiceInterface) MentorUserHandlerInterface {
	return &MentorUserHandler{
		service: service,
	}
}

// @Summary		Register mentor
// @Description	Register mentor
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			user	body		RegisterMentorDto	true	"User"
// @Success		200		{object}	RegisterMentorResponse
// @Failure		400		{object}	Error
// @Failure		500		{object}	Error
// @Router			/users/registermentor [post]
func (s *MentorUserHandler) RegisterMentor(c *gin.Context) {
	var request dto.RegisterMentorDto

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	access_token, refresh_token, err := s.service.RegisterMentor(&request.Email, &request.Password, &request.Fullname)

	if err != nil {
		switch err {
		case custom_error.ErrAlreadyMentor, custom_error.ErrChangeDetails:
			c.JSON(http.StatusBadRequest, custom_error.Error{Error: err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, custom_error.Error{Error: err.Error()})
			return
		}
	}

	utils.SetCookie(c, access_token, refresh_token)

	c.JSON(http.StatusOK, dto.RegisterMentorResponse{
		Message: "mentor registered successfully",
	})
}
