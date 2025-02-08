package cmd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nurmuh-alhakim18/evermos-project/helpers"
	"github.com/nurmuh-alhakim18/evermos-project/internal/handlers/healthH"
	"github.com/nurmuh-alhakim18/evermos-project/internal/interfaces/healthI"
	"github.com/nurmuh-alhakim18/evermos-project/internal/services/healthS"
)

func ServeHTTP() {
	app := fiber.New()
	dependency := dependencyInject()

	app.Get("/health", dependency.healthHandler.HealthCheck)

	port := helpers.GetEnv("PORT", "8080")
	app.Listen(":" + port)
}

type Dependency struct {
	healthHandler healthI.HealthHandlerInterface
}

func dependencyInject() Dependency {
	healthSvc := &healthS.HealthService{}
	healthHandler := &healthH.HealthHandler{Service: healthSvc}

	return Dependency{
		healthHandler: healthHandler,
	}
}
