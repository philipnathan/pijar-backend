package seed

import (
	model "github.com/philipnathan/pijar-backend/internal/notification/model"
	"gorm.io/gorm"
)

func SeedNotificationType(db *gorm.DB) error {
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	notificationTypes := []model.NotificationType{
		{
			Type: "schedule_change",
		},
		{
			Type: "next_class",
		},
	}

	var count int64
	if err := db.Model(&model.NotificationType{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	for _, notificationType := range notificationTypes {
		if err := db.Create(&notificationType).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
