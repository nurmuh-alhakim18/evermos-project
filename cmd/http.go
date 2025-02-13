package cmd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nurmuh-alhakim18/evermos-project/external"
	"github.com/nurmuh-alhakim18/evermos-project/helpers"
	healthhandler "github.com/nurmuh-alhakim18/evermos-project/internal/handlers/health_handler"
	userhandler "github.com/nurmuh-alhakim18/evermos-project/internal/handlers/user_handler"
	healthinterface "github.com/nurmuh-alhakim18/evermos-project/internal/interfaces/health_interface"
	userinterface "github.com/nurmuh-alhakim18/evermos-project/internal/interfaces/user_interface"
	userrepository "github.com/nurmuh-alhakim18/evermos-project/internal/repositories/user_repository"
	healthservice "github.com/nurmuh-alhakim18/evermos-project/internal/services/health_service"
	userservice "github.com/nurmuh-alhakim18/evermos-project/internal/services/user_service"
	"gorm.io/gorm"
)

func ServeHTTP() {
	app := fiber.New()
	dependency := dependencyInject(helpers.DB)

	api := app.Group("/api/v1")
	api.Get("/health", dependency.AuthMiddleware, dependency.AdminMiddleware, dependency.healthHandler.HealthCheck)

	api.Post("/auth/register", dependency.userHandler.Register)
	api.Post("/auth/login", dependency.userHandler.Login)

	port := helpers.GetEnv("PORT", "8080")
	app.Listen(":" + port)
}

type Dependency struct {
	userRepository userinterface.UserRepositoryInterface

	healthHandler healthinterface.HealthHandlerInterface
	userHandler   userinterface.UserHandlerInterface
}

func dependencyInject(db *gorm.DB) Dependency {
	external := &external.External{}

	healthSvc := &healthservice.HealthService{}
	healthHandler := &healthhandler.HealthHandler{Service: healthSvc}

	userRepo := &userrepository.UserRepository{DB: db}
	userSvc := &userservice.UserService{UserRepository: userRepo, External: external}
	userHandler := &userhandler.UserHandler{UserService: userSvc}

	return Dependency{
		userRepository: userRepo,
		healthHandler:  healthHandler,
		userHandler:    userHandler,
	}
}
