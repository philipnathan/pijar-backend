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

var LearnerBioProviderSet = wire.NewSet(
	repo.NewLearnerBioRepository,
	service.NewLearnerBioService,
	handler.NewLearnerBioHandler,
)

func InitializedLearnerBio(db *gorm.DB) (handler.LearnerBioHandlerInterface, error) {
	wire.Build(LearnerBioProviderSet)
	return nil, nil
}
