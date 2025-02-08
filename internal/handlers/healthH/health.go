package healthH

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nurmuh-alhakim18/evermos-project/helpers"
	"github.com/nurmuh-alhakim18/evermos-project/internal/interfaces/healthI"
)

type HealthHandler struct {
	Service healthI.HealthServiceInterface
}

func (h *HealthHandler) HealthCheck(ctx *fiber.Ctx) error {
	msg, err := h.Service.HealthCheck()
	if err != nil {
		helpers.SendResponse(ctx, fiber.StatusInternalServerError, false, "Failed to check service", err.Error(), nil)
	}

	return helpers.SendResponse(ctx, fiber.StatusOK, true, msg, nil, nil)
}
