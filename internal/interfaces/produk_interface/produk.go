package produkinterface

import (
	"context"

	"github.com/gofiber/fiber/v2"
	produkmodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/produk_model"
)

type ProdukRepositoryInterface interface {
	CreateProduk(ctx context.Context, produk *produkmodel.Produk) (*int, error)
	GetProduks(ctx context.Context, queries produkmodel.GetProdukQueriesDB) ([]produkmodel.Produk, error)
	GetProdukByID(ctx context.Context, produkID int) (*produkmodel.Produk, error)
	UpdateProduk(ctx context.Context, produkID int, produkInput produkmodel.UpdateProduk) error
	DeleteProduk(ctx context.Context, produkID int) error
}

type ProdukServiceInterface interface {
	CreateProduk(ctx context.Context, userID int, req produkmodel.Produk) (*int, error)
	GetProduks(ctx context.Context, queries produkmodel.GetProdukQueries) ([]produkmodel.GetProdukResp, error)
	GetProdukByID(ctx context.Context, produkID int) (*produkmodel.GetProdukResp, error)
	UpdateProduk(ctx context.Context, produkID int, req produkmodel.UpdateProduk) error
	DeleteProduk(ctx context.Context, produkID int) error
}

type ProdukHandlerInterface interface {
	CreateProduk(ctx *fiber.Ctx) error
	GetProduks(ctx *fiber.Ctx) error
	GetProdukByID(ctx *fiber.Ctx) error
	UpdateProduk(ctx *fiber.Ctx) error
	DeleteProduk(ctx *fiber.Ctx) error
}
