package seed

import (
	"log"
	"time"

	model "github.com/philipnathan/pijar-backend/internal/session/model"
	"gorm.io/gorm"
)

func SeedSession(db *gorm.DB) error {
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	sessions := []model.MentorSession{
		{
			UserID:           1, // Replace with an actual user ID
			CategoryID:       1, // Replace with an actual category ID
			Title:            "Introduction to Python",
			ShortDescription: "Learn the basics of Python programming.",
			Detail:           "This session will cover Python syntax, data types, and basic programming concepts.",
			Schedule:         time.Date(2024, 12, 11, 10, 0, 0, 0, time.UTC),
			EstimateDuration: 120, // Duration in minutes
			ImageURL:         "https://example.com/python-intro.jpg",
			Link:             "https://example.com/python-session-link",
			CreatedAt:        time.Now(),
			UpdatedAt:        time.Now(),
		},
		{
			UserID:           2,
			CategoryID:       2,
			Title:            "Advanced JavaScript Techniques",
			ShortDescription: "Dive deep into advanced JavaScript concepts.",
			Detail:           "This session will cover closures, asynchronous programming, and ES6+ features.",
			Schedule:         time.Date(2024, 12, 15, 14, 0, 0, 0, time.UTC),
			EstimateDuration: 90,
			ImageURL:         "https://example.com/js-advanced.jpg",
			Link:             "https://example.com/js-session-link",
			CreatedAt:        time.Now(),
			UpdatedAt:        time.Now(),
		},
	}

	var count int64
	if err := db.Model(&model.MentorSession{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	for _, session := range sessions {
		if err := tx.Create(&session).Error; err != nil {
			tx.Rollback()
			log.Printf("Error seeding session: %v", err)
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		log.Printf("Error committing transaction: %v", err)
		return err
	}

	return nil
}
