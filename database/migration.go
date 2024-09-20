package database

import (
	"fmt"
	"green_environment_app/models"
	"green_environment_app/utils"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitializeDatabase sets up the connection to the database
func InitializeDatabase() {
	// Get database credentials from the environment
	dbUser := utils.GetConfig("DB_USER")
	dbPassword := utils.GetConfig("DB_PASSWORD")
	dbHost := utils.GetConfig("DB_HOST")
	dbPort := utils.GetConfig("DB_PORT")
	dbName := utils.GetConfig("DB_NAME")

	// Construct the DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	// Connect to the database
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	log.Println("Database connection established successfully")

	// Perform migrations
	runMigrations()
}

// runMigrations migrates all the necessary models to the database
func runMigrations() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Transaction{},
		&models.Challenge{},
		&models.Reward{},
		&models.Leaderboard{},
		&models.Information{},
		&models.Discussion{},
	)
	if err != nil {
		log.Fatalf("Error migrating models: %v", err)
	}

	log.Println("Database migrated successfully")
}
