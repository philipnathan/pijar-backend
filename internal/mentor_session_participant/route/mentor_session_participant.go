package mentor_session_participant

import (
	"github.com/gin-gonic/gin"
	learnerRepo "github.com/philipnathan/pijar-backend/internal/learner/repository"
	handler "github.com/philipnathan/pijar-backend/internal/mentor_session_participant/handler"
	repo "github.com/philipnathan/pijar-backend/internal/mentor_session_participant/repository"
	service "github.com/philipnathan/pijar-backend/internal/mentor_session_participant/service"
	sessionRepository "github.com/philipnathan/pijar-backend/internal/session/repository"
	sessionService "github.com/philipnathan/pijar-backend/internal/session/service"
	userRepository "github.com/philipnathan/pijar-backend/internal/user/repository"
	userService "github.com/philipnathan/pijar-backend/internal/user/service"
	middleware "github.com/philipnathan/pijar-backend/middleware"
	"gorm.io/gorm"
)

func MentorSessionParticipantRoute(r *gin.Engine, db *gorm.DB) {
	apiV1 := "/api/v1/sessions/"

	repo := repo.NewMentorSessionParticipantRepository(db)
	userRepo := userRepository.NewUserRepository(db)
	sessionRepo := sessionRepository.NewSessionRepository(db)
	learnerRepo := learnerRepo.NewLearnerRepository(db)
	srv := service.NewMentorSessionParticipantService(
		repo,
		userService.NewUserService(userRepo),
		sessionService.NewSessionService(sessionRepo, learnerRepo),
	)
	hnd := handler.NewMentorSessionParticipantHandler(srv)

	protectedRoutes := r.Group(apiV1)
	{
		protectedRoutes.Use(middleware.AuthMiddleware())
		protectedRoutes.POST("/:session_id/enroll", hnd.CreateMentorSessionParticipantHandler)
		protectedRoutes.GET("/enrollments", hnd.GetLearnerEnrollmentsHandler)
	}
}
