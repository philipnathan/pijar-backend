package session

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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

	c.JSON(http.StatusOK, sessions)
}
