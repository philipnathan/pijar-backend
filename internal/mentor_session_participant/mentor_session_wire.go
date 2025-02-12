//go:build wireinject
// +build wireinject

package mentor_session_participant

import (
	"github.com/google/wire"
	learnerInit "github.com/philipnathan/pijar-backend/internal/learner"
	handler "github.com/philipnathan/pijar-backend/internal/mentor_session_participant/handler"
	repo "github.com/philipnathan/pijar-backend/internal/mentor_session_participant/repository"
	service "github.com/philipnathan/pijar-backend/internal/mentor_session_participant/service"
	sessionInit "github.com/philipnathan/pijar-backend/internal/session"
	userInit "github.com/philipnathan/pijar-backend/internal/user"
	"gorm.io/gorm"
)

var MentorSessionParticipantProviderSet = wire.NewSet(
	repo.NewMentorSessionParticipantRepository,
	service.NewMentorSessionParticipantService,
	handler.NewMentorSessionParticipantHandler,
)

func InitializedMentorSessionParticipant(db *gorm.DB) (handler.MentorSessionParticipantHandlerInterface, error) {
	wire.Build(MentorSessionParticipantProviderSet, sessionInit.SessionProviderSet, learnerInit.LearnerProviderSet, userInit.UserProviderSet)
	return nil, nil
}
