package trxrepository

import (
	"context"

	trxmodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/trx_model"
	"gorm.io/gorm"
)

type TrxRepository struct {
	DB *gorm.DB
}

func (r *TrxRepository) CreateTrx(ctx context.Context, trx *trxmodel.Trx) (*int, error) {
	err := r.DB.WithContext(ctx).Create(trx).Error
	if err != nil {
		return nil, err
	}

	return &trx.ID, nil
}

func (r *TrxRepository) GetTrxByUserID(ctx context.Context, userID, limit, offset int) ([]trxmodel.Trx, error) {
	var trxs []trxmodel.Trx
	err := r.DB.WithContext(ctx).Limit(limit).Offset(offset).Where("id_user = ?", userID).Find(&trxs).Error
	if err != nil {
		return nil, err
	}

	return trxs, nil
}

func (r *TrxRepository) GetTrxByID(ctx context.Context, trxID int) (*trxmodel.Trx, error) {
	var trx trxmodel.Trx
	err := r.DB.WithContext(ctx).First(&trx, trxID).Error
	if err != nil {
		return nil, err
	}

	return &trx, nil
}

func (r *TrxRepository) UpdateTotalPrice(ctx context.Context, trxID, totalPrice int) error {
	var trx trxmodel.Trx
	return r.DB.WithContext(ctx).Model(&trx).Where("id = ?", trxID).Update("harga_total", totalPrice).Error
}
