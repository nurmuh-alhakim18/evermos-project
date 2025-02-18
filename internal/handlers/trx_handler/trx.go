package trxhandler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/nurmuh-alhakim18/evermos-project/constants"
	"github.com/nurmuh-alhakim18/evermos-project/helpers"
	trxinterface "github.com/nurmuh-alhakim18/evermos-project/internal/interfaces/trx_interface"
	trxmodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/trx_model"
)

type TrxHandler struct {
	TrxService trxinterface.TrxServiceInterface
}

func (h *TrxHandler) CreateTrx(ctx *fiber.Ctx) error {
	var req trxmodel.TrxReq

	userID, ok := ctx.Locals("userID").(int)
	if !ok {
		return helpers.SendResponse(ctx, fiber.StatusUnauthorized, false, constants.FailedPostMessage, constants.InvalidUserIDErr, nil)
	}

	if err := ctx.BodyParser(&req); err != nil {
		return helpers.SendResponse(ctx, fiber.StatusBadRequest, false, constants.FailedPostMessage, err.Error(), nil)
	}

	trxID, err := h.TrxService.CreateTrx(ctx.Context(), userID, req)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusInternalServerError, false, constants.FailedPostMessage, err.Error(), nil)
	}

	return helpers.SendResponse(ctx, fiber.StatusOK, true, constants.SucceedPostMessage, nil, trxID)
}

func (h *TrxHandler) GetTrxByUserID(ctx *fiber.Ctx) error {
	userID, ok := ctx.Locals("userID").(int)
	if !ok {
		return helpers.SendResponse(ctx, fiber.StatusUnauthorized, false, constants.FailedPostMessage, constants.InvalidUserIDErr, nil)
	}

	search := ctx.Query("search", "")
	pageString := ctx.Query("page", "1")
	limitString := ctx.Query("limit", "10")

	page, err := strconv.Atoi(pageString)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusBadRequest, false, constants.FailedGetMessage, err.Error(), nil)
	}

	limit, err := strconv.Atoi(limitString)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusBadRequest, false, constants.FailedGetMessage, err.Error(), nil)
	}

	trxs, err := h.TrxService.GetTrxByUserID(ctx.Context(), userID, limit, page, search)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusInternalServerError, false, constants.FailedGetMessage, err.Error(), nil)
	}

	return helpers.SendResponse(ctx, fiber.StatusOK, true, constants.SucceedGetMessage, nil, trxs)
}

func (h *TrxHandler) GetTrxByID(ctx *fiber.Ctx) error {
	userID, ok := ctx.Locals("userID").(int)
	if !ok {
		return helpers.SendResponse(ctx, fiber.StatusUnauthorized, false, constants.FailedGetMessage, constants.InvalidUserIDErr, nil)
	}

	trxIDString := ctx.Params("id")
	trxID, err := strconv.Atoi(trxIDString)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusBadRequest, false, constants.FailedGetMessage, err.Error(), nil)
	}

	trx, err := h.TrxService.GetTrxByID(ctx.Context(), trxID)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusNotFound, false, constants.FailedGetMessage, err.Error(), nil)
	}

	if trx.IdUser != userID {
		return helpers.SendResponse(ctx, fiber.StatusForbidden, false, constants.FailedDeleteMessage, constants.NotAuthorizedErr, nil)
	}

	return helpers.SendResponse(ctx, fiber.StatusOK, true, constants.SucceedGetMessage, nil, trx)
}
