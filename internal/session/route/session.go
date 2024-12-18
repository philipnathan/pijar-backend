package route

import (
    "github.com/gin-gonic/gin"
    handler "github.com/philipnathan/pijar-backend/internal/session/handler"
    repository "github.com/philipnathan/pijar-backend/internal/session/repository"
    service "github.com/philipnathan/pijar-backend/internal/session/service"
    "gorm.io/gorm"
)

func SessionRoute(r *gin.Engine, db *gorm.DB) {
    repo := repository.NewSessionRepository(db)
    srv := service.NewSessionService(repo)
    hnd := handler.NewSessionHandler(srv)

    r.GET("/api/v1/sessions/:user_id", hnd.GetSessions)
    r.GET("/api/v1/sessions/upcoming", hnd.GetUpcomingSessions)
}