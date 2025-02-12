//go:build wireinject
// +build wireinject

package mentor

import (
	"github.com/google/wire"
	learnerProvider "github.com/philipnathan/pijar-backend/internal/learner"
	handler "github.com/philipnathan/pijar-backend/internal/mentor/handler"
	repo "github.com/philipnathan/pijar-backend/internal/mentor/repository"
	service "github.com/philipnathan/pijar-backend/internal/mentor/service"
	"gorm.io/gorm"
)

var MentorProviderSet = wire.NewSet(
	repo.NewMentorRepository,
	service.NewMentorService,
	handler.NewMentorHandler,
)

func InitializedMentor(db *gorm.DB) (handler.MentorHandlerInterface, error) {
	wire.Build(MentorProviderSet, learnerProvider.LearnerProviderSet)
	return nil, nil
}
