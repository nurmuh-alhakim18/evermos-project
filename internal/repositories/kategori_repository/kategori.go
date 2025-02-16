package kategorirepository

import (
	"context"

	kategorimodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/kategori_model"
	"gorm.io/gorm"
)

type KategoriRepository struct {
	DB *gorm.DB
}

func (r *KategoriRepository) CreateKategori(ctx context.Context, kategori *kategorimodel.Kategori) (*int, error) {
	err := r.DB.WithContext(ctx).Create(kategori).Error
	if err != nil {
		return nil, err
	}

	return &kategori.ID, nil
}

func (r *KategoriRepository) GetKategoris(ctx context.Context) ([]kategorimodel.Kategori, error) {
	var kategoris []kategorimodel.Kategori
	err := r.DB.WithContext(ctx).Find(&kategoris).Error
	if err != nil {
		return nil, err
	}

	return kategoris, nil
}

func (r *KategoriRepository) GetKategoriByID(ctx context.Context, kategoriID int) (*kategorimodel.Kategori, error) {
	var kategori kategorimodel.Kategori
	err := r.DB.WithContext(ctx).First(&kategori, kategoriID).Error
	if err != nil {
		return nil, err
	}

	return &kategori, nil
}
func (r *KategoriRepository) UpdateKategori(ctx context.Context, kategoriID int, kategoriInput kategorimodel.UpdateKategori) error {
	var kategori kategorimodel.Kategori
	err := r.DB.WithContext(ctx).First(&kategori, kategoriID).Error
	if err != nil {
		return err
	}

	return r.DB.WithContext(ctx).Model(&kategori).Updates(kategoriInput).Error
}

func (r *KategoriRepository) DeleteKategori(ctx context.Context, kategoriID int) error {
	var kategori kategorimodel.Kategori
	return r.DB.WithContext(ctx).Delete(&kategori, kategoriID).Error
}
