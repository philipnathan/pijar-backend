package database

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConfig struct {
	DB_Username string
	DB_Password string
	DB_Host     string
	DB_Port     string
	DB_Database string
}

func ConnectToDatabase() (*gorm.DB, error) {
	APP_ENV := os.Getenv("APP_ENV")
	config := DBConfig{}

	if APP_ENV == "production" {
		config = DBConfig{
			DB_Username: os.Getenv("AWS_RDS_USERNAME"),
			DB_Password: os.Getenv("AWS_RDS_PASSWORD"),
			DB_Host:     os.Getenv("AWS_RDS_ENDPOINT"),
			DB_Port:     os.Getenv("AWS_RDS_PORT"),
			DB_Database: os.Getenv("AWS_RDS_DBNAME"),
		}
	}

	if APP_ENV == "development" {
		config = DBConfig{
			DB_Username: os.Getenv("DB_USERNAME"),
			DB_Password: os.Getenv("DB_PASSWORD"),
			DB_Host:     os.Getenv("DB_HOST"),
			DB_Port:     os.Getenv("DB_PORT"),
			DB_Database: os.Getenv("DB_DATABASE"),
		}
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DB_Username, config.DB_Password, config.DB_Host, config.DB_Port, config.DB_Database)

	maxRetries := 10
	for i := 0; i < maxRetries; i++ {
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			fmt.Println("Error connecting to database:", err)
			time.Sleep(5 * time.Second)
			continue
		}

		return db, nil
	}

	return nil, fmt.Errorf("failed to connect to database after %d retries", maxRetries)
}
