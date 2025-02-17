package produkhandler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/nurmuh-alhakim18/evermos-project/constants"
	"github.com/nurmuh-alhakim18/evermos-project/helpers"
	produkinterface "github.com/nurmuh-alhakim18/evermos-project/internal/interfaces/produk_interface"
	tokointerface "github.com/nurmuh-alhakim18/evermos-project/internal/interfaces/toko_interface"
	produkmodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/produk_model"
)

type ProdukHandler struct {
	ProdukService produkinterface.ProdukServiceInterface
	TokoService   tokointerface.TokoServiceInterface
}

func (h *ProdukHandler) CreateProduk(ctx *fiber.Ctx) error {
	userID, ok := ctx.Locals("userID").(int)
	if !ok {
		return helpers.SendResponse(ctx, fiber.StatusUnauthorized, false, constants.FailedPostMessage, constants.InvalidUserIDErr, nil)
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusBadRequest, false, constants.FailedPostMessage, err.Error(), nil)
	}

	namaProduk := form.Value["nama_produk"][0]
	categoryIDString := form.Value["category_id"][0]
	hargaReseller := form.Value["harga_reseller"][0]
	hargaKonsumen := form.Value["harga_konsumen"][0]
	stokString := form.Value["stok"][0]
	deskripsi := form.Value["deskripsi"][0]
	photos := form.File["photos"]

	categoryID, err := strconv.Atoi(categoryIDString)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusBadRequest, false, constants.FailedPostMessage, err.Error(), nil)
	}

	stok, err := strconv.Atoi(stokString)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusBadRequest, false, constants.FailedPostMessage, err.Error(), nil)
	}

	req := produkmodel.ProdukReq{
		NamaProduk:    namaProduk,
		HargaReseller: hargaReseller,
		HargaKonsumen: hargaKonsumen,
		Stok:          stok,
		Deskripsi:     deskripsi,
		IdKategori:    categoryID,
		Photos:        photos,
	}

	produkID, err := h.ProdukService.CreateProduk(ctx.Context(), userID, req)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusInternalServerError, false, constants.FailedPostMessage, err.Error(), nil)
	}

	return helpers.SendResponse(ctx, fiber.StatusOK, true, constants.SucceedPostMessage, nil, produkID)
}

func (h *ProdukHandler) GetProduks(ctx *fiber.Ctx) error {
	namaProduk := ctx.Query("nama_produk", "")
	limitString := ctx.Query("limit", "10")
	pageString := ctx.Query("page", "1")
	categoryIDString := ctx.Query("category_id", "")
	tokoIDString := ctx.Query("toko_id", "")
	maxHarga := ctx.Query("max_harga", "")
	minHarga := ctx.Query("min_harga", "")

	limit, err := strconv.Atoi(limitString)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusBadRequest, false, constants.FailedGetMessage, err.Error(), nil)
	}

	page, err := strconv.Atoi(pageString)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusBadRequest, false, constants.FailedGetMessage, err.Error(), nil)
	}

	var categoryID int
	if categoryIDString != "" {
		categoryID, err = strconv.Atoi(categoryIDString)
		if err != nil {
			return helpers.SendResponse(ctx, fiber.StatusBadRequest, false, constants.FailedGetMessage, err.Error(), nil)
		}
	}

	var tokoID int
	if tokoIDString != "" {
		tokoID, err = strconv.Atoi(tokoIDString)
		if err != nil {
			return helpers.SendResponse(ctx, fiber.StatusBadRequest, false, constants.FailedGetMessage, err.Error(), nil)
		}
	}

	queries := produkmodel.GetProdukQueries{
		NamaProduk: namaProduk,
		Limit:      limit,
		Page:       page,
		CategoryId: categoryID,
		TokoId:     tokoID,
		MaxHarga:   maxHarga,
		MinHarga:   minHarga,
	}

	produks, err := h.ProdukService.GetProduks(ctx.Context(), queries)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusInternalServerError, false, constants.FailedGetMessage, err.Error(), nil)
	}

	return helpers.SendResponse(ctx, fiber.StatusOK, true, constants.SucceedGetMessage, nil, produks)
}

func (h *ProdukHandler) GetProdukByID(ctx *fiber.Ctx) error {
	produkIDString := ctx.Params("id")
	produkID, err := strconv.Atoi(produkIDString)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusBadRequest, false, constants.FailedGetMessage, err.Error(), nil)
	}

	produk, err := h.ProdukService.GetProdukByID(ctx.Context(), produkID)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusNotFound, false, constants.FailedGetMessage, err.Error(), nil)
	}

	return helpers.SendResponse(ctx, fiber.StatusOK, true, constants.SucceedGetMessage, nil, produk)
}

func (h *ProdukHandler) UpdateProduk(ctx *fiber.Ctx) error {
	userID, ok := ctx.Locals("userID").(int)
	if !ok {
		return helpers.SendResponse(ctx, fiber.StatusUnauthorized, false, constants.FailedUpdateMessage, constants.InvalidUserIDErr, nil)
	}

	produkIDString := ctx.Params("id")
	produkID, err := strconv.Atoi(produkIDString)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusBadRequest, false, constants.FailedUpdateMessage, err.Error(), nil)
	}

	produk, err := h.ProdukService.GetProdukByID(ctx.Context(), produkID)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusNotFound, false, constants.FailedUpdateMessage, err.Error(), nil)
	}

	toko, err := h.TokoService.GetTokoByUserID(ctx.Context(), userID)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusNotFound, false, constants.FailedUpdateMessage, err.Error(), nil)
	}

	if toko.ID != produk.IdToko {
		return helpers.SendResponse(ctx, fiber.StatusForbidden, false, constants.FailedUpdateMessage, constants.NotAuthorizedErr, nil)
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusBadRequest, false, constants.FailedUpdateMessage, err.Error(), nil)
	}

	var namaProduk string
	var categoryIDString string
	var hargaReseller string
	var hargaKonsumen string
	var stokString string
	var deskripsi string
	namaProdukForm := form.Value["nama_produk"]
	if len(namaProdukForm) > 0 {
		namaProduk = namaProdukForm[0]
	}

	categoryIDForm := form.Value["category_id"]
	if len(categoryIDForm) > 0 {
		categoryIDString = categoryIDForm[0]
	}

	hargaResellerForm := form.Value["harga_reseller"]
	if len(hargaResellerForm) > 0 {
		hargaReseller = hargaResellerForm[0]
	}

	hargaKonsumenForm := form.Value["harga_konsumen"]
	if len(hargaKonsumenForm) > 0 {
		hargaKonsumen = hargaKonsumenForm[0]
	}

	stokForm := form.Value["stok"]
	if len(stokForm) > 0 {
		stokString = stokForm[0]
	}

	deskripsiForm := form.Value["deskripsi"]
	if len(deskripsiForm) > 0 {
		deskripsi = deskripsiForm[0]
	}

	photos := form.File["photos"]

	var categoryID int
	var stok int
	if categoryIDString != "" {
		categoryID, err = strconv.Atoi(categoryIDString)
		if err != nil {
			return helpers.SendResponse(ctx, fiber.StatusBadRequest, false, constants.FailedUpdateMessage, err.Error(), nil)
		}
	}

	if stokString != "" {
		stok, err = strconv.Atoi(stokString)
		if err != nil {
			return helpers.SendResponse(ctx, fiber.StatusBadRequest, false, constants.FailedUpdateMessage, err.Error(), nil)
		}
	}

	req := produkmodel.ProdukReq{
		NamaProduk:    namaProduk,
		HargaReseller: hargaReseller,
		HargaKonsumen: hargaKonsumen,
		Stok:          stok,
		Deskripsi:     deskripsi,
		IdKategori:    categoryID,
		Photos:        photos,
	}

	err = h.ProdukService.UpdateProduk(ctx.Context(), produkID, req)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusInternalServerError, false, constants.FailedUpdateMessage, err.Error(), nil)
	}

	return helpers.SendResponse(ctx, fiber.StatusOK, false, constants.SucceedUpdateMessage, nil, constants.SucceedUpdateData)
}

func (h *ProdukHandler) DeleteProduk(ctx *fiber.Ctx) error {
	userID, ok := ctx.Locals("userID").(int)
	if !ok {
		return helpers.SendResponse(ctx, fiber.StatusUnauthorized, false, constants.FailedDeleteMessage, constants.InvalidUserIDErr, nil)
	}

	produkIDString := ctx.Params("id")
	produkID, err := strconv.Atoi(produkIDString)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusBadRequest, false, constants.FailedDeleteMessage, err.Error(), nil)
	}

	produk, err := h.ProdukService.GetProdukByID(ctx.Context(), produkID)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusNotFound, false, constants.FailedDeleteMessage, err.Error(), nil)
	}

	toko, err := h.TokoService.GetTokoByUserID(ctx.Context(), userID)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusNotFound, false, constants.FailedDeleteMessage, err.Error(), nil)
	}

	if toko.ID != produk.IdToko {
		return helpers.SendResponse(ctx, fiber.StatusForbidden, false, constants.FailedDeleteMessage, constants.NotAuthorizedErr, nil)
	}

	err = h.ProdukService.DeleteProduk(ctx.Context(), produkID)
	if err != nil {
		return helpers.SendResponse(ctx, fiber.StatusInternalServerError, false, constants.FailedDeleteMessage, err.Error(), nil)
	}

	return helpers.SendResponse(ctx, fiber.StatusOK, false, constants.SucceedDeleteMessage, nil, constants.SucceedDeleteData)
}
