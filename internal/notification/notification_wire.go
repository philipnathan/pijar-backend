//go:build wireinject
// +build wireinject

package notification

import (
	"github.com/google/wire"
	handler "github.com/philipnathan/pijar-backend/internal/notification/handler"
	repo "github.com/philipnathan/pijar-backend/internal/notification/repository"
	service "github.com/philipnathan/pijar-backend/internal/notification/service"
	"gorm.io/gorm"
)

var NotificationProviderSet = wire.NewSet(
	repo.NewNotificationRepository,
	service.NewNotificationService,
	handler.NewNotificationHandler,
)

func InitializedNotification(db *gorm.DB) (handler.NotificationHandlerInterface, error) {
	wire.Build(NotificationProviderSet)
	return nil, nil
}
