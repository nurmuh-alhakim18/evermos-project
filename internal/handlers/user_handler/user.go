package userhandler

import (
	"time"

	"github.com/gofiber/fiber/v2"
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
		return helpers.SendResponse(ctx, fiber.StatusBadRequest, false, "Failed to POST data", err.Error(), nil)
	}

	layout := "02/01/2006"
	date, err := time.Parse(layout, req.TanggalLahir)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusBadRequest, false, "Failed to POST data", err.Error(), nil)
	}

	req.TanggalLahir = date.String()

	if err := req.Validate(); err != nil {
		return helpers.SendResponse(ctx, fiber.StatusBadRequest, false, "Failed to POST data", err.Error(), nil)
	}

	err = h.UserService.Register(ctx.Context(), req)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusInternalServerError, false, "Failed to POST data", err.Error(), nil)
	}

	return helpers.SendResponse(ctx, fiber.StatusOK, true, "Succeed to POST data", nil, "Register Succeed")
}

func (h *UserHandler) Login(ctx *fiber.Ctx) error {
	var req usermodel.LoginRequest

	if err := ctx.BodyParser(&req); err != nil {
		return helpers.SendResponse(ctx, fiber.StatusBadRequest, false, "Failed to POST data", err.Error(), nil)
	}

	resp, err := h.UserService.Login(ctx.Context(), req)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusUnauthorized, false, "Failed to POST data", "No Telp atau kata sandi salah", nil)
	}

	return helpers.SendResponse(ctx, fiber.StatusOK, true, "Succeed to POST data", nil, resp)
}
