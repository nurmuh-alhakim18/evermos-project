package logprodukrepository

import (
	"context"

	logprodukmodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/log_produk_model"
	"gorm.io/gorm"
)

type LogProdukRepository struct {
	DB *gorm.DB
}

func (r *LogProdukRepository) CreateLogProduk(ctx context.Context, logProduk *logprodukmodel.LogProduk) (*int, error) {
	err := r.DB.WithContext(ctx).Create(logProduk).Error
	if err != nil {
		return nil, err
	}

	return &logProduk.ID, nil
}

func (r *LogProdukRepository) GetLogProdukByID(ctx context.Context, logProdukID int) (*logprodukmodel.LogProduk, error) {
	var logProduk logprodukmodel.LogProduk
	err := r.DB.WithContext(ctx).First(&logProduk, logProdukID).Error
	if err != nil {
		return nil, err
	}

	return &logProduk, nil
}
