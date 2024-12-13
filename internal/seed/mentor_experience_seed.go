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
			Occupation:  "Software Engineer",
			CompanyName: "Google",
			StartDate:   model.CustomTime{Time: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)},
			EndDate:     model.CustomTime{Time: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
		},
		{
			UserID:      1,
			Occupation:  "Software Engineer",
			CompanyName: "Facebook",
			StartDate:   model.CustomTime{Time: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
			EndDate:     model.CustomTime{Time: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)},
		},
		{
			UserID:      4,
			Occupation:  "Software Engineer",
			CompanyName: "Google",
			StartDate:   model.CustomTime{Time: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)},
			EndDate:     model.CustomTime{Time: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
		},
		{
			UserID:      4,
			Occupation:  "Software Engineer",
			CompanyName: "Facebook",
			StartDate:   model.CustomTime{Time: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
			EndDate:     model.CustomTime{Time: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)},
		},
		{
			UserID:      5,
			Occupation:  "Software Engineer",
			CompanyName: "Google",
			StartDate:   model.CustomTime{Time: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)},
			EndDate:     model.CustomTime{Time: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
		},
		{
			UserID:      5,
			Occupation:  "Software Engineer",
			CompanyName: "Facebook",
			StartDate:   model.CustomTime{Time: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
			EndDate:     model.CustomTime{Time: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)},
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
