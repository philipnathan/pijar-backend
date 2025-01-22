package user

import (
	"github.com/gin-gonic/gin"
	handler "github.com/philipnathan/pijar-backend/internal/user/handler"
	repo "github.com/philipnathan/pijar-backend/internal/user/repository"
	service "github.com/philipnathan/pijar-backend/internal/user/service"
	"gorm.io/gorm"
)

func GoogleAuthRoute(r *gin.Engine, db *gorm.DB) {
	apiV1 := "/api/v1/"

	repo := repo.NewGoogleAuthRepo(db)
	services := service.NewGoogleAuthService(repo)
	handler := handler.NewGoogleAuthHandler(services)

	nonProtected := r.Group(apiV1)
	{
		nonProtected.GET("/auth/google/register", handler.GoogleRegisterCallback)
		nonProtected.GET("/auth/google/login", handler.GoogleLoginCallback)
	}
}
