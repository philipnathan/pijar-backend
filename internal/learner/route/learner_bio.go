package learner

import (
	"github.com/gin-gonic/gin"
	initLearnerBio "github.com/philipnathan/pijar-backend/internal/learner"
	middleware "github.com/philipnathan/pijar-backend/middleware"
	"gorm.io/gorm"
)

func LearnerBioRoute(r *gin.Engine, db *gorm.DB) {
	apiV1 := "/api/v1/learners/biographies"

	handler, err := initLearnerBio.InitializedLearnerBio(db)
	if err != nil {
		panic(err)
	}

	protectedRoutes := r.Group(apiV1)
	{
		protectedRoutes.Use(middleware.AuthMiddleware())
		protectedRoutes.POST("/", handler.CreateLearnerBio)
		protectedRoutes.GET("/", handler.GetLearnerBio)
		protectedRoutes.PUT("/", handler.UpdateLearnerBio)
	}
}
