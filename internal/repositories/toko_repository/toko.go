package tokorepository

import (
	"context"

	tokomodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/toko_model"
	"gorm.io/gorm"
)

type TokoRepository struct {
	DB *gorm.DB
}

func (r *TokoRepository) CreateToko(ctx context.Context, req *tokomodel.Toko) error {
	return r.DB.WithContext(ctx).Create(req).Error
}

func (r *TokoRepository) GetTokos(ctx context.Context, nama string, offset, limit int) ([]tokomodel.Toko, error) {
	var tokos []tokomodel.Toko

	query := r.DB.WithContext(ctx).Limit(limit).Offset(offset)

	if nama != "" {
		query = query.Where("LOWER(nama_toko) = LOWER(?)", nama)
	}

	err := query.Find(&tokos).Error
	if err != nil {
		return nil, err
	}

	return tokos, nil
}

func (r *TokoRepository) GetTokoByUserID(ctx context.Context, userID int) (*tokomodel.Toko, error) {
	var toko tokomodel.Toko
	err := r.DB.WithContext(ctx).Where("id_user = ?", userID).First(&toko).Error
	if err != nil {
		return nil, err
	}

	return &toko, nil
}

func (r *TokoRepository) GetTokoByID(ctx context.Context, tokoID int) (*tokomodel.Toko, error) {
	var toko tokomodel.Toko

	err := r.DB.WithContext(ctx).First(&toko, tokoID).Error
	if err != nil {
		return nil, err
	}

	return &toko, nil
}

func (r *TokoRepository) UpdateToko(ctx context.Context, tokoID int, tokoInput tokomodel.UpdateToko) error {
	var toko tokomodel.Toko

	err := r.DB.WithContext(ctx).First(&toko, tokoID).Error
	if err != nil {
		return err
	}

	err = r.DB.WithContext(ctx).Model(&toko).Updates(tokoInput).Error
	if err != nil {
		return nil
	}

	return nil
}
