//go:build wireinject
// +build wireinject

package user

import (
	"github.com/google/wire"
	hndl "github.com/philipnathan/pijar-backend/internal/user/handler"
	repo "github.com/philipnathan/pijar-backend/internal/user/repository"
	serv "github.com/philipnathan/pijar-backend/internal/user/service"
	"gorm.io/gorm"
)

var UserProviderSet = wire.NewSet(
	repo.NewUserRepository,
	serv.NewUserService,
	hndl.NewUserHandler,
)

func InitializedUser(db *gorm.DB) (hndl.UserHandlerInterface, error) {
	wire.Build(UserProviderSet)

	return nil, nil
}
