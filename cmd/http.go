package cmd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nurmuh-alhakim18/evermos-project/helpers"
	"github.com/nurmuh-alhakim18/evermos-project/internal/handlers"
	"github.com/nurmuh-alhakim18/evermos-project/internal/interfaces"
	"github.com/nurmuh-alhakim18/evermos-project/internal/services"
)

func ServeHTTP() {
	app := fiber.New()
	dependency := dependencyInject()

	app.Get("/health", dependency.healthHandler.HealthCheck)

	port := helpers.GetEnv("PORT", "8080")
	app.Listen(":" + port)
}

type Dependency struct {
	healthHandler interfaces.HealthHandlerInterface
}

func dependencyInject() Dependency {
	healthSvc := &services.HealthService{}
	healthHandler := &handlers.HealthHandler{Service: healthSvc}

	return Dependency{
		healthHandler: healthHandler,
	}
}
