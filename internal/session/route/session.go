package session

import (
	"github.com/gin-gonic/gin"
	initSession "github.com/philipnathan/pijar-backend/internal/session"
	middleware "github.com/philipnathan/pijar-backend/middleware"
	"gorm.io/gorm"
)

func SessionRoute(r *gin.Engine, db *gorm.DB) {
	apiV1 := "/api/v1/sessions"

	hnd, err := initSession.InitializedSession(db)
	if err != nil {
		panic(err)
	}

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
