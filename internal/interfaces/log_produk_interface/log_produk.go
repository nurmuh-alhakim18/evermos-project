package logprodukinterface

import (
	"context"

	logprodukmodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/log_produk_model"
)

type LogProdukRepositoryInterface interface {
	CreateLogProduk(ctx context.Context, logProduk *logprodukmodel.LogProduk) (*int, error)
	GetLogProdukByID(ctx context.Context, logProdukID int) (*logprodukmodel.LogProduk, error)
}

type LogProdukServiceInterface interface {
	CreateLogProduk(ctx context.Context, req logprodukmodel.LogProduk) (*int, error)
	GetLogProdukByID(ctx context.Context, logProdukID int) (*logprodukmodel.LogProduk, error)
}
