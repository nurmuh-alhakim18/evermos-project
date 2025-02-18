package trxinterface

import (
	"context"

	"github.com/gofiber/fiber/v2"
	trxmodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/trx_model"
)

type TrxRepostoryInterface interface {
	CreateTrx(ctx context.Context, trx *trxmodel.Trx) (*int, error)
	GetTrxByUserID(ctx context.Context, userID, limit, offset int) ([]trxmodel.Trx, error)
	GetTrxByID(ctx context.Context, trxID int) (*trxmodel.Trx, error)
}

type TrxServiceInterface interface {
	CreateTrx(ctx context.Context, userID int, req trxmodel.TrxReq) (*int, error)
	GetTrxByUserID(ctx context.Context, userID, limit, page int, search string) ([]trxmodel.Trx, error)
	GetTrxByID(ctx context.Context, trxID int) (*trxmodel.Trx, error)
}

type TrxHandlerInterface interface {
	CreateTrx(ctx *fiber.Ctx) error
	GetTrxByUserID(ctx *fiber.Ctx) error
	GetTrxByID(ctx *fiber.Ctx) error
}
