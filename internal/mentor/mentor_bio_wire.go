//go:build wireinject
// +build wireinject

package mentor

import (
	"github.com/google/wire"
	handler "github.com/philipnathan/pijar-backend/internal/mentor/handler"
	repo "github.com/philipnathan/pijar-backend/internal/mentor/repository"
	service "github.com/philipnathan/pijar-backend/internal/mentor/service"
	"gorm.io/gorm"
)

var MentorBioProviderSet = wire.NewSet(
	repo.NewMentorBioRepository,
	service.NewMentorBioService,
	handler.NewMentorBioHandler,
)

func InitializedMentorBio(db *gorm.DB) (handler.MentorBioHandlerInterface, error) {
	wire.Build(MentorBioProviderSet)
	return nil, nil
}
