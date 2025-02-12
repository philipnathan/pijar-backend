package mentor_session_participant

import (
	"github.com/gin-gonic/gin"
	initMentorSession "github.com/philipnathan/pijar-backend/internal/mentor_session_participant"
	middleware "github.com/philipnathan/pijar-backend/middleware"
	"gorm.io/gorm"
)

func MentorSessionParticipantRoute(r *gin.Engine, db *gorm.DB) {
	apiV1 := "/api/v1/sessions/"

	hnd, err := initMentorSession.InitializedMentorSessionParticipant(db)
	if err != nil {
		panic(err)
	}

	protectedRoutes := r.Group(apiV1)
	{
		protectedRoutes.Use(middleware.AuthMiddleware())
		protectedRoutes.POST("/:session_id/enroll", hnd.CreateMentorSessionParticipantHandler)
		protectedRoutes.GET("/enrollments", hnd.GetLearnerEnrollmentsHandler)
	}
}
