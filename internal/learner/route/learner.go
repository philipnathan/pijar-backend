package learner

import (
	"github.com/gin-gonic/gin"
	initLearner "github.com/philipnathan/pijar-backend/internal/learner"
	middleware "github.com/philipnathan/pijar-backend/middleware"
	"gorm.io/gorm"
)

func LearnerRoute(r *gin.Engine, db *gorm.DB) {
	apiV1 := "/api/v1/learners"

	handler, err := initLearner.InitializedLearner(db)
	if err != nil {
		panic(err)
	}

	protectedRoutes := r.Group(apiV1)
	{
		protectedRoutes.Use(middleware.AuthMiddleware())
		protectedRoutes.GET("/interests", handler.GetLearnerInterests)
		protectedRoutes.POST("/interests", handler.AddLearnerInterests)
		protectedRoutes.DELETE("/interests", handler.DeleteLearnerInterests)
	}
}
