package handler

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