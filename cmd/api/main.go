package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/philipnathan/pijar-backend/database"
	categoryRoute "github.com/philipnathan/pijar-backend/internal/category/route"
	learnerRoute "github.com/philipnathan/pijar-backend/internal/learner/route"
	mentor "github.com/philipnathan/pijar-backend/internal/mentor/route"
	mentorSessionParticipant "github.com/philipnathan/pijar-backend/internal/mentor_session_participant/route"
	notification "github.com/philipnathan/pijar-backend/internal/notification/route"
	searchRoute "github.com/philipnathan/pijar-backend/internal/search/route"
	seed "github.com/philipnathan/pijar-backend/internal/seed"
	sessionRoute "github.com/philipnathan/pijar-backend/internal/session/route"
	userRoute "github.com/philipnathan/pijar-backend/internal/user/route"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"

	_ "github.com/philipnathan/pijar-backend/docs"

	"github.com/gin-contrib/cors"
)

//	@title			Pijar API
//	@version		1.0
//	@description	This is a Pijar API

//	@host		localhost:8080
//	@BasePath	/api/v1

// @securityDefinitions.apikey	Bearer
// @in							header
// @name						Authorization
// @description				"Type 'Bearer TOKEN' to correctly set the API Key"
func main() {
	db, err := database.ConnectToDatabase()

	if err != nil {
		fmt.Println("Failed to connect to database:", err)
	}

	fmt.Println("Connected to database!")

	database.MigrateDatabase(db)

	r := gin.Default()

	// docs.SwaggerInfo.Title = "Swagger Example API"
	// docs.SwaggerInfo.Version = "1.0"
	// docs.SwaggerInfo.Host = "108.136.220.233"
	// docs.SwaggerInfo.BasePath = "/api/v1"
	// docs.SwaggerInfo.Schemes = []string{"http"}

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
	}))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	seedDatabase(db)

	userRoute.UserRoute(r, db)
	categoryRoute.CategoryRoute(r, db)
	learnerRoute.LearnerRoute(r, db)
	learnerRoute.LearnerBioRoute(r, db)
	mentor.MentorBioRoute(r, db)
	notification.NotificationRoute(r, db)
	sessionRoute.SessionRoute(r, db)
	searchRoute.SearchRoute(r, db)
	mentorSessionParticipant.MentorSessionParticipantRoute(r, db)

	if err := r.Run(":8080"); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}

func seedDatabase(db *gorm.DB) error {
	seeds := []func(db *gorm.DB) error{
		seed.SeedUser,
		seed.SeedCategory,
		seed.SeedSubCategory,
		seed.SeedMentorBio,
		seed.SeedMentorExperience,
		seed.SeedMentorExpertise,
		seed.SeedNotificationType,
		seed.SeedNotification,
		seed.SeedSession,
		seed.LearnerBio,
		seed.SeedLearnerInterests,
		seed.MentorSessionParticipant,
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
