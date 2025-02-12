//go:build wireinject
// +build wireinject

package learner

import (
	"github.com/google/wire"
	handler "github.com/philipnathan/pijar-backend/internal/learner/handler"
	repo "github.com/philipnathan/pijar-backend/internal/learner/repository"
	service "github.com/philipnathan/pijar-backend/internal/learner/service"
	"gorm.io/gorm"
)

var LearnerProviderSet = wire.NewSet(
	repo.NewLearnerRepository,
	service.NewLearnerService,
	handler.NewLearnerHandler,
)

func InitializedLearner(db *gorm.DB) (handler.LearnerHandlerInterface, error) {
	wire.Build(LearnerProviderSet)
	return nil, nil
}
