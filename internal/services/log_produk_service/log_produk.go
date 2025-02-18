package logprodukservice

import (
	"context"
	"fmt"

	logprodukinterface "github.com/nurmuh-alhakim18/evermos-project/internal/interfaces/log_produk_interface"
	logprodukmodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/log_produk_model"
)

type LogProdukService struct {
	LogProdukRepository logprodukinterface.LogProdukRepositoryInterface
}

func (s *LogProdukService) CreateLogProduk(ctx context.Context, req logprodukmodel.LogProduk) (*int, error) {
	logProdukID, err := s.LogProdukRepository.CreateLogProduk(ctx, &req)
	if err != nil {
		return nil, fmt.Errorf("failed to create log produk: %v", err)
	}

	return logProdukID, nil
}

func (s *LogProdukService) GetLogProdukByID(ctx context.Context, logProdukID int) (*logprodukmodel.LogProduk, error) {
	logProduk, err := s.LogProdukRepository.GetLogProdukByID(ctx, logProdukID)
	if err != nil {
		return nil, fmt.Errorf("failed to get log produk: %v", err)
	}

	return logProduk, nil
}
