package wilayahinterface

import (
	"context"

	"github.com/gofiber/fiber/v2"
	usermodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/user_model"
)

type WilayahRepositoryInterface interface {
	GetProvinces(ctx context.Context, search string, offset, limit int) ([]usermodel.Provinsi, error)
	GetCities(ctx context.Context, provinceID string) ([]usermodel.Kota, error)
	GetProvinceDetail(ctx context.Context, provinceID string) (*usermodel.Provinsi, error)
	GetCityDetail(ctx context.Context, cityID string) (*usermodel.Kota, error)
}

type WilayahServiceInterface interface {
	GetProvincies(ctx context.Context, search string, page, limit int) ([]usermodel.Provinsi, error)
	GetCities(ctx context.Context, provinceID string) ([]usermodel.Kota, error)
	GetProvinceDetail(ctx context.Context, provinceID string) (*usermodel.Provinsi, error)
	GetCityDetail(ctx context.Context, cityID string) (*usermodel.Kota, error)
}

type WilayahHandlerInterface interface {
	GetProvinces(ctx *fiber.Ctx) error
	GetCities(ctx *fiber.Ctx) error
	GetProvinceDetail(ctx *fiber.Ctx) error
	GetCityDetail(ctx *fiber.Ctx) error
}
