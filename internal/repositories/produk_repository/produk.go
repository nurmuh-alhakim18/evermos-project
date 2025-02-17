package produkrepository

import (
	"context"

	produkmodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/produk_model"
	"gorm.io/gorm"
)

type ProdukRepository struct {
	DB *gorm.DB
}

func (r *ProdukRepository) CreateProduk(ctx context.Context, produk *produkmodel.Produk) (*int, error) {
	err := r.DB.WithContext(ctx).Create(produk).Error
	if err != nil {
		return nil, err
	}

	return &produk.ID, nil
}

func (r *ProdukRepository) GetProduks(ctx context.Context, queries produkmodel.GetProdukQueriesDB) ([]produkmodel.Produk, error) {
	var produks []produkmodel.Produk
	query := r.DB.WithContext(ctx).Limit(queries.Limit).Offset(queries.Offset)

	if queries.NamaProduk != "" {
		query = query.Where("LOWER(nama_produk) LIKE LOWER(?)", "%"+queries.NamaProduk+"%")
	}

	if queries.TokoId != 0 {
		query = query.Where("id_toko = ?", queries.TokoId)
	}

	if queries.CategoryId != 0 {
		query = query.Where("id_kategori = ?", queries.CategoryId)
	}

	if queries.MinHarga != "" {
		query = query.Where("CAST(harga_reseller AS SIGNED) >= ?", queries.MinHarga)
		query = query.Where("CAST(harga_konsumen AS SIGNED) >= ?", queries.MinHarga)
	}

	if queries.MaxHarga != "" {
		query = query.Where("CAST(harga_reseller AS SIGNED) <= ?", queries.MaxHarga)
		query = query.Where("CAST(harga_konsumen AS SIGNED) <= ?", queries.MaxHarga)
	}

	err := query.Find(&produks).Error
	if err != nil {
		return nil, err
	}

	return produks, nil
}

func (r *ProdukRepository) GetProdukByID(ctx context.Context, produkID int) (*produkmodel.Produk, error) {
	var produk produkmodel.Produk
	err := r.DB.WithContext(ctx).First(&produk, produkID).Error
	if err != nil {
		return nil, err
	}

	return &produk, nil
}

func (r *ProdukRepository) UpdateProduk(ctx context.Context, produkID int, produkInput produkmodel.UpdateProduk) error {
	var produk produkmodel.Produk
	err := r.DB.WithContext(ctx).First(&produk, produkID).Error
	if err != nil {
		return err
	}

	return r.DB.WithContext(ctx).Model(&produk).Updates(produkInput).Error
}

func (r *ProdukRepository) DeleteProduk(ctx context.Context, produkID int) error {
	var produk produkmodel.Produk
	return r.DB.WithContext(ctx).Delete(&produk, produkID).Error
}
