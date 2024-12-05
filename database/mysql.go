package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"
)

func ConnectToDatabase() (*sql.DB, error) {

	dbUser := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	maxRetries := 10
	for i := 0; i < maxRetries; i++ {
		db, err := sql.Open("mysql", dsn)

		if err != nil {
			fmt.Println("Error connecting to database:", err)
			time.Sleep(5 * time.Second)
			continue
		}

		if err := db.Ping(); err != nil {
			fmt.Println("Error pinging database:", err)
			time.Sleep(5 * time.Second)
			continue
		}

		return db, nil
	}

	return nil, fmt.Errorf("failed to connect to database after %d retries", maxRetries)
}