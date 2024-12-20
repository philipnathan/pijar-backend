package notification

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	custom_error "github.com/philipnathan/pijar-backend/internal/notification/custom_error"
	dto "github.com/philipnathan/pijar-backend/internal/notification/dto"
	service "github.com/philipnathan/pijar-backend/internal/notification/service"
	"gorm.io/gorm"
)

type NotificationHandler struct {
	service service.NotificationServiceInterface
}

func NewNotificationHandler(service service.NotificationServiceInterface) *NotificationHandler {
	return &NotificationHandler{
		service: service,
	}
}

// @Summary	Get all user's notifications
// @Schemes
// @Description	Get all user's notifications
// @Tags			Notification
// @Produce		json
// @Security		Bearer
// @Success		200	{object}	GetAllNotificationsResponseDto
// @Failure		401	{object}	Error	"Unauthorized"
// @Failure		500	{object}	Error	"Internal server error"
// @Router			/users/notifications [get]
func (h *NotificationHandler) GetAllNotificationsHandler(c *gin.Context) {
	UserID, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusUnauthorized, custom_error.Error{Error: "Unauthorized"})
		return
	}

	id, ok := UserID.(float64)
	if !ok {
		c.JSON(http.StatusUnauthorized, custom_error.Error{Error: "Unauthorized"})
		return
	}

	notifications, err := h.service.GetAllNotifications(uint(id))

	fmt.Println(notifications)

	emptyResponse := dto.GetAllNotificationsResponseDto{
		Message:       "No notifications found",
		Notifications: []dto.Notif{},
	}

	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, emptyResponse)
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, custom_error.Error{Error: err.Error()})
		return
	}

	if len(notifications) == 0 {
		c.JSON(http.StatusOK, emptyResponse)
		return
	}
	var data []dto.Notif

	for _, notification := range notifications {
		data = append(data, dto.Notif{
			ID:      notification.ID,
			Type:    notification.NotificationType.Type,
			Message: notification.Message,
			IsRead:  notification.IsRead},
		)
	}

	response := dto.GetAllNotificationsResponseDto{
		Message:       "Notifications found",
		Notifications: data,
	}

	c.JSON(http.StatusOK, response)
}

// @Summary	Read notification
// @Schemes
// @Description	Change IsRead Field to True
// @Tags			Notification
// @Produce		json
// @Security		Bearer
// @Param			notificationid	path		string	true	"Notification ID"
// @Success		200				{object}	ReadNotificationResponseDto
// @Failure		401				{object}	Error	"Unauthorized"
// @Failure		500				{object}	Error	"Internal server error"
// @Router			/users/notifications/read/{notificationid} [put]
func (h *NotificationHandler) ReadNotificationHandler(c *gin.Context) {
	UserID, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusUnauthorized, custom_error.Error{Error: "Unauthorized"})
		return
	}

	id, ok := UserID.(float64)
	if !ok {
		c.JSON(http.StatusUnauthorized, custom_error.Error{Error: "Unauthorized"})
		return
	}

	notificationID := c.Param("notificationid")
	if !exist {
		c.JSON(http.StatusBadRequest, custom_error.Error{Error: "Please provide a notification_id"})
		return
	}

	notificationIDuint64, err := strconv.ParseUint(notificationID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, custom_error.Error{Error: err.Error()})
		return
	}

	err = h.service.ReadNotification(uint(id), uint(notificationIDuint64))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, custom_error.Error{Error: err.Error()})
			return
		}
		if err == custom_error.ErrNotificationHasBeenRead {
			c.JSON(http.StatusBadRequest, custom_error.Error{Error: err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, custom_error.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ReadNotificationResponseDto{Message: "Notification with ID " + notificationID + " has been read"})
}
