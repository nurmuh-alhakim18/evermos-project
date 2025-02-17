package fotoprodukrepository

import (
	"context"

	fotoprodukmodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/foto_produk_model"
	"gorm.io/gorm"
)

type FotoProdukRepository struct {
	DB *gorm.DB
}

func (r *FotoProdukRepository) CreateFotoProduk(ctx context.Context, req *fotoprodukmodel.FotoProduk) error {
	return r.DB.WithContext(ctx).Create(req).Error
}

func (r *FotoProdukRepository) GetFotoProdukByProdukID(ctx context.Context, produkID int) ([]fotoprodukmodel.FotoProduk, error) {
	var fotoProduks []fotoprodukmodel.FotoProduk
	err := r.DB.WithContext(ctx).Where("id_produk = ?", produkID).Find(&fotoProduks).Error
	if err != nil {
		return nil, err
	}

	return fotoProduks, nil
}

func (r *FotoProdukRepository) DeleteFotoProdukByProdukID(ctx context.Context, produkID int) error {
	var fotoProduk fotoprodukmodel.FotoProduk
	return r.DB.WithContext(ctx).Where("id_produk = ?", produkID).Delete(&fotoProduk).Error
}
