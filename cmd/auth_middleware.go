package cmd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nurmuh-alhakim18/evermos-project/constants"
	"github.com/nurmuh-alhakim18/evermos-project/helpers"
)

func (d *Dependency) AuthMiddleware(ctx *fiber.Ctx) error {
	token, err := helpers.GetBearerToken(ctx)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusUnauthorized, false, constants.MiddlewareFailedMessage, err.Error(), nil)
	}

	userID, err := helpers.ValidateJWT(token)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusUnauthorized, false, constants.MiddlewareFailedMessage, err.Error(), nil)
	}

	ctx.Locals("userID", userID)
	return ctx.Next()
}

func (d *Dependency) AdminMiddleware(ctx *fiber.Ctx) error {
	userID, ok := ctx.Locals("userID").(int)
	if !ok {
		return helpers.SendResponse(ctx, fiber.StatusUnauthorized, false, constants.MiddlewareFailedMessage, constants.InvalidUserIDErr, nil)
	}

	user, err := d.userRepository.GetUserByID(ctx.Context(), userID)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusUnauthorized, false, constants.MiddlewareFailedMessage, err.Error(), nil)
	}

	if !user.IsAdmin {
		return helpers.SendResponse(ctx, fiber.StatusUnauthorized, false, constants.MiddlewareFailedMessage, constants.NotAuthorizedErr, nil)
	}

	return ctx.Next()
}
