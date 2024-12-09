package learner

import (
	"github.com/gin-gonic/gin"
	handler "github.com/philipnathan/pijar-backend/internal/learner/handler"
	repository "github.com/philipnathan/pijar-backend/internal/learner/repository"
	service "github.com/philipnathan/pijar-backend/internal/learner/service"
	middleware "github.com/philipnathan/pijar-backend/middleware"
	"gorm.io/gorm"
)

func LearnerRoute(r *gin.Engine, db *gorm.DB) {
	apiV1 := "/api/v1/learners"

	repo := repository.NewLearnerRepository(db)
	services := service.NewLearnerService(repo)
	handler := handler.NewLearnerHandler(services)

	protectedRoutes := r.Group(apiV1)
	{
		protectedRoutes.Use(middleware.AuthMiddleware())
		protectedRoutes.GET("/interests", handler.GetLearnerInterests)
		protectedRoutes.POST("/interests", handler.AddLearnerInterests)
	}
}
