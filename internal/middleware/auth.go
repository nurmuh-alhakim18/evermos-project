package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nurmuh-alhakim18/evermos-project/helpers"
)

func AuthMiddleware(ctx *fiber.Ctx) error {
	token, err := helpers.GetBearerToken(ctx)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusUnauthorized, false, "Failed to proceed request", err.Error(), nil)
	}

	userID, err := helpers.ValidateJWT(token)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusUnauthorized, false, "Failed to proceed request", err.Error(), nil)
	}

	ctx.Locals("userID", userID)
	return ctx.Next()
}
