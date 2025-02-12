package session_review

import (
	"github.com/gin-gonic/gin"
	initSessionReview "github.com/philipnathan/pijar-backend/internal/session_review"
	middleware "github.com/philipnathan/pijar-backend/middleware"
	"gorm.io/gorm"
)

func SessionReviewRoute(r *gin.Engine, db *gorm.DB) {
	apiV1 := "/api/v1/sessions"

	hnd, err := initSessionReview.InitializedSessionReview(db)
	if err != nil {
		panic(err)
	}

	nonProtectedRoutes := r.Group(apiV1)
	{
		nonProtectedRoutes.GET("/:session_id/review", hnd.GetSessionReviewsHandler)
	}

	protectedRoutes := r.Group(apiV1)
	{
		protectedRoutes.Use(middleware.AuthMiddleware())
		protectedRoutes.POST("/:session_id/review", hnd.CreateSessionReviewHandler)
	}
}
