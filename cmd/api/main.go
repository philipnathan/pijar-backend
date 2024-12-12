package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/philipnathan/pijar-backend/database"
	_ "github.com/philipnathan/pijar-backend/docs"
	categoryRoute "github.com/philipnathan/pijar-backend/internal/category/route"
	learnerRoute "github.com/philipnathan/pijar-backend/internal/learner/route"
	mentor "github.com/philipnathan/pijar-backend/internal/mentor/route"
	notification "github.com/philipnathan/pijar-backend/internal/notification/route"
	seed "github.com/philipnathan/pijar-backend/internal/seed"
	userRoute "github.com/philipnathan/pijar-backend/internal/user/route"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
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
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	seedDatabase(db)

	userRoute.UserRoute(r, db)
	categoryRoute.CategoryRoute(r, db)
	learnerRoute.LearnerRoute(r, db)
	learnerRoute.LearnerBioRoute(r, db)
	mentor.MentorBioRoute(r, db)
	notification.NotificationRoute(r, db)

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
