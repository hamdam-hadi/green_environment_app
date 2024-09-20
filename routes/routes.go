package routes

import (
	"green_environment_app/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// InitializeRoutes sets up all routes for the application
func InitializeRoutes(e *echo.Echo, userController *controllers.UserController, productController *controllers.ProductController, challengeController *controllers.ChallengeController, rewardController *controllers.RewardController) {

	// Middleware for JWT Authentication
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes for Users
	e.POST("/users/register", userController.RegisterUser)
	e.POST("/users/login", userController.LoginUser)
	e.GET("/users/:id", userController.GetUserByID, middleware.JWTWithConfig(userController.JWTConfig))

	// Routes for Products
	e.GET("/products/:id", productController.GetProductByID)
	e.POST("/products", productController.CreateProduct, middleware.JWTWithConfig(productController.JWTConfig))
	e.PUT("/products/:id", productController.UpdateProduct, middleware.JWTWithConfig(productController.JWTConfig))
	e.DELETE("/products/:id", productController.DeleteProduct, middleware.JWTWithConfig(productController.JWTConfig))

	// Routes for Challenges
	e.GET("/challenges/:id", challengeController.GetChallengeByID)
	e.POST("/challenges", challengeController.CreateChallenge, middleware.JWTWithConfig(challengeController.JWTConfig))
	e.PUT("/challenges/:id", challengeController.UpdateChallenge, middleware.JWTWithConfig(challengeController.JWTConfig))
	e.DELETE("/challenges/:id", challengeController.DeleteChallenge, middleware.JWTWithConfig(challengeController.JWTConfig))

	// Routes for Rewards
	e.GET("/rewards/:id", rewardController.GetRewardByID)
	e.POST("/rewards", rewardController.CreateReward, middleware.JWTWithConfig(rewardController.JWTConfig))
	e.PUT("/rewards/:id", rewardController.UpdateReward, middleware.JWTWithConfig(rewardController.JWTConfig))
	e.DELETE("/rewards/:id", rewardController.DeleteReward, middleware.JWTWithConfig(rewardController.JWTConfig))

	//routes for user by email
	e.GET("/users/email/:email", userController.GetUserByEmail)
}
