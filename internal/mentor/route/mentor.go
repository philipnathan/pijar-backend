package mentor

import (
	"github.com/gin-gonic/gin"
	learnerInterestsRepo "github.com/philipnathan/pijar-backend/internal/learner/repository"
	repository "github.com/philipnathan/pijar-backend/internal/mentor/repository"
	service "github.com/philipnathan/pijar-backend/internal/mentor/service"
	middleware "github.com/philipnathan/pijar-backend/middleware"
	"gorm.io/gorm"

	mentor "github.com/philipnathan/pijar-backend/internal/mentor/handler"
)

func MentorBioRoute(r *gin.Engine, db *gorm.DB) {
	apiV1 := "/api/v1/mentors"

	repo := repository.NewMentorBioRepository(db)
	services := service.NewMentorBioService(repo)
	handler := mentor.NewMentorBioHandler(services)

	learnerInterestsRepo := learnerInterestsRepo.NewLearnerRepository(db)

	mentorRepo := repository.NewMentorRepository(db)
	mentorServices := service.NewMentorService(mentorRepo, learnerInterestsRepo)
	mentorHandler := mentor.NewMentorHandler(mentorServices)

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
