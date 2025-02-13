package alamatinterface

import (
	"context"

	"github.com/gofiber/fiber/v2"
	alamatmodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/alamat_model"
)

type AlamatRepositoryInterface interface {
	CreateAlamat(ctx context.Context, alamat *alamatmodel.Alamat) error
	GetAlamat(ctx context.Context, userID int) ([]alamatmodel.Alamat, error)
	GetAlamatByID(ctx context.Context, alamatID int) (*alamatmodel.Alamat, error)
	UpdateAlamat(ctx context.Context, alamatID int, alamatInput alamatmodel.UpdateAlamat) (*alamatmodel.UpdateAlamat, error)
	DeleteAlamat(ctx context.Context, alamatID int) error
}

type AlamatServiceInterface interface {
	CreateAlamat(ctx context.Context, req alamatmodel.Alamat) error
	GetAlamat(ctx context.Context, userID int) ([]alamatmodel.Alamat, error)
	GetAlamatByID(ctx context.Context, alamatID int) (*alamatmodel.Alamat, error)
	UpdateAlamat(ctx context.Context, alamatID int, req alamatmodel.UpdateAlamat) (*alamatmodel.UpdateAlamat, error)
	DeleteAlamat(ctx context.Context, alamatID int) error
}

type AlamatHandlerInterface interface {
	CreateAlamat(ctx *fiber.Ctx) error
	GetAlamat(ctx *fiber.Ctx) error
	GetAlamatByID(ctx *fiber.Ctx) error
	UpdateAlamat(ctx *fiber.Ctx) error
	DeleteAlamat(ctx *fiber.Ctx) error
}
