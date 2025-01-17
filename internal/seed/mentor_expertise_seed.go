package seed

import (
	model "github.com/philipnathan/pijar-backend/internal/mentor/model"
	repo "github.com/philipnathan/pijar-backend/internal/mentor/repository"
	"gorm.io/gorm"
)

func SeedMentorExpertise(db *gorm.DB) error {
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	mentors := []model.MentorExpertises{
		{UserID: 1, Expertise: "Pertanian Organik", Category_id: 1},
		{UserID: 1, Expertise: "Kewirausahaan Desa", Category_id: 2},
		{UserID: 3, Expertise: "Kerajinan Daur Ulang", Category_id: 3},
		{UserID: 3, Expertise: "Pengolahan Hasil Pertanian", Category_id: 4},
		{UserID: 5, Expertise: "Kerajinan dari Anyaman Bambu", Category_id: 3},
		{UserID: 7, Expertise: "Pengolahan Tebu", Category_id: 4},
		{UserID: 9, Expertise: "JavaScript", Category_id: 7},
		{UserID: 9, Expertise: "Laravel", Category_id: 7},
		{UserID: 10, Expertise: "Perdagangan Internasional", Category_id: 8},
		{UserID: 10, Expertise: "Konsultasi Ekspor-Impor", Category_id: 8},
	}

	var count int64
	if err := db.Model(&model.MentorExpertises{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	for _, mentor := range mentors {
		if err := repo.NewMentorExpertiseRepository(db).SaveMentorExpertise(&mentor); err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
