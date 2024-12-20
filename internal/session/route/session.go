package route

import (
	"github.com/gin-gonic/gin"
	handler "github.com/philipnathan/pijar-backend/internal/session/handler"
	repository "github.com/philipnathan/pijar-backend/internal/session/repository"
	service "github.com/philipnathan/pijar-backend/internal/session/service"
	middleware "github.com/philipnathan/pijar-backend/middleware"
	"gorm.io/gorm"
)

func SessionRoute(r *gin.Engine, db *gorm.DB) {
	apiV1 := "/api/v1/sessions"

	repo := repository.NewSessionRepository(db)
	srv := service.NewSessionService(repo)
	hnd := handler.NewSessionHandler(srv)

	nonProtectedRoutes := r.Group(apiV1)
	{
		nonProtectedRoutes.GET("/:user_id", hnd.GetSessions)
		nonProtectedRoutes.GET("/upcoming", hnd.GetUpcomingSessions)
	}

	protectedRoutes := r.Group(apiV1)
	{
		protectedRoutes.Use(middleware.AuthMiddleware())
		protectedRoutes.GET("/histories", hnd.GetLearnerHistorySession)
	}
}
