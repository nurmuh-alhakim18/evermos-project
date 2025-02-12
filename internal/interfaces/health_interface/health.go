package healthinterface

import "github.com/gofiber/fiber/v2"

type HealthRepositoryInterface interface {
}

type HealthServiceInterface interface {
	HealthCheck() (string, error)
}

type HealthHandlerInterface interface {
	HealthCheck(ctx *fiber.Ctx) error
}
