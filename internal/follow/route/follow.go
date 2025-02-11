package follow

import (
	"github.com/gin-gonic/gin"
	initFollow "github.com/philipnathan/pijar-backend/internal/follow"
	middleware "github.com/philipnathan/pijar-backend/middleware"
	"gorm.io/gorm"
)

func FollowRoute(r *gin.Engine, db *gorm.DB) {
	apiV1 := "/api/v1/mentors/"

	hnd, err := initFollow.InitializedFollow(db)
	if err != nil {
		panic(err)
	}

	protectedRoutes := r.Group(apiV1)
	{
		protectedRoutes.Use(middleware.AuthMiddleware())

		protectedRoutes.POST("/:mentor_id/follow", hnd.FollowUnfollowHandler)
		protectedRoutes.GET("/:mentor_id/status", hnd.IsFollowingHandler)
	}
}
