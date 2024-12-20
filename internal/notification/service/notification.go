package notification

import (
	custom_error "github.com/philipnathan/pijar-backend/internal/notification/custom_error"
	model "github.com/philipnathan/pijar-backend/internal/notification/model"
	repo "github.com/philipnathan/pijar-backend/internal/notification/repository"
	"gorm.io/gorm"
)

type NotificationServiceInterface interface {
	GetAllNotifications(userID uint) ([]model.Notification, error)
	ReadNotification(userID uint, notificationID uint) error
}

type NotificationService struct {
	repo repo.NotificationRepositoryInterface
}

func NewNotificationService(repo repo.NotificationRepositoryInterface) NotificationServiceInterface {
	return &NotificationService{
		repo: repo,
	}
}

func (s *NotificationService) GetAllNotifications(userID uint) ([]model.Notification, error) {
	notifications, err := s.repo.GetAllNotifications(&userID)
	if err == gorm.ErrRecordNotFound {
		return nil, gorm.ErrRecordNotFound
	}
	if err != nil {
		return nil, err
	}

	return notifications, nil
}

func (s *NotificationService) ReadNotification(userID uint, notificationID uint) error {
	notification, err := s.repo.GetNotificationByUserIDandNotificationID(&userID, &notificationID)
	if err == gorm.ErrRecordNotFound {
		return gorm.ErrRecordNotFound
	}
	if err != nil {
		return err
	}

	if notification.IsRead == true {
		return custom_error.ErrNotificationHasBeenRead
	}

	notification.IsRead = true
	return s.repo.SaveNotification(notification)
}
