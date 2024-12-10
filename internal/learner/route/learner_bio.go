package learner

import (
	"github.com/gin-gonic/gin"
	middleware "github.com/philipnathan/pijar-backend/middleware"
	"gorm.io/gorm"
)

func LearnerBioRoute(r *gin.Engine, db *gorm.DB) {
	apiV1 := "/api/v1/learners/bio"

	protectedRoutes := r.Group(apiV1)
	{
		protectedRoutes.Use(middleware.AuthMiddleware())
	}
}
