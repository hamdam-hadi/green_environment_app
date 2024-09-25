package main

import (
	"fmt"
	"green_environment_app/controllers"
	"green_environment_app/models"
	"green_environment_app/repositories"
	"green_environment_app/routes"
	"green_environment_app/services"
	"green_environment_app/utils"

	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeDatabase() (*gorm.DB, error) {
	var dsn string = fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.GetConfig("DB_USER"),
		utils.GetConfig("DB_PASSWORD"),
		utils.GetConfig("DB_HOST"),
		utils.GetConfig("DB_PORT"),
		utils.GetConfig("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	// Initialize the database
	db, err := InitializeDatabase()
	if err != nil {
		log.Fatalf("couldn't connect to the database: %v", err)
	}
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalf("couldn't get the database from gorm: %v", err)
		}
		sqlDB.Close()
	}()
	err = db.AutoMigrate(
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

	// Initialize repositories and services
	userRepo := repositories.NewUserRepository(db)
	productRepo := repositories.NewProductRepository(db)
	challengeRepo := repositories.NewChallengeRepository(db)
	rewardRepo := repositories.NewRewardRepository(db)

	userService := services.NewUserService(userRepo)
	productService := services.NewProductService(productRepo)
	challengeService := services.NewChallengeService(challengeRepo)
	rewardService := services.NewRewardService(rewardRepo)

	userController := controllers.NewUserController(userService)
	productController := controllers.NewProductController(productService)
	challengeController := controllers.NewChallengeController(challengeService)
	rewardController := controllers.NewRewardController(rewardService)

	// Initialize Echo
	e := echo.New()

	// Apply JWT middleware globally
	e.POST("/login", login)

	// Login route (no JWT needed for this route)
	//authconfig := middleware.JWTConfig{
	//	SecretKey: utils.GetConfig("SECRET_KEY"),
	//}

	//e.Use(echojwt.WithConfig(authconfig.Init()))

	// Initialize routes
	routes.InitializeRoutes(e, userController, productController, challengeController, rewardController)

	// Start the server
	if err := e.Start(":2024"); err != nil {
		log.Fatalf("could not start the server: %v", err)
	}
}

func login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "admin" && password == "admin" {

		token, err := utils.GenerateJWT(1, utils.JWTOptions{
			ExpiresDuration: 24,
			SecretKey:       utils.GetConfig("SECRET_KEY"),
		})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to generate JWT"})
		}
		return c.JSON(http.StatusOK, map[string]string{"token": token})
	}
	return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid credentials"})
}
