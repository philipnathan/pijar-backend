//go:build wireinject
// +build wireinject

package session

import (
	"github.com/google/wire"
	learnerProvider "github.com/philipnathan/pijar-backend/internal/learner"
	handler "github.com/philipnathan/pijar-backend/internal/session/handler"
	repo "github.com/philipnathan/pijar-backend/internal/session/repository"
	service "github.com/philipnathan/pijar-backend/internal/session/service"
	"gorm.io/gorm"
)

var SessionProviderSet = wire.NewSet(
	repo.NewSessionRepository,
	service.NewSessionService,
	handler.NewSessionHandler,
)

func InitializedSession(db *gorm.DB) (handler.SessionHandlerInterface, error) {
	wire.Build(SessionProviderSet, learnerProvider.LearnerProviderSet)
	return nil, nil
}
