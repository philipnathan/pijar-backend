package follow

import (
	"github.com/gin-gonic/gin"
	handler "github.com/philipnathan/pijar-backend/internal/follow/handler"
	repo "github.com/philipnathan/pijar-backend/internal/follow/repository"
	service "github.com/philipnathan/pijar-backend/internal/follow/service"
	userRepo "github.com/philipnathan/pijar-backend/internal/user/repository"
	userService "github.com/philipnathan/pijar-backend/internal/user/service"
	middleware "github.com/philipnathan/pijar-backend/middleware"
	"gorm.io/gorm"
)

func FollowRoute(r *gin.Engine, db *gorm.DB) {
	apiV1 := "/api/v1/mentors/"

	repo := repo.NewFollowRepository(db)
	userRepo := userRepo.NewUserRepository(db)
	userService := userService.NewUserService(userRepo)
	srv := service.NewFollowService(
		repo,
		userService,
	)
	hnd := handler.NewFollowHandler(srv)

	protectedRoutes := r.Group(apiV1)
	{
		protectedRoutes.Use(middleware.AuthMiddleware())

		protectedRoutes.POST("/:mentor_id/follow", hnd.FollowUnfollowHandler)
		protectedRoutes.GET("/:mentor_id/status", hnd.IsFollowingHandler)
	}
}
