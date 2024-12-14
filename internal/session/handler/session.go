package mentor

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	custom_error "github.com/philipnathan/pijar-backend/internal/session/custom_error"
	service "github.com/philipnathan/pijar-backend/internal/session/service"
	middleware "github.com/philipnathan/pijar-backend/middleware"
)

type SessionHandler struct {
	service service.SessionService
}

func NewSessionHandler(service service.SessionService) *SessionHandler {
	return &SessionHandler{service: service}
}

func (h *SessionHandler) GetSessions(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Query("user_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	sessions, err := h.service.GetSessions(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"sessions": sessions})
}
