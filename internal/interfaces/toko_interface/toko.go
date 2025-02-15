package tokointerface

import (
	"context"

	"github.com/gofiber/fiber/v2"
	tokomodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/toko_model"
)

type TokoRepositoryInterface interface {
	CreateToko(ctx context.Context, req *tokomodel.Toko) error
	GetTokos(ctx context.Context, nama string, offset, limit int) ([]tokomodel.Toko, error)
	GetTokoByUserID(ctx context.Context, userID int) (*tokomodel.Toko, error)
	GetTokoByID(ctx context.Context, tokoID int) (*tokomodel.Toko, error)
	UpdateToko(ctx context.Context, tokoID int, tokoInput tokomodel.UpdateToko) error
}

type TokoServiceInterface interface {
	CreateToko(ctx context.Context, req tokomodel.Toko) error
	GetTokos(ctx context.Context, nama string, page, limit int) ([]tokomodel.Toko, error)
	GetTokoByUserID(ctx context.Context, userID int) (*tokomodel.Toko, error)
	GetTokoByID(ctx context.Context, tokoID int) (*tokomodel.Toko, error)
	UpdateToko(ctx context.Context, tokoID int, req tokomodel.UpdateTokoReq) error
}
type TokoHandlerInterface interface {
	GetTokos(ctx *fiber.Ctx) error
	GetTokoByUserID(ctx *fiber.Ctx) error
	GetTokoByID(ctx *fiber.Ctx) error
	UpdateToko(ctx *fiber.Ctx) error
}
