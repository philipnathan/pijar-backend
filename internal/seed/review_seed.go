package seed

import (
	model "github.com/philipnathan/pijar-backend/internal/session_review/model"
	"gorm.io/gorm"
)

func Review(db *gorm.DB) error {
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	reviews := []model.SessionReview{
		{
			UserID:    1,
			SessionID: 5,
			Review:    nil,
			Rating:    5,
		},
		{
			UserID:    2,
			SessionID: 1,
			Review:    StrPtr("Kelasnya bagus & materi yang diberikan detail. Saya senang sekali belajar dari mentor ini. Terima kasih"),
			Rating:    5,
		},
		{
			UserID:    4,
			SessionID: 4,
			Review:    StrPtr("Penjelasan kurang baik, tetapi materi cukup lengkap."),
			Rating:    3,
		},
		{
			UserID:    6,
			SessionID: 4,
			Review:    nil,
			Rating:    5,
		},
		{
			UserID:    8,
			SessionID: 4,
			Review:    StrPtr("Kelasnya bagus & materi yang diberikan detail. Saya senang sekali belajar dari mentor ini. Terima kasih"),
			Rating:    4,
		},
		{
			UserID:    2,
			SessionID: 7,
			Review:    StrPtr("Kelasnya sangat menarik, materi yang diberikan sangat lengkap, materi lengkap."),
			Rating:    5,
		},
		{
			UserID:    2,
			SessionID: 8,
			Review:    StrPtr("Kelasnya bagus & materi yang diberikan detail. Saya senang sekali belajar dari mentor ini. Terima kasih"),
			Rating:    5,
		},
		{
			UserID:    2,
			SessionID: 9,
			Review:    StrPtr("Kelasnya bagus & materi yang diberikan detail. Saya senang sekali belajar dari mentor ini. Terima kasih"),
			Rating:    5,
		},
		{
			UserID:    2,
			SessionID: 10,
			Review:    StrPtr("Kelasnya bagus & materi yang diberikan detail. Saya senang sekali belajar dari mentor ini. Terima kasih"),
			Rating:    5,
		},
	}

	var count int64
	if err := tx.Model(&model.SessionReview{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	for _, review := range reviews {
		if err := tx.Create(&review).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func StrPtr(s string) *string {
	return &s
}
