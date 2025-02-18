package detailtrxrepository

import (
	"context"

	detailtrxmodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/detail_trx_model"
	"gorm.io/gorm"
)

type DetailTrxRepository struct {
	DB *gorm.DB
}

func (r *DetailTrxRepository) CreateDetailTrx(ctx context.Context, req *detailtrxmodel.DetailTrx) (*int, error) {
	err := r.DB.WithContext(ctx).Create(req).Error
	if err != nil {
		return nil, err
	}

	return &req.ID, nil
}

func (r *DetailTrxRepository) GetDetailTrxByTrxID(ctx context.Context, trxID int) ([]detailtrxmodel.DetailTrx, error) {
	var detailTrxs []detailtrxmodel.DetailTrx
	err := r.DB.WithContext(ctx).Where("id_trx = ?", trxID).Find(&detailTrxs).Error
	if err != nil {
		return nil, err
	}

	return detailTrxs, err
}
