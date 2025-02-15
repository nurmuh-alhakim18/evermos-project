package tokohandler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/nurmuh-alhakim18/evermos-project/helpers"
	tokointerface "github.com/nurmuh-alhakim18/evermos-project/internal/interfaces/toko_interface"
	tokomodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/toko_model"
)

type TokoHandler struct {
	TokoService tokointerface.TokoServiceInterface
}

func (h *TokoHandler) GetTokos(ctx *fiber.Ctx) error {
	_, ok := ctx.Locals("userID").(int)
	if !ok {
		return helpers.SendResponse(ctx, fiber.StatusUnauthorized, false, "Failed to GET data", "Invalid user ID", nil)
	}

	nama := ctx.Query("nama", "")
	pageString := ctx.Query("page", "1")
	limitString := ctx.Query("limit", "10")

	page, err := strconv.Atoi(pageString)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusBadRequest, false, "Failed to GET data", err.Error(), nil)
	}

	limit, err := strconv.Atoi(limitString)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusBadRequest, false, "Failed to GET data", err.Error(), nil)
	}

	tokos, err := h.TokoService.GetTokos(ctx.Context(), nama, page, limit)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusInternalServerError, false, "Failed to GET data", err.Error(), nil)
	}

	return helpers.SendResponse(ctx, fiber.StatusOK, true, "Succeed to GET data", nil, tokos)
}

func (h *TokoHandler) GetTokoByUserID(ctx *fiber.Ctx) error {
	userID, ok := ctx.Locals("userID").(int)
	if !ok {
		return helpers.SendResponse(ctx, fiber.StatusUnauthorized, false, "Failed to GET data", "Invalid user ID", nil)
	}

	toko, err := h.TokoService.GetTokoByUserID(ctx.Context(), userID)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusInternalServerError, false, "Failed to GET data", "Invalid user ID", nil)
	}

	return helpers.SendResponse(ctx, fiber.StatusOK, true, "Succeed to GET data", nil, toko)
}

func (h *TokoHandler) GetTokoByID(ctx *fiber.Ctx) error {
	_, ok := ctx.Locals("userID").(int)
	if !ok {
		return helpers.SendResponse(ctx, fiber.StatusUnauthorized, false, "Failed to GET data", "Invalid user ID", nil)
	}

	tokoIDString := ctx.Params("id_toko")
	tokoID, err := strconv.Atoi(tokoIDString)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusBadRequest, false, "Failed to GET data", err.Error(), nil)
	}

	toko, err := h.TokoService.GetTokoByID(ctx.Context(), tokoID)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusNotFound, false, "Failed to GET data", err.Error(), nil)
	}

	return helpers.SendResponse(ctx, fiber.StatusOK, true, "Succeed to GET data", nil, toko)
}

func (h *TokoHandler) UpdateToko(ctx *fiber.Ctx) error {
	_, ok := ctx.Locals("userID").(int)
	if !ok {
		return helpers.SendResponse(ctx, fiber.StatusUnauthorized, false, "Failed to PUT data", "Invalid user ID", nil)
	}

	tokoIDString := ctx.Params("id_toko")
	tokoID, err := strconv.Atoi(tokoIDString)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusBadRequest, false, "Failed to PUT data", err.Error(), nil)
	}

	name := ctx.FormValue("nama_toko")
	photo, _ := ctx.FormFile("photo")

	req := tokomodel.UpdateTokoReq{NamaToko: name, Photo: photo}

	err = h.TokoService.UpdateToko(ctx.Context(), tokoID, req)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusInternalServerError, false, "Failed to PUT data", "Invalid user ID", nil)
	}

	return helpers.SendResponse(ctx, fiber.StatusOK, true, "Succeed to PUT data", nil, "Update toko succeed")
}
