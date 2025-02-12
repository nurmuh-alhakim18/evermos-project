package cmd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nurmuh-alhakim18/evermos-project/external"
	"github.com/nurmuh-alhakim18/evermos-project/helpers"
	"github.com/nurmuh-alhakim18/evermos-project/internal/handlers/healthH"
	userhandler "github.com/nurmuh-alhakim18/evermos-project/internal/handlers/user_handler"
	"github.com/nurmuh-alhakim18/evermos-project/internal/interfaces/healthI"
	userinterface "github.com/nurmuh-alhakim18/evermos-project/internal/interfaces/user_interface"
	"github.com/nurmuh-alhakim18/evermos-project/internal/middleware"
	userrepository "github.com/nurmuh-alhakim18/evermos-project/internal/repositories/user_repository"
	"github.com/nurmuh-alhakim18/evermos-project/internal/services/healthS"
	userservice "github.com/nurmuh-alhakim18/evermos-project/internal/services/user_service"
	"gorm.io/gorm"
)

func ServeHTTP() {
	app := fiber.New()
	dependency := dependencyInject(helpers.DB)

	api := app.Group("/api/v1")
	api.Get("/health", middleware.AuthMiddleware, dependency.healthHandler.HealthCheck)

	api.Post("/auth/register", dependency.userHandler.Register)
	api.Post("/auth/login", dependency.userHandler.Login)

	port := helpers.GetEnv("PORT", "8080")
	app.Listen(":" + port)
}

type Dependency struct {
	healthHandler healthI.HealthHandlerInterface
	userHandler   userinterface.UserHandlerInterface
}

func dependencyInject(db *gorm.DB) Dependency {
	external := &external.External{}

	healthSvc := &healthS.HealthService{}
	healthHandler := &healthH.HealthHandler{Service: healthSvc}

	userRepo := &userrepository.UserRepository{DB: db}
	userSvc := &userservice.UserService{UserRepository: userRepo, External: external}
	userHandler := &userhandler.UserHandler{UserService: userSvc}

	return Dependency{
		healthHandler: healthHandler,
		userHandler:   userHandler,
	}
}
