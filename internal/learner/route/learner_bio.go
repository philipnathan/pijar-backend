package learner

import (
	"github.com/gin-gonic/gin"
	handler "github.com/philipnathan/pijar-backend/internal/learner/handler"
	repository "github.com/philipnathan/pijar-backend/internal/learner/repository"
	service "github.com/philipnathan/pijar-backend/internal/learner/service"
	middleware "github.com/philipnathan/pijar-backend/middleware"
	"gorm.io/gorm"
)

func LearnerBioRoute(r *gin.Engine, db *gorm.DB) {
	apiV1 := "/api/v1/learners/biographies"

	repo := repository.NewLearnerBioRepository(db)
	services := service.NewLearnerBioService(repo)
	handler := handler.NewLearnerBioHandler(services)

	protectedRoutes := r.Group(apiV1)
	{
		protectedRoutes.Use(middleware.AuthMiddleware())
		protectedRoutes.POST("/", handler.CreateLearnerBio)
		protectedRoutes.GET("/", handler.GetLearnerBio)
	}
}
