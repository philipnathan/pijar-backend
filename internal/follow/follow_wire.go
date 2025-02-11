//go:build wireinject
// +build wireinject

package follow

import (
	"github.com/google/wire"
	handler "github.com/philipnathan/pijar-backend/internal/follow/handler"
	repo "github.com/philipnathan/pijar-backend/internal/follow/repository"
	service "github.com/philipnathan/pijar-backend/internal/follow/service"
	userWire "github.com/philipnathan/pijar-backend/internal/user"
	"gorm.io/gorm"
)

var FollowProviderSet = wire.NewSet(
	repo.NewFollowRepository,
	service.NewFollowService,
	handler.NewFollowHandler,
	userWire.UserProviderSet,
)

func InitializedFollow(db *gorm.DB) (handler.FollowHandlerInterface, error) {
	wire.Build(FollowProviderSet)
	return nil, nil
}
