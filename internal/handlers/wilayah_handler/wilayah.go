package wilayahhandler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/nurmuh-alhakim18/evermos-project/helpers"
	wilayahinterface "github.com/nurmuh-alhakim18/evermos-project/internal/interfaces/wilayah_interface"
)

type WilayahHandler struct {
	WilayahService wilayahinterface.WilayahServiceInterface
}

func (h *WilayahHandler) GetProvinces(ctx *fiber.Ctx) error {
	search := ctx.Query("search", "")
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

	provinces, err := h.WilayahService.GetProvincies(ctx.Context(), search, page, limit)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusInternalServerError, false, "Failed to GET data", err.Error(), nil)
	}

	return helpers.SendResponse(ctx, fiber.StatusOK, true, "Succeed to GET data", nil, provinces)
}

func (h *WilayahHandler) GetCities(ctx *fiber.Ctx) error {
	provID := ctx.Params("prov_id")

	cities, err := h.WilayahService.GetCities(ctx.Context(), provID)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusInternalServerError, false, "Failed to GET data", err.Error(), nil)
	}

	return helpers.SendResponse(ctx, fiber.StatusOK, true, "Succeed to GET data", nil, cities)
}

func (h *WilayahHandler) GetProvinceDetail(ctx *fiber.Ctx) error {
	provID := ctx.Params("prov_id")

	province, err := h.WilayahService.GetProvinceDetail(ctx.Context(), provID)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusInternalServerError, false, "Failed to GET data", err.Error(), nil)
	}

	return helpers.SendResponse(ctx, fiber.StatusOK, true, "Succeed to GET data", nil, province)
}

func (h *WilayahHandler) GetCityDetail(ctx *fiber.Ctx) error {
	cityID := ctx.Params("city_id")

	city, err := h.WilayahService.GetCityDetail(ctx.Context(), cityID)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusInternalServerError, false, "Failed to GET data", err.Error(), nil)
	}

	return helpers.SendResponse(ctx, fiber.StatusOK, true, "Succeed to GET data", nil, city)
}
