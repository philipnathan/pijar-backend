package seed

import (
	model "github.com/philipnathan/pijar-backend/internal/notification/model"
	repo "github.com/philipnathan/pijar-backend/internal/notification/repository"
	"gorm.io/gorm"
)

func SeedNotification(db *gorm.DB) error {
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	notifications := []model.Notification{
		{
			NotificationTypeID: 1,
			UserID:             2,
			Message:            "Sesi mentor 'Teknik Bertani Modern' akan dimulai pada 22 Desember 2024 pukul 09:00.",
			IsRead:             false,
		},
		{
			NotificationTypeID: 1,
			UserID:             4,
			Message:            "Sesi mentor 'Membuat Keripik Pisang Berkualitas Ekspor' akan dimulai pada 30 Desember 2024 pukul 15:00.",
			IsRead:             false,
		},
	}

	var count int64
	if err := db.Model(&model.Notification{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	for _, notification := range notifications {
		if err := repo.NewNotificationRepository(db).SaveNotification(&notification); err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
