package database

import (
	"fmt"

	"github.com/philipnathan/pijar-backend/internal/seed"
	"gorm.io/gorm"
)

func SeedDatabase(db *gorm.DB) error {
	seeds := []func(db *gorm.DB) error{
		seed.SeedUser,
		seed.SeedCategory,
		seed.SeedMentorBio,
		seed.SeedMentorExperience,
		seed.SeedMentorExpertise,
		seed.SeedNotificationType,
		seed.SeedNotification,
		seed.SeedSession,
		seed.LearnerBio,
		seed.SeedLearnerInterests,
		seed.MentorSessionParticipant,
		seed.Review,
	}

	for _, seed := range seeds {
		if err := seed(db); err != nil {
			fmt.Println("Failed to seed database:", err)
			return nil
		}
	}

	fmt.Println("Database seeded successfully")
	return nil
}
