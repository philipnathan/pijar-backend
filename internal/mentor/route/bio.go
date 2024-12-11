package mentor

import (
	"github.com/gin-gonic/gin"
	middleware "github.com/philipnathan/pijar-backend/middleware"
	"gorm.io/gorm"

	handler "github.com/philipnathan/pijar-backend/internal/mentor/handler"
	repository "github.com/philipnathan/pijar-backend/internal/mentor/repository"
	service "github.com/philipnathan/pijar-backend/internal/mentor/service"
)

func MentorBioRoute(r *gin.Engine, db *gorm.DB) {
	apiV1 := "/api/v1/mentors/biographies"

	repo := repository.NewMentorBioRepository(db)
	services := service.NewMentorBioService(repo)
	handler := handler.NewMentorBioHandler(services)

	protectedRoutes := r.Group(apiV1)
	{
		protectedRoutes.Use(middleware.AuthMiddleware())
		protectedRoutes.GET("/", handler.MentorGetBio)
	}
}
