package main

import (
	"fmt"
	"green_environment_app/controllers"
	"green_environment_app/repositories"
	"green_environment_app/routes"
	"green_environment_app/services"
	"green_environment_app/utils"
	"log"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// connecting to mysql database
func InitializeDatabase() (*gorm.DB, error) {
	var dsn string = fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.GetConfig("DB_USER"),
		utils.GetConfig("DB_PASSWORD"),
		utils.GetConfig("DB_HOST"),
		utils.GetConfig("DB_PORT"),
		utils.GetConfig("DB_NAME"),
	)
	//dsn := "root:@tcp(127.0.0.1:3306)/green_environment_app?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {

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

	// Initializing repositories
	userRepo := repositories.NewUserRepository(db)
	productRepo := repositories.NewProductRepository(db)
	challengeRepo := repositories.NewChallengeRepository(db)
	rewardRepo := repositories.NewRewardRepository(db)

	// Initializing services
	userService := services.NewUserService(userRepo)
	productService := services.NewProductService(productRepo)
	challengeService := services.NewChallengeService(challengeRepo)
	rewardService := services.NewRewardService(rewardRepo)

	// Initializing controllers with services
	userController := controllers.NewUserController(userService)
	productController := controllers.NewProductController(productService)
	challengeController := controllers.NewChallengeController(challengeService)
	rewardController := controllers.NewRewardController(rewardService)

	// Initializing Echo
	e := echo.New()

	// Initializing the routes and starting the server
	routes.InitializeRoutes(e, userController, productController, challengeController, rewardController)
	if err := e.Start(":2024"); err != nil {
		log.Fatalf("can not start the server: %v", err)
	}
}
