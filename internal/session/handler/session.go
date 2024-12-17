package session

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	dto "github.com/philipnathan/pijar-backend/internal/session/dto"
	service "github.com/philipnathan/pijar-backend/internal/session/service"
)

type SessionHandler struct {
	service service.SessionService
}

func NewSessionHandler(service service.SessionService) *SessionHandler {
	return &SessionHandler{service: service}
}

// GetSessions godoc
// @Summary Get sessions for a user
// @Description Get all sessions for a specific user by user ID
// @Tags Sessions
// @Accept  json
// @Produce  json
// @Param user_id path int true "User ID"
// @Success 200 {array} session.MentorSession
// @Failure 400 {object} gin.H{"error": "Invalid user ID"}
// @Failure 500 {object} gin.H{"error": "Internal Server Error"}
// @Router /sessions/{user_id} [get]
func (h *SessionHandler) GetSessions(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	sessions, err := h.service.GetSessions(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
