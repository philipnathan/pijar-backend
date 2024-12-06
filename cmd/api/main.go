package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/philipnathan/pijar-backend/database"
	"github.com/philipnathan/pijar-backend/internal/routes"
)

func main() {
	db, err := database.ConnectToDatabase()

	if err != nil {
		fmt.Println("Failed to connect to database:", err)
	}

	fmt.Println("Connected to database!")

	database.MigrateDatabase(db)

	r := gin.Default()

	routes.UserRoute(r, db)

	if err := r.Run(":8080"); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}