package alamathandler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/nurmuh-alhakim18/evermos-project/constants"
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
		return helpers.SendResponse(ctx, fiber.StatusUnauthorized, false, constants.FailedPostMessage, constants.InvalidUserIDErr, nil)
	}

	if err := ctx.BodyParser(&req); err != nil {
		return helpers.SendResponse(ctx, fiber.StatusBadRequest, false, constants.FailedPostMessage, err.Error(), nil)
	}

	if err := req.Validate(); err != nil {
		return helpers.SendResponse(ctx, fiber.StatusBadRequest, false, constants.FailedPostMessage, err.Error(), nil)
	}

	newAlamat := req
	newAlamat.IdUser = userID

	err := h.AlamatService.CreateAlamat(ctx.Context(), newAlamat)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusInternalServerError, false, constants.FailedPostMessage, err.Error(), nil)
	}

	return helpers.SendResponse(ctx, fiber.StatusOK, true, constants.SucceedPostMessage, nil, constants.SucceedCreateData)
}

func (h *AlamatHandler) GetAlamat(ctx *fiber.Ctx) error {
	userID, ok := ctx.Locals("userID").(int)
	if !ok {
		return helpers.SendResponse(ctx, fiber.StatusUnauthorized, false, constants.FailedGetMessage, constants.InvalidUserIDErr, nil)
	}

	alamats, err := h.AlamatService.GetAlamat(ctx.Context(), userID)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusInternalServerError, false, constants.FailedGetMessage, err.Error(), nil)
	}

	return helpers.SendResponse(ctx, fiber.StatusOK, true, constants.SucceedGetMessage, nil, alamats)
}

func (h *AlamatHandler) GetAlamatByID(ctx *fiber.Ctx) error {
	userID, ok := ctx.Locals("userID").(int)
	if !ok {
		return helpers.SendResponse(ctx, fiber.StatusUnauthorized, false, constants.FailedGetMessage, constants.InvalidUserIDErr, nil)
	}

	alamatIDString := ctx.Params("id")
	alamatID, err := strconv.Atoi(alamatIDString)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusBadRequest, false, constants.FailedGetMessage, err.Error(), nil)
	}

	alamat, err := h.AlamatService.GetAlamatByID(ctx.Context(), alamatID)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusInternalServerError, false, constants.FailedGetMessage, err.Error(), nil)
	}

	if alamat.IdUser != userID {
		return helpers.SendResponse(ctx, fiber.StatusForbidden, false, constants.FailedGetMessage, constants.NotAuthorizedErr, nil)
	}

	return helpers.SendResponse(ctx, fiber.StatusOK, true, constants.SucceedGetMessage, nil, alamat)
}

func (h *AlamatHandler) UpdateAlamat(ctx *fiber.Ctx) error {
	var req alamatmodel.UpdateAlamat

	userID, ok := ctx.Locals("userID").(int)
	if !ok {
		return helpers.SendResponse(ctx, fiber.StatusUnauthorized, false, constants.FailedUpdateMessage, constants.InvalidUserIDErr, nil)
	}

	alamatIDString := ctx.Params("id")
	alamatID, err := strconv.Atoi(alamatIDString)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusBadRequest, false, constants.FailedUpdateMessage, err.Error(), nil)
	}

	if err := ctx.BodyParser(&req); err != nil {
		return helpers.SendResponse(ctx, fiber.StatusBadRequest, false, constants.FailedUpdateMessage, err.Error(), nil)
	}

	alamatDB, err := h.AlamatService.GetAlamatByID(ctx.Context(), alamatID)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusInternalServerError, false, constants.FailedUpdateMessage, err.Error(), nil)
	}

	if alamatDB.IdUser != userID {
		return helpers.SendResponse(ctx, fiber.StatusForbidden, false, constants.FailedUpdateMessage, constants.NotAuthorizedErr, nil)
	}

	alamat, err := h.AlamatService.UpdateAlamat(ctx.Context(), alamatID, req)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusInternalServerError, false, constants.FailedUpdateMessage, err.Error(), nil)
	}

	return helpers.SendResponse(ctx, fiber.StatusOK, true, constants.SucceedUpdateMessage, nil, alamat)
}

func (h *AlamatHandler) DeleteAlamat(ctx *fiber.Ctx) error {
	userID, ok := ctx.Locals("userID").(int)
	if !ok {
		return helpers.SendResponse(ctx, fiber.StatusUnauthorized, false, constants.FailedDeleteMessage, constants.InvalidUserIDErr, nil)
	}

	alamatIDString := ctx.Params("id")
	alamatID, err := strconv.Atoi(alamatIDString)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusBadRequest, false, constants.FailedDeleteMessage, err.Error(), nil)
	}

	alamatDB, err := h.AlamatService.GetAlamatByID(ctx.Context(), alamatID)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusInternalServerError, false, constants.FailedDeleteMessage, err.Error(), nil)
	}

	if alamatDB.IdUser != userID {
		return helpers.SendResponse(ctx, fiber.StatusForbidden, false, constants.FailedDeleteMessage, constants.NotAuthorizedErr, nil)
	}

	err = h.AlamatService.DeleteAlamat(ctx.Context(), alamatID)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusInternalServerError, false, constants.FailedDeleteMessage, err.Error(), nil)
	}

	return helpers.SendResponse(ctx, fiber.StatusOK, true, constants.SucceedDeleteMessage, nil, constants.SucceedDeleteData)
}
