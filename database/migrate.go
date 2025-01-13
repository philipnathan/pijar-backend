package database

import (
	"fmt"

	category "github.com/philipnathan/pijar-backend/internal/category/model"
	learner "github.com/philipnathan/pijar-backend/internal/learner/model"
	mentor "github.com/philipnathan/pijar-backend/internal/mentor/model"
	mentorSessionParticipant "github.com/philipnathan/pijar-backend/internal/mentor_session_participant/model"
	notification "github.com/philipnathan/pijar-backend/internal/notification/model"
	session "github.com/philipnathan/pijar-backend/internal/session/model"
	user "github.com/philipnathan/pijar-backend/internal/user/model"
	"gorm.io/gorm"
)

func MigrateDatabase(db *gorm.DB) {
	err := db.AutoMigrate(
		&user.User{},
		&category.Category{},
		&category.SubCategory{},
		&learner.LearnerBio{},
		&learner.LearnerInterest{},
		&mentor.MentorBiographies{},
		&mentor.MentorExperiences{},
		&mentor.MentorExpertises{},
		&notification.Notification{},
		&notification.NotificationType{},
		&session.MentorSession{},
		&mentorSessionParticipant.MentorSessionParticipant{},
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("Database migrated successfully")
}
