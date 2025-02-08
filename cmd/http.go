package cmd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nurmuh-alhakim18/evermos-project/helpers"
	"github.com/nurmuh-alhakim18/evermos-project/internal/handlers"
	"github.com/nurmuh-alhakim18/evermos-project/internal/services"
)

func ServeHTTP() {
	app := fiber.New()

	healthSvc := &services.HealthService{}
	healthHandler := &handlers.HealthHandler{Service: healthSvc}
	app.Get("/health", healthHandler.HealthCheck)

	port := helpers.GetEnv("PORT", "8080")
	app.Listen(":" + port)
}
