package mentor_session_participant

import (
	"net/http"

	"github.com/gin-gonic/gin"
	custom_error "github.com/philipnathan/pijar-backend/internal/mentor_session_participant/custom_error"
	dto "github.com/philipnathan/pijar-backend/internal/mentor_session_participant/dto"
	service "github.com/philipnathan/pijar-backend/internal/mentor_session_participant/service"
)

type MentorSessionParticipantHandler struct {
	service service.MentorSessionParticipantServiceInterface
}

func NewMentorSessionParticipantHandler(service service.MentorSessionParticipantServiceInterface) *MentorSessionParticipantHandler {
	return &MentorSessionParticipantHandler{
		service: service,
	}
}

func (h *MentorSessionParticipantHandler) CreateMentorSessionParticipantHandler(c *gin.Context) {
	UserID, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusUnauthorized, custom_error.NewCustomError("Unauthorized"))
	}

	_, ok := UserID.(float64)
	if !ok {
		c.JSON(http.StatusUnauthorized, custom_error.NewCustomError("Unauthorized"))
	}

	// Get all params
	var input dto.RegistrationRequestDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, custom_error.NewCustomError("Invalid request body"))
		return
	}

	err := h.service.CreateMentorSessionParticipant(
		UserID.(*uint),
		input.MentorSessionID,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Mentor session participant created successfully"})
}
