package alamatrepository

import (
	"context"

	alamatmodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/alamat_model"
	"gorm.io/gorm"
)

type AlamatRepository struct {
	DB *gorm.DB
}

func (r *AlamatRepository) CreateAlamat(ctx context.Context, alamat *alamatmodel.Alamat) (*int, error) {
	err := r.DB.WithContext(ctx).Create(alamat).Error
	if err != nil {
		return nil, err
	}

	return &alamat.ID, nil
}

func (r *AlamatRepository) GetAlamats(ctx context.Context, userID int) ([]alamatmodel.Alamat, error) {
	var alamats []alamatmodel.Alamat
	err := r.DB.WithContext(ctx).Where("id_user = ?", userID).Find(&alamats).Error
	if err != nil {
		return nil, err
	}

	return alamats, nil
}

func (r *AlamatRepository) GetAlamatByID(ctx context.Context, alamatID int) (*alamatmodel.Alamat, error) {
	var alamat alamatmodel.Alamat
	err := r.DB.WithContext(ctx).First(&alamat, alamatID).Error
	if err != nil {
		return nil, err
	}

	return &alamat, nil
}

func (r *AlamatRepository) UpdateAlamat(ctx context.Context, alamatID int, alamatInput alamatmodel.UpdateAlamat) error {
	var alamat alamatmodel.Alamat
	err := r.DB.WithContext(ctx).First(&alamat, alamatID).Error
	if err != nil {
		return err
	}

	err = r.DB.WithContext(ctx).Model(&alamat).Updates(alamatInput).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *AlamatRepository) DeleteAlamat(ctx context.Context, alamatID int) error {
	var alamat alamatmodel.Alamat
	return r.DB.WithContext(ctx).Delete(&alamat, alamatID).Error
}
