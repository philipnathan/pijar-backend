package seed

import (
	"time"

	model "github.com/philipnathan/pijar-backend/internal/mentor/model"
	repo "github.com/philipnathan/pijar-backend/internal/mentor/repository"
	"gorm.io/gorm"
)

func SeedMentorExperience(db *gorm.DB) error {
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	mentors := []model.MentorExperiences{
		{
			UserID:      1,
			Occupation:  "Penyuluh Pertanian",
			CompanyName: "Dinas Pertanian",
			StartDate:   model.CustomTime{Time: time.Date(2012, 1, 1, 0, 0, 0, 0, time.UTC)},
		},
		{
			UserID:      3,
			Occupation:  "Pelatih Kerajinan",
			CompanyName: "Komunitas Kreatif Indonesia",
			StartDate:   model.CustomTime{Time: time.Date(2016, 5, 1, 0, 0, 0, 0, time.UTC)},
		},
		{
			UserID:      5,
			Occupation:  "Pengerajin Anyaman Bambu",
			CompanyName: "Pengerajin Bambu Denpasar",
			StartDate:   model.CustomTime{Time: time.Date(2005, 5, 1, 0, 0, 0, 0, time.UTC)},
		},
		{
			UserID:      7,
			Occupation:  "Business Owner",
			CompanyName: "UD. Minum Yang Manis",
			StartDate:   model.CustomTime{Time: time.Date(2000, 5, 1, 0, 0, 0, 0, time.UTC)},
		},
		{
			UserID:      9,
			Occupation:  "Dosen",
			CompanyName: "Universitas Pasundan",
			StartDate:   model.CustomTime{Time: time.Date(2012, 1, 1, 0, 0, 0, 0, time.UTC)},
		},
		{
			UserID:      9,
			Occupation:  "Konten Kreator",
			CompanyName: "Youtube",
			StartDate:   model.CustomTime{Time: time.Date(2012, 1, 1, 0, 0, 0, 0, time.UTC)},
		},
		{
			UserID:      10,
			Occupation:  "Konsultan Ekspor Impor",
			CompanyName: "PT. Konsultan Ekspor Impor",
			StartDate:   model.CustomTime{Time: time.Date(2003, 1, 1, 0, 0, 0, 0, time.UTC)},
		},
	}

	var count int64
	if err := db.Model(&model.MentorExperiences{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	for _, mentor := range mentors {
		if err := repo.NewMentorExperienceRepository(db).SaveMentorExperience(&mentor); err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
