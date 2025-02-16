package kategorihandler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/nurmuh-alhakim18/evermos-project/constants"
	"github.com/nurmuh-alhakim18/evermos-project/helpers"
	kategoriinterface "github.com/nurmuh-alhakim18/evermos-project/internal/interfaces/kategori_interface"
	kategorimodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/kategori_model"
)

type KategoriHandler struct {
	KategoriService kategoriinterface.KategoriServiceInterface
}

func (h *KategoriHandler) CreateKategori(ctx *fiber.Ctx) error {
	var req kategorimodel.Kategori

	_, ok := ctx.Locals("userID").(int)
	if !ok {
		return helpers.SendResponse(ctx, fiber.StatusUnauthorized, false, constants.FailedPostMessage, constants.InvalidUserIDErr, nil)
	}

	if err := ctx.BodyParser(&req); err != nil {
		return helpers.SendResponse(ctx, fiber.StatusBadRequest, false, constants.FailedPostMessage, err.Error(), nil)
	}

	kategoriID, err := h.KategoriService.CreateKategori(ctx.Context(), req)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusInternalServerError, false, constants.FailedPostMessage, err.Error(), nil)
	}

	return helpers.SendResponse(ctx, fiber.StatusOK, true, constants.SucceedPostMessage, nil, kategoriID)
}

func (h *KategoriHandler) GetKategoris(ctx *fiber.Ctx) error {
	kategoris, err := h.KategoriService.GetKategoris(ctx.Context())
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusInternalServerError, false, constants.FailedGetMessage, err.Error(), nil)
	}

	return helpers.SendResponse(ctx, fiber.StatusOK, true, constants.SucceedGetMessage, nil, kategoris)
}

func (h *KategoriHandler) GetKategoriByID(ctx *fiber.Ctx) error {
	_, ok := ctx.Locals("userID").(int)
	if !ok {
		return helpers.SendResponse(ctx, fiber.StatusUnauthorized, false, constants.FailedGetMessage, constants.InvalidUserIDErr, nil)
	}

	kategoriIDString := ctx.Params("id")
	kategoriID, err := strconv.Atoi(kategoriIDString)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusBadRequest, false, constants.FailedGetMessage, err.Error(), nil)
	}

	kategori, err := h.KategoriService.GetKategoriByID(ctx.Context(), kategoriID)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusInternalServerError, false, constants.FailedGetMessage, err.Error(), nil)
	}

	return helpers.SendResponse(ctx, fiber.StatusOK, true, constants.SucceedGetMessage, nil, kategori)
}

func (h *KategoriHandler) UpdateKategori(ctx *fiber.Ctx) error {
	var req kategorimodel.UpdateKategori

	_, ok := ctx.Locals("userID").(int)
	if !ok {
		return helpers.SendResponse(ctx, fiber.StatusUnauthorized, false, constants.FailedUpdateMessage, constants.InvalidUserIDErr, nil)
	}

	kategoriIDString := ctx.Params("id")
	kategoriID, err := strconv.Atoi(kategoriIDString)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusBadRequest, false, constants.FailedUpdateMessage, err.Error(), nil)
	}

	if err := ctx.BodyParser(&req); err != nil {
		return helpers.SendResponse(ctx, fiber.StatusBadRequest, false, constants.FailedUpdateMessage, err.Error(), nil)
	}

	err = h.KategoriService.UpdateKategori(ctx.Context(), kategoriID, req)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusInternalServerError, false, constants.FailedUpdateMessage, err.Error(), nil)
	}

	return helpers.SendResponse(ctx, fiber.StatusOK, true, constants.SucceedUpdateMessage, nil, constants.SucceedUpdateData)
}

func (h *KategoriHandler) DeleteKategori(ctx *fiber.Ctx) error {
	_, ok := ctx.Locals("userID").(int)
	if !ok {
		return helpers.SendResponse(ctx, fiber.StatusUnauthorized, false, constants.FailedDeleteMessage, constants.InvalidUserIDErr, nil)
	}

	kategoriIDString := ctx.Params("id")
	kategoriID, err := strconv.Atoi(kategoriIDString)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusBadRequest, false, constants.FailedDeleteMessage, err.Error(), nil)
	}

	_, err = h.KategoriService.GetKategoriByID(ctx.Context(), kategoriID)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusInternalServerError, false, constants.FailedDeleteMessage, err.Error(), nil)
	}

	err = h.KategoriService.DeleteKategori(ctx.Context(), kategoriID)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusInternalServerError, false, constants.FailedDeleteMessage, err.Error(), nil)
	}

	return helpers.SendResponse(ctx, fiber.StatusOK, true, constants.SucceedDeleteMessage, nil, constants.SucceedDeleteData)
}
