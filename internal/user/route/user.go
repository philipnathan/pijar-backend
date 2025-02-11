package user

import (
	"github.com/gin-gonic/gin"
	userInit "github.com/philipnathan/pijar-backend/internal/user"
	"github.com/philipnathan/pijar-backend/middleware"
	"gorm.io/gorm"
)

func UserRoute(r *gin.Engine, db *gorm.DB) {
	apiV1 := "/api/v1/users"

	handler, err := userInit.InitializedUser(db)
	if err != nil {
		panic(err)
	}

	userRoutes := r.Group(apiV1)
	{
		userRoutes.POST("/register", handler.RegisterUser)
		userRoutes.POST("/login", handler.LoginUser)
		userRoutes.POST("/logout", handler.UserLogout)
	}

	protectedRoutes := r.Group(apiV1)
	{
		protectedRoutes.Use(middleware.AuthMiddleware())
		protectedRoutes.GET("/me", handler.GetUser)
		protectedRoutes.DELETE("/me", handler.DeleteUserHandler)
		protectedRoutes.PATCH("/me/password", handler.UpdateUserPasswordHandler)
		protectedRoutes.PATCH("/me/details", handler.UpdateUserDetailsHandler)
		protectedRoutes.GET("me/profile", handler.GetUserProfile)
	}
}
