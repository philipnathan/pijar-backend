//go:build wireinject
// +build wireinject

package session_review

import (
	"github.com/google/wire"
	learnerProvider "github.com/philipnathan/pijar-backend/internal/learner"
	participantProvider "github.com/philipnathan/pijar-backend/internal/mentor_session_participant"
	sessionProvider "github.com/philipnathan/pijar-backend/internal/session"
	handler "github.com/philipnathan/pijar-backend/internal/session_review/handler"
	repo "github.com/philipnathan/pijar-backend/internal/session_review/repository"
	service "github.com/philipnathan/pijar-backend/internal/session_review/service"
	userProvider "github.com/philipnathan/pijar-backend/internal/user"
	"gorm.io/gorm"
)

var SessionReviewProviderSet = wire.NewSet(
	repo.NewSessionReviewRepository,
	service.NewSessionReviewService,
	handler.NewSessionReviewHandler,
)

func InitializedSessionReview(db *gorm.DB) (handler.SessionReviewHandlerInterface, error) {
	wire.Build(
		SessionReviewProviderSet, learnerProvider.LearnerProviderSet,
		participantProvider.MentorSessionParticipantProviderSet,
		sessionProvider.SessionProviderSet,
		userProvider.UserProviderSet,
	)

	return nil, nil
}
