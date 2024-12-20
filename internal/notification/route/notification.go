package notification

import (
	handler "github.com/philipnathan/pijar-backend/internal/notification/handler"
	repository "github.com/philipnathan/pijar-backend/internal/notification/repository"
	service "github.com/philipnathan/pijar-backend/internal/notification/service"
	middleware "github.com/philipnathan/pijar-backend/middleware"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func NotificationRoute(r *gin.Engine, db *gorm.DB) {
	apiV1 := "/api/v1/users/notifications"

	repo := repository.NewNotificationRepository(db)
	services := service.NewNotificationService(repo)
	handler := handler.NewNotificationHandler(services)

	protectedRoutes := r.Group(apiV1)
	{
		protectedRoutes.Use(middleware.AuthMiddleware())
		protectedRoutes.GET("/", handler.GetAllNotificationsHandler)
		protectedRoutes.PUT("/read/:notificationid", handler.ReadNotificationHandler)
	}
}
