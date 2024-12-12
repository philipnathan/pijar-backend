package notification

import (
	model "github.com/philipnathan/pijar-backend/internal/notification/model"
	"gorm.io/gorm"
)

type NotificationRepositoryInterface interface {
	GetAllNotifications(userID *uint) ([]model.Notification, error)
	SaveNotification(notification *model.Notification) error
}

type NotificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) NotificationRepositoryInterface {
	return &NotificationRepository{
		db: db,
	}
}

func (r *NotificationRepository) GetAllNotifications(userID *uint) ([]model.Notification, error) {
	var notifications []model.Notification
	if err := r.db.Preload("NotificationType").Where("user_id = ?", userID).Find(&notifications).Error; err != nil {
		return nil, err
	}
	return notifications, nil
}

func (r *NotificationRepository) SaveNotification(notification *model.Notification) error {
	return r.db.Save(notification).Error
}
