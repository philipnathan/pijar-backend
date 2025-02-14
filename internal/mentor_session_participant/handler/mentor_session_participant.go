package mentor_session_participant

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	custom_error "github.com/philipnathan/pijar-backend/internal/mentor_session_participant/custom_error"
	dto "github.com/philipnathan/pijar-backend/internal/mentor_session_participant/dto"
	service "github.com/philipnathan/pijar-backend/internal/mentor_session_participant/service"
)

type MentorSessionParticipantHandlerInterface interface {
	CreateMentorSessionParticipantHandler(c *gin.Context)
	GetLearnerEnrollmentsHandler(c *gin.Context)
}

type MentorSessionParticipantHandler struct {
	service service.MentorSessionParticipantServiceInterface
}

func NewMentorSessionParticipantHandler(service service.MentorSessionParticipantServiceInterface) MentorSessionParticipantHandlerInterface {
	return &MentorSessionParticipantHandler{
		service: service,
	}
}

// @Summary	Used for learner to join a mentor session
// @Schemes
// @Description	Used for learner to join a mentor session
// @Tags			Session Enrollments
// @Produce		json
// @Param			session_id	path		int	true	"Session ID"
// @Success		200			{object}	RegistrationResponse
// @Failure		400			{object}	CustomError	"Invalid session ID"
// @Failure		500			{object}	CustomError	"Internal server error"
// @Router			/sessions/{session_id}/enroll [post]
func (h *MentorSessionParticipantHandler) CreateMentorSessionParticipantHandler(c *gin.Context) {
	UserID, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusUnauthorized, custom_error.NewCustomError("Unauthorized"))
	}

	userIDFloat, ok := UserID.(float64)
	if !ok {
		c.JSON(http.StatusUnauthorized, custom_error.NewCustomError("Unauthorized"))
	}

	// Get all params
	SessionIDStr := c.Param("session_id")
	if SessionIDStr == "" {
		c.JSON(http.StatusBadRequest, custom_error.ErrSessionNotFound)
		return
	}

	SessionID, err := strconv.ParseUint(SessionIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, custom_error.ErrSessionNotFound)
		return
	}

	uintSessionID := uint(SessionID)
	uintUserID := uint(userIDFloat)
	ctx := c.Request.Context()

	err = h.service.CreateMentorSessionParticipant(
		ctx,
		&uintUserID,
		&uintSessionID,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, custom_error.NewCustomError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, dto.RegistrationResponse{Message: "Successfully registered"})
}

// @Summary	Get learner enrollments
// @Schemes
// @Description	Get learner enrollments
// @Tags			Session Enrollments
// @Produce		json
// @Param			page		query		int	false	"Page number"
// @Param			page_size	query		int	false	"Page size"
// @Success		200			{object}	EnrollmentResponse
// @Failure		500			{object}	CustomError	"Internal server error"
// @Router			/sessions/enrollments [get]
func (h *MentorSessionParticipantHandler) GetLearnerEnrollmentsHandler(c *gin.Context) {
	UserID, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusUnauthorized, custom_error.NewCustomError("Unauthorized"))
	}

	userIDFloat, ok := UserID.(float64)
	if !ok {
		c.JSON(http.StatusUnauthorized, custom_error.NewCustomError("Unauthorized"))
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "10"))

	uintUserID := uint(userIDFloat)

	data, total, err := h.service.GetLearnerEnrollments(&uintUserID, &page, &pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, custom_error.NewCustomError(err.Error()))
		return
	}

	var details []dto.EnrollmentDetails
	var sessionImage string
	for _, d := range *data {
		if d.MentorSession.ImageURL == "" {
			sessionImage = ""
		} else {
			sessionImage = d.MentorSession.ImageURL
		}

		details = append(details, dto.EnrollmentDetails{
			MentorSessionParticipantID: d.ID,
			SessionDetails: dto.SessionDetails{
				MentorSessionID:    d.MentorSession.ID,
				MentorSessionTitle: d.MentorSession.Title,
				ShortDescription:   d.MentorSession.ShortDescription,
				ImageURL:           sessionImage,
				Schedule:           d.MentorSession.Schedule,
			},
			Status: string(d.Status),
		})
	}

	response := dto.EnrollmentResponse{
		Message:     "Enrollments fetched successfully",
		Enrollments: details,
		Total:       total,
		Page:        page,
		PageSize:    pageSize,
	}

	c.JSON(http.StatusOK, response)
}
