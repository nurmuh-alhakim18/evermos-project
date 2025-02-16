package kategoriinterface

import (
	"context"

	"github.com/gofiber/fiber/v2"
	kategorimodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/kategori_model"
)

type KategoriRepositoryInterface interface {
	CreateKategori(ctx context.Context, kategori *kategorimodel.Kategori) (*int, error)
	GetKategoris(ctx context.Context) ([]kategorimodel.Kategori, error)
	GetKategoriByID(ctx context.Context, kategoriID int) (*kategorimodel.Kategori, error)
	UpdateKategori(ctx context.Context, kategoriID int, kategoriInput kategorimodel.UpdateKategori) error
	DeleteKategori(ctx context.Context, kategoriID int) error
}

type KategoriServiceInterface interface {
	CreateKategori(ctx context.Context, req kategorimodel.Kategori) (*int, error)
	GetKategoris(ctx context.Context) ([]kategorimodel.Kategori, error)
	GetKategoriByID(ctx context.Context, kategoriID int) (*kategorimodel.Kategori, error)
	UpdateKategori(ctx context.Context, kategoriID int, req kategorimodel.UpdateKategori) error
	DeleteKategori(ctx context.Context, kategoriID int) error
}

type KategoriHandlerInterface interface {
	CreateKategori(ctx *fiber.Ctx) error
	GetKategoris(ctx *fiber.Ctx) error
	GetKategoriByID(ctx *fiber.Ctx) error
	UpdateKategori(ctx *fiber.Ctx) error
	DeleteKategori(ctx *fiber.Ctx) error
}
