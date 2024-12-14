package seed

import (
	model "github.com/philipnathan/pijar-backend/internal/session/model"
	repo "github.com/philipnathan/pijar-backend/internal/session/repository"
	"gorm.io/gorm"
)

func SessionSeed(db *gorm.DB) {

	tx := db.Begin()
	if tx.Error != nil {
		return
	}

	sessions := []model.Session{
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

	for _, session := range sessions {
		if err := db.Model(&model.Session{}).Count(&count).Error; err != nil {	
			tx.Rollback()
			return err
		}
	}
	

	if err := tx.Commit().Error; err != nil {
		return
	}

}