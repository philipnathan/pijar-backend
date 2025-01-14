package session_review

import (
	"github.com/gin-gonic/gin"
	learnerRepo "github.com/philipnathan/pijar-backend/internal/learner/repository"
	participantRepo "github.com/philipnathan/pijar-backend/internal/mentor_session_participant/repository"
	participantService "github.com/philipnathan/pijar-backend/internal/mentor_session_participant/service"
	sessionRepo "github.com/philipnathan/pijar-backend/internal/session/repository"
	sessionService "github.com/philipnathan/pijar-backend/internal/session/service"
	handler "github.com/philipnathan/pijar-backend/internal/session_review/handler"
	repo "github.com/philipnathan/pijar-backend/internal/session_review/repository"
	service "github.com/philipnathan/pijar-backend/internal/session_review/service"
	userRepo "github.com/philipnathan/pijar-backend/internal/user/repository"
	userService "github.com/philipnathan/pijar-backend/internal/user/service"
	middleware "github.com/philipnathan/pijar-backend/middleware"
	"gorm.io/gorm"
)

func SessionReviewRoute(r *gin.Engine, db *gorm.DB) {
	apiV1 := "/api/v1/reviews"

	repo := repo.NewSessionReviewRepository(db)
	userRepo := userRepo.NewUserRepository(db)
	learnerRepo := learnerRepo.NewLearnerRepository(db)
	sessionRepo := sessionRepo.NewSessionRepository(db)
	userService := userService.NewUserService(userRepo)
	sessionService := sessionService.NewSessionService(sessionRepo, learnerRepo)
	participantRepo := participantRepo.NewMentorSessionParticipantRepository(db)
	srv := service.NewSessionReviewService(
		repo,
		userService,
		sessionService,
		participantService.NewMentorSessionParticipantService(
			participantRepo,
			userService,
			sessionService,
		),
	)
	hnd := handler.NewSessionReviewHandler(srv)

	protectedRoutes := r.Group(apiV1)
	{
		protectedRoutes.Use(middleware.AuthMiddleware())
		protectedRoutes.POST("/:session_id", hnd.CreateSessionReviewHandler)
	}
}
