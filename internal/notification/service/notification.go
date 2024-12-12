package notification

import (
	model "github.com/philipnathan/pijar-backend/internal/notification/model"
	repo "github.com/philipnathan/pijar-backend/internal/notification/repository"
	"gorm.io/gorm"
)

type NotificationServiceInterface interface {
	GetAllNotifications(userID uint) ([]model.Notification, error)
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
