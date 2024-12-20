package session

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	custom_error "github.com/philipnathan/pijar-backend/internal/session/custom_error"
	dto "github.com/philipnathan/pijar-backend/internal/session/dto"
	service "github.com/philipnathan/pijar-backend/internal/session/service"
)

type SessionHandler struct {
	service service.SessionService
}

func NewSessionHandler(service service.SessionService) *SessionHandler {
	return &SessionHandler{service: service}
}

// @Summary		Get sessions for a user
// @Description	Get all sessions for a specific user by user ID
// @Tags			Mentor
// @Accept			json
// @Produce		json
// @Param			user_id	path		int	true	"User ID"
// @Success		200		{array}		MentorSessionResponse
// @Failure		400		{object}	Error	"Invalid user ID"
// @Failure		500		{object}	Error	"Internal server error"
// @Router			/sessions/{user_id} [get]
func (h *SessionHandler) GetSessions(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, custom_error.Error{Error: "Invalid user ID"})
		return
	}

	sessions, err := h.service.GetSessions(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, custom_error.Error{Error: err.Error()})
		return
	}

	var response []dto.MentorSessionResponse
	var mentorImageURL string
	var registered bool = false

	for _, session := range sessions {
		if session.User.ImageURL != nil {
			mentorImageURL = *session.User.ImageURL
		} else {
			mentorImageURL = ""
		}

		response = append(response, dto.MentorSessionResponse{
			MentorSessionTitle: session.Title,
			ShortDescription:   session.ShortDescription,
			Schedule:           session.Schedule,
			Registered:         registered,
			MentorDetails: dto.MentorDetails{
				Id:       session.User.ID,
				Fullname: session.User.Fullname,
				ImageURL: mentorImageURL,
			},
		})
	}

	c.JSON(http.StatusOK, response)
}

// @Summary Get upcoming sessions
// @Description Get all upcoming sessions
// @Tags Session
// @Produce json
// @Success 200 {object} session.GetUpcomingSessionResponse
// @Failure 500 {object} Error "Internal server error"
// @Router /sessions/upcoming [get]
func (h *SessionHandler) GetUpcomingSessions(c *gin.Context) {
	sessions, err := h.service.GetUpcomingSessions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var sessionDetails []dto.SessionDetail
	for _, session := range sessions {
		sessionDetails = append(sessionDetails, dto.SessionDetail{
			Day:              session.Schedule.Weekday().String(),
			Time:             session.Schedule.Format("03:04 PM"),
			Title:            session.Title,
			ShortDescription: session.ShortDescription,
			Schedule:         session.Schedule.Format(time.RFC3339),
			ImageURL:         session.ImageURL,
			Registered:       true,
			Duration:         session.EstimateDuration,
		})
	}
	c.JSON(http.StatusOK, dto.GetUpcomingSessionResponse{Sessions: sessionDetails})
}

// @Summary		Get learner history session
// @Description	Get learner history session
// @Tags			Learner
// @Produce		json
// @Security		Bearer
// @Success		200	{object}	GetUserHistorySessionResponseDto
// @Failure		401	{object}	Error	"Unauthorized"
// @Failure		500	{object}	Error	"Internal server error"
// @Router			/sessions/histories [get]
func (h *SessionHandler) GetLearnerHistorySession(c *gin.Context) {
	UserID, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusUnauthorized, custom_error.Error{Error: "Unauthorized"})
		return
	}

	id, ok := UserID.(float64)
	if !ok {
		c.JSON(http.StatusUnauthorized, custom_error.Error{Error: "Unauthorized"})
		return
	}

	histories, err := h.service.GetLearnerHistorySession(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, custom_error.Error{Error: err.Error()})
		return
	}

	var response dto.GetUserHistorySessionResponseDto
	var mentorImageURL string

	for _, history := range *histories {
		if history.MentorSession.Schedule.After(time.Now()) {
			continue
		}
		if history.MentorSession.User.ImageURL != nil {
			mentorImageURL = *history.MentorSession.User.ImageURL
		} else {
			mentorImageURL = ""
		}

		response.Histories = append(response.Histories, dto.History{
			MentorSessionTitle: history.MentorSession.Title,
			ShortDescription:   history.MentorSession.ShortDescription,
			Schedule:           history.MentorSession.Schedule,
			Status:             string(history.Status),
			MentorDetails: dto.MentorDetails{
				Id:       history.MentorSession.User.ID,
				Fullname: history.MentorSession.User.Fullname,
				ImageURL: mentorImageURL,
			},
		})
	}

	c.JSON(http.StatusOK, response)
}
