package healthhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nurmuh-alhakim18/evermos-project/helpers"
	healthinterface "github.com/nurmuh-alhakim18/evermos-project/internal/interfaces/health_interface"
)

type HealthHandler struct {
	Service healthinterface.HealthServiceInterface
}

func (h *HealthHandler) HealthCheck(ctx *fiber.Ctx) error {
	msg, err := h.Service.HealthCheck()
	if err != nil {
		helpers.SendResponse(ctx, fiber.StatusInternalServerError, false, "Failed to check service", err.Error(), nil)
	}

	return helpers.SendResponse(ctx, fiber.StatusOK, true, msg, nil, nil)
}
