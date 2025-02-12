package mentor

import (
	"github.com/gin-gonic/gin"
	middleware "github.com/philipnathan/pijar-backend/middleware"
	"gorm.io/gorm"

	initMentor "github.com/philipnathan/pijar-backend/internal/mentor"
)

func MentorBioRoute(r *gin.Engine, db *gorm.DB) {
	apiV1 := "/api/v1/mentors"

	handler, err := initMentor.InitializedMentorBio(db)
	if err != nil {
		panic(err)
	}

	mentorHandler, err := initMentor.InitializedMentor(db)
	if err != nil {
		panic(err)
	}

	nonProtectedRoutes := r.Group(apiV1)
	{
		nonProtectedRoutes.GET("/:mentor_id", mentorHandler.UserGetMentorDetails)
		nonProtectedRoutes.GET("/landingpage", mentorHandler.UserGetMentorLandingPage)
	}

	protectedRoutes := r.Group(apiV1)
	{
		protectedRoutes.Use(middleware.AuthMiddleware())
		protectedRoutes.GET("/me/bio", handler.MentorGetBio)
	}
}
