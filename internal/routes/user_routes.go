package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/philipnathan/pijar-backend/internal/handlers"
	"github.com/philipnathan/pijar-backend/internal/repositories"
	"github.com/philipnathan/pijar-backend/internal/services"
	"github.com/philipnathan/pijar-backend/middleware"
	"gorm.io/gorm"
)

func UserRoute(r *gin.Engine, db *gorm.DB) {
	apiV1 := "/api/v1/users"

	repo := repositories.NewUserRepository(db)
	service := services.NewUserService(repo)
	handler := handlers.NewUserHandler(service)


	userRoutes := r.Group(apiV1)
	{
		userRoutes.POST("/register", handler.RegisterUser)
		userRoutes.POST("/login", handler.LoginUser)
	}

	protectedRoutes := r.Group(apiV1)
	{
		protectedRoutes.Use(middleware.AuthMiddleware())
		protectedRoutes.GET("/me", handler.GetUser)
	}
}