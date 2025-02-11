//go:build wireinject
// +build wireinject

package user

import (
	wire "github.com/google/wire"
	handler "github.com/philipnathan/pijar-backend/internal/user/handler"
	repo "github.com/philipnathan/pijar-backend/internal/user/repository"
	service "github.com/philipnathan/pijar-backend/internal/user/service"
	"gorm.io/gorm"
)

var GoogleProviderSet = wire.NewSet(
	repo.NewGoogleAuthRepo,
	service.NewGoogleAuthService,
	handler.NewGoogleAuthHandler,
)

func InitializedGoogleAuth(db *gorm.DB) (handler.GoogleAuthHandlerInterface, error) {
	wire.Build(GoogleProviderSet)
	return nil, nil
}
