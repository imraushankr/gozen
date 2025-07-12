package app

import (
	"gozen/src/configs"
	"gozen/src/internal/delivery/http/middleware"
	v1 "gozen/src/internal/delivery/http/v1"
	"gozen/src/internal/pkg/database"
	"gozen/src/internal/repository"
	"gozen/src/internal/usecase"

	"github.com/gin-gonic/gin"
)

func SetupRouter(cfg *configs.Config, db *database.DB) *gin.Engine {
	if cfg.App.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	// Initialize handlers
	healthHandler := v1.NewHealthHandler()

	// Middleware
	router.Use(
		middleware.CORS(&cfg.CORS),
		middleware.RateLimiter(100, 50),
		middleware.Environment(cfg),
	)

	// Initialize reposiroies
	userRepo := repository.NewUserRespository(db)

	// Initialize usecases
	userUsecase := usecase.NewUserUsecase(userRepo, &cfg.JWT)

	// Initialize handlers
	userHandler := v1.NewUserHandler(&userUsecase, &cfg.JWT)

	// API routes
	api := router.Group("/api/v1")
	{
		// Health check endpoint
		api.GET("/health", healthHandler.HealthCheck)

		// Rest of your routes...
		auth := api.Group("/auth")
		{
			auth.POST("/signup", userHandler.Signup)
			auth.POST("/signin", userHandler.Signin)
			auth.POST("/signout", middleware.AuthMiddleware(cfg.JWT.AccessTokenSecret), userHandler.Signout)
			auth.POST("/refresh", userHandler.Signout)
		}

		// Protected user routes 
		// user := api.Group("/users", middleware.AuthMiddleware(cfg.JWT.AccessTokenSecret)) {
		// 	// user.GET("/me", userHandler.GetProfile)
		// }
 	}

	return router
}
