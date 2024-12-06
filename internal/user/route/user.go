package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/philipnathan/pijar-backend/internal/user/handler"
	"github.com/philipnathan/pijar-backend/internal/user/repository"
	"github.com/philipnathan/pijar-backend/internal/user/service"
	"github.com/philipnathan/pijar-backend/middleware"
	"gorm.io/gorm"
)

func UserRoute(r *gin.Engine, db *gorm.DB) {
	apiV1 := "/api/v1/users"

	repo := repository.NewUserRepository(db)
	services := service.NewUserService(repo)
	handler := handler.NewUserHandler(services)


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