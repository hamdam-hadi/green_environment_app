package routes

import (
	"green_environment_app/controllers"
	mid "green_environment_app/middleware"
	"green_environment_app/utils"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// InitializeRoutes sets up all routes for the application
func InitializeRoutes(e *echo.Echo, userController *controllers.UserController, productController *controllers.ProductController, challengeController *controllers.ChallengeController, rewardController *controllers.RewardController) {

	// Middleware for JWT Authentication
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	authconfig := mid.JWTConfig{
		SecretKey: utils.GetConfig("SECRET_KEY"),
	}

	// Routes for Users
	userroutes := e.Group("/users")

	userroutes.POST("/register", userController.RegisterUser)
	userroutes.POST("/login", userController.LoginUser)
	userroutes.GET("/:id", userController.GetUserByID, echojwt.WithConfig(authconfig.Init()), mid.VerifyToken)
	userroutes.GET("/email/:email", userController.GetUserByEmail, echojwt.WithConfig(authconfig.Init()), mid.VerifyToken)

	// Routes for Products
	productroutes := e.Group("/products", echojwt.WithConfig(authconfig.Init()), mid.VerifyToken)
	productroutes.GET("/:id", productController.GetProductByID)
	productroutes.POST("", productController.CreateProduct)
	productroutes.PUT("/:id", productController.UpdateProduct)
	productroutes.DELETE("/:id", productController.DeleteProduct)

	// Routes for Challenges
	challengeroutes := e.Group("/challenges", echojwt.WithConfig(authconfig.Init()), mid.VerifyToken)
	challengeroutes.GET("/:id", challengeController.GetChallengeByID)
	challengeroutes.POST("", challengeController.CreateChallenge)
	challengeroutes.PUT("/:id", challengeController.UpdateChallenge)
	challengeroutes.DELETE("/:id", challengeController.DeleteChallenge)

	// Routes for Rewards
	rewardroutes := e.Group("/rewards", echojwt.WithConfig(authconfig.Init()), mid.VerifyToken)
	rewardroutes.GET("/:id", rewardController.GetRewardByID)
	rewardroutes.POST("", rewardController.CreateReward)
	rewardroutes.PUT("/:id", rewardController.UpdateReward)
	rewardroutes.DELETE("/:id", rewardController.DeleteReward)

	//routes for user by email

}
