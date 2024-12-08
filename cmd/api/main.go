package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/philipnathan/pijar-backend/database"
	_ "github.com/philipnathan/pijar-backend/docs"
	categoryRoute "github.com/philipnathan/pijar-backend/internal/category/route"
	seed "github.com/philipnathan/pijar-backend/internal/seed"
	userRoute "github.com/philipnathan/pijar-backend/internal/user/route"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Pijar API
//	@version		1.0
//	@description	This is a Pijar API

// @host		localhost:8080
// @BasePath	/api/v1
func main() {
	db, err := database.ConnectToDatabase()

	if err != nil {
		fmt.Println("Failed to connect to database:", err)
	}

	fmt.Println("Connected to database!")

	database.MigrateDatabase(db)

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := seed.SeedDatabase(db); err != nil {
		fmt.Println("Failed to seed database:", err)
	} else {
		fmt.Println("Database seeded successfully")
	}

	userRoute.UserRoute(r, db)
	categoryRoute.CategoryRoute(r, db)

	if err := r.Run(":8080"); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
