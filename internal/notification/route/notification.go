package notification

import (
	middleware "github.com/philipnathan/pijar-backend/middleware"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	initNotification "github.com/philipnathan/pijar-backend/internal/notification"
)

func NotificationRoute(r *gin.Engine, db *gorm.DB) {
	apiV1 := "/api/v1/users/notifications"

	handler, err := initNotification.InitializedNotification(db)
	if err != nil {
		panic(err)
	}

	protectedRoutes := r.Group(apiV1)
	{
		protectedRoutes.Use(middleware.AuthMiddleware())
		protectedRoutes.GET("", handler.GetAllNotificationsHandler)
		protectedRoutes.PUT("/read/:notificationid", handler.ReadNotificationHandler)
	}
}
