package alamathandler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/nurmuh-alhakim18/evermos-project/helpers"
	alamatinterface "github.com/nurmuh-alhakim18/evermos-project/internal/interfaces/alamat_interface"
	alamatmodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/alamat_model"
)

type AlamatHandler struct {
	AlamatService alamatinterface.AlamatServiceInterface
}

func (h *AlamatHandler) CreateAlamat(ctx *fiber.Ctx) error {
	var req alamatmodel.Alamat

	userID, ok := ctx.Locals("userID").(int)
	if !ok {
		return helpers.SendResponse(ctx, fiber.StatusUnauthorized, false, "Failed to POST data", "Invalid user ID", nil)
	}

	if err := ctx.BodyParser(&req); err != nil {
		return helpers.SendResponse(ctx, fiber.StatusBadRequest, false, "Failed to POST data", err.Error(), nil)
	}

	if err := req.Validate(); err != nil {
		return helpers.SendResponse(ctx, fiber.StatusBadRequest, false, "Failed to POST data", err.Error(), nil)
	}

	newAlamat := req
	newAlamat.IdUser = userID

	err := h.AlamatService.CreateAlamat(ctx.Context(), newAlamat)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusInternalServerError, false, "Failed to POST data", err.Error(), nil)
	}

	return helpers.SendResponse(ctx, fiber.StatusOK, true, "Succeed to POST data", nil, "Create Alamat Succeed")
}

func (h *AlamatHandler) GetAlamat(ctx *fiber.Ctx) error {
	userID, ok := ctx.Locals("userID").(int)
	if !ok {
		return helpers.SendResponse(ctx, fiber.StatusUnauthorized, false, "Failed to GET data", "Invalid user ID", nil)
	}

	alamats, err := h.AlamatService.GetAlamat(ctx.Context(), userID)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusInternalServerError, false, "Failed to GET data", err.Error(), nil)
	}

	return helpers.SendResponse(ctx, fiber.StatusOK, true, "Succeed to GET data", nil, alamats)
}

func (h *AlamatHandler) GetAlamatByID(ctx *fiber.Ctx) error {
	userID, ok := ctx.Locals("userID").(int)
	if !ok {
		return helpers.SendResponse(ctx, fiber.StatusUnauthorized, false, "Failed to GET data", "Invalid user ID", nil)
	}

	alamatIDString := ctx.Params("id")
	alamatID, err := strconv.Atoi(alamatIDString)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusBadRequest, false, "Failed to GET data", err.Error(), nil)
	}

	alamat, err := h.AlamatService.GetAlamatByID(ctx.Context(), alamatID)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusInternalServerError, false, "Failed to GET data", err.Error(), nil)
	}

	if alamat.IdUser != userID {
		return helpers.SendResponse(ctx, fiber.StatusForbidden, false, "Failed to GET data", "You are not authorized to access this resource", nil)
	}

	return helpers.SendResponse(ctx, fiber.StatusOK, true, "Succeed to GET data", nil, alamat)
}

func (h *AlamatHandler) UpdateAlamat(ctx *fiber.Ctx) error {
	var req alamatmodel.UpdateAlamat

	userID, ok := ctx.Locals("userID").(int)
	if !ok {
		return helpers.SendResponse(ctx, fiber.StatusUnauthorized, false, "Failed to PUT data", "Invalid user ID", nil)
	}

	alamatIDString := ctx.Params("id")
	alamatID, err := strconv.Atoi(alamatIDString)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusBadRequest, false, "Failed to PUT data", err.Error(), nil)
	}

	if err := ctx.BodyParser(&req); err != nil {
		return helpers.SendResponse(ctx, fiber.StatusBadRequest, false, "Failed to PUT data", err.Error(), nil)
	}

	alamatDB, err := h.AlamatService.GetAlamatByID(ctx.Context(), alamatID)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusInternalServerError, false, "Failed to PUT data", err.Error(), nil)
	}

	if alamatDB.IdUser != userID {
		return helpers.SendResponse(ctx, fiber.StatusForbidden, false, "Failed to PUT data", "You are not authorized to access this resource", nil)
	}

	alamat, err := h.AlamatService.UpdateAlamat(ctx.Context(), alamatID, req)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusInternalServerError, false, "Failed to PUT data", err.Error(), nil)
	}

	return helpers.SendResponse(ctx, fiber.StatusOK, true, "Succeed to PUT data", nil, alamat)
}

func (h *AlamatHandler) DeleteAlamat(ctx *fiber.Ctx) error {
	userID, ok := ctx.Locals("userID").(int)
	if !ok {
		return helpers.SendResponse(ctx, fiber.StatusUnauthorized, false, "Failed to PUT data", "Invalid user ID", nil)
	}

	alamatIDString := ctx.Params("id")
	alamatID, err := strconv.Atoi(alamatIDString)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusBadRequest, false, "Failed to DELETE data", err.Error(), nil)
	}

	alamatDB, err := h.AlamatService.GetAlamatByID(ctx.Context(), alamatID)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusInternalServerError, false, "Failed to DELETE data", err.Error(), nil)
	}

	if alamatDB.IdUser != userID {
		return helpers.SendResponse(ctx, fiber.StatusForbidden, false, "Failed to DELETE data", "You are not authorized to access this resource", nil)
	}

	err = h.AlamatService.DeleteAlamat(ctx.Context(), alamatID)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusInternalServerError, false, "Failed to DELETE data", err.Error(), nil)
	}

	return helpers.SendResponse(ctx, fiber.StatusOK, true, "Succeed to PUT data", nil, "Delete Alamat Succeed")
}
