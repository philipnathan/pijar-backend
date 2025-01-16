package session

import (
	"github.com/gin-gonic/gin"
	learnerInterestsRepo "github.com/philipnathan/pijar-backend/internal/learner/repository"
	handler "github.com/philipnathan/pijar-backend/internal/session/handler"
	repository "github.com/philipnathan/pijar-backend/internal/session/repository"
	service "github.com/philipnathan/pijar-backend/internal/session/service"
	middleware "github.com/philipnathan/pijar-backend/middleware"
	"gorm.io/gorm"
)

func SessionRoute(r *gin.Engine, db *gorm.DB) {
	apiV1 := "/api/v1/sessions"

	learnerInterestsRepo := learnerInterestsRepo.NewLearnerRepository(db)

	repo := repository.NewSessionRepository(db)
	srv := service.NewSessionService(repo, learnerInterestsRepo)
	hnd := handler.NewSessionHandler(srv)

	nonProtectedRoutes := r.Group(apiV1)
	{
		// nonProtectedRoutes.GET("", hnd.GetSessions)
		nonProtectedRoutes.GET("/upcoming", hnd.GetUpcommingSessionsLandingPage)
		nonProtectedRoutes.GET("", hnd.GetAllSessionsWithFilter)
		nonProtectedRoutes.GET("/:session_id", hnd.GetSessionDetailById)
	}

	protectedRoutes := r.Group(apiV1)
	{
		protectedRoutes.Use(middleware.AuthMiddleware())
		protectedRoutes.GET("/histories", hnd.GetLearnerHistorySession)
	}
}
