//go:build wireinject
// +build wireinject

package user

import (
	"github.com/google/wire"
	handler "github.com/philipnathan/pijar-backend/internal/user/handler"
	repo "github.com/philipnathan/pijar-backend/internal/user/repository"
	service "github.com/philipnathan/pijar-backend/internal/user/service"
	"gorm.io/gorm"
)

var MentorProviderSet = wire.NewSet(
	repo.NewMentorUserRepository,
	service.NewMentorUserService,
	handler.NewMentorUserHandler,
)

func InitializedMentor(db *gorm.DB) (handler.MentorUserHandlerInterface, error) {
	wire.Build(MentorProviderSet)

	return nil, nil
}
