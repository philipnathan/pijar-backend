package seed

import (
	model "github.com/philipnathan/pijar-backend/internal/learner/model"
	"gorm.io/gorm"
)

func LearnerBio(db *gorm.DB) error {
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	learners := []model.LearnerBio{
		{
			UserID:      2,
			Bio:         "Ingin belajar cara bertani modern untuk meningkatkan hasil panen.",
			Occupation:  "Petani",
			Institution: "Kelompok Tani Sejahtera",
		},
		{
			UserID:      4,
			Bio:         "Tertarik belajar cara mengolah hasil pertanian menjadi produk bernilai jual.",
			Occupation:  "Petani",
			Institution: "Kelompok Tani Maju",
		},
	}

	var count int64
	if err := db.Model(&model.LearnerBio{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	for _, learner := range learners {
		if err := db.Create(&learner).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
