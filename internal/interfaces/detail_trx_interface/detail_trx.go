package detailtrxinterface

import (
	"context"

	detailtrxmodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/detail_trx_model"
)

type DetailTrxRepositoryInterface interface {
	CreateDetailTrx(ctx context.Context, req *detailtrxmodel.DetailTrx) (*int, error)
	GetDetailTrxByTrxID(ctx context.Context, trxID int) ([]detailtrxmodel.DetailTrx, error)
}

type DetailTrxServiceInterface interface {
	CreateDetailTrx(ctx context.Context, req detailtrxmodel.DetailTrx) (*int, error)
	GetDetailTrxByTrxID(ctx context.Context, trxID int) ([]detailtrxmodel.DetailTrx, error)
}
