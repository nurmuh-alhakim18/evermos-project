package kategorimodel

import (
	"time"

	logprodukmodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/log_produk_model"
	produkmodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/produk_model"
)

type Kategori struct {
	ID           int                      `json:"id" gorm:"primaryKey"`
	NamaKategori string                   `json:"nama_category" gorm:"type:varchar(255)"`
	Produk       produkmodel.Produk       `json:"-" gorm:"foreignKey:IdKategori"`
	LogProduk    logprodukmodel.LogProduk `json:"-" gorm:"foreignKey:IdKategori"`
	UpdatedAt    time.Time                `json:"-"`
	CreatedAt    time.Time                `json:"-"`
}

func (Kategori) TableName() string {
	return "kategoris"
}

type UpdateKategori struct {
	NamaKategori string    `json:"nama_category"`
	UpdatedAt    time.Time `json:"-"`
}
