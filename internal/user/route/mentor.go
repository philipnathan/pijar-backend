package user

import (
	"github.com/gin-gonic/gin"
	handler "github.com/philipnathan/pijar-backend/internal/user/handler"
	repo "github.com/philipnathan/pijar-backend/internal/user/repository"
	service "github.com/philipnathan/pijar-backend/internal/user/service"
	"gorm.io/gorm"
)

func MentorUserRoute(r *gin.Engine, db *gorm.DB) {
	apiV1 := "/api/v1/users"

	repo := repo.NewMentorUserRepository(db)
	services := service.NewMentorUserService(repo)
	handler := handler.NewMentorUserHandler(services)

	nonProtectedRoutes := r.Group(apiV1)
	{
		nonProtectedRoutes.POST("/registermentor", handler.RegisterMentor)
	}
}
