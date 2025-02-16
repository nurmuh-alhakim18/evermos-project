package userhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nurmuh-alhakim18/evermos-project/constants"
	"github.com/nurmuh-alhakim18/evermos-project/helpers"
	userinterface "github.com/nurmuh-alhakim18/evermos-project/internal/interfaces/user_interface"
	usermodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/user_model"
)

type UserHandler struct {
	UserService userinterface.UserServiceInterface
}

func (h *UserHandler) Register(ctx *fiber.Ctx) error {
	var req usermodel.User

	if err := ctx.BodyParser(&req); err != nil {
		return helpers.SendResponse(ctx, fiber.StatusBadRequest, false, constants.FailedPostMessage, err.Error(), nil)
	}

	err := h.UserService.Register(ctx.Context(), req)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusInternalServerError, false, constants.FailedPostMessage, err.Error(), nil)
	}

	return helpers.SendResponse(ctx, fiber.StatusOK, true, constants.SucceedPostMessage, nil, "Register Succeed")
}

func (h *UserHandler) Login(ctx *fiber.Ctx) error {
	var req usermodel.LoginRequest

	if err := ctx.BodyParser(&req); err != nil {
		return helpers.SendResponse(ctx, fiber.StatusBadRequest, false, constants.FailedPostMessage, err.Error(), nil)
	}

	resp, err := h.UserService.Login(ctx.Context(), req)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusUnauthorized, false, constants.FailedPostMessage, "No Telp atau kata sandi salah", nil)
	}

	return helpers.SendResponse(ctx, fiber.StatusOK, true, constants.SucceedPostMessage, nil, resp)
}

func (h *UserHandler) GetProfile(ctx *fiber.Ctx) error {
	userID, ok := ctx.Locals("userID").(int)
	if !ok {
		return helpers.SendResponse(ctx, fiber.StatusUnauthorized, false, constants.FailedGetMessage, constants.InvalidUserIDErr, nil)
	}

	user, err := h.UserService.GetProfile(ctx.Context(), userID)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusInternalServerError, false, constants.FailedGetMessage, err.Error(), nil)
	}

	resp := user
	resp.KataSandi = ""

	return helpers.SendResponse(ctx, fiber.StatusOK, true, constants.SucceedGetMessage, nil, resp)
}

func (h *UserHandler) UpdateUser(ctx *fiber.Ctx) error {
	var req usermodel.UpdateUser

	userID, ok := ctx.Locals("userID").(int)
	if !ok {
		return helpers.SendResponse(ctx, fiber.StatusUnauthorized, false, constants.FailedUpdateMessage, constants.InvalidUserIDErr, nil)
	}

	if err := ctx.BodyParser(&req); err != nil {
		return helpers.SendResponse(ctx, fiber.StatusBadRequest, false, constants.FailedUpdateMessage, err.Error(), nil)
	}

	err := h.UserService.UpdateUser(ctx.Context(), userID, req)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusInternalServerError, false, constants.FailedUpdateMessage, err.Error(), nil)
	}

	return helpers.SendResponse(ctx, fiber.StatusOK, true, constants.SucceedUpdateMessage, nil, constants.SucceedUpdateData)
}
