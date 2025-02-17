package kategorimodel

import (
	"time"

	produkmodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/produk_model"
)

type Kategori struct {
	ID           int                `json:"id" gorm:"primaryKey"`
	NamaKategori string             `json:"nama_category" gorm:"type:varchar(255)"`
	Produk       produkmodel.Produk `json:"-" gorm:"foreignKey:IdKategori"`
	UpdatedAt    time.Time          `json:"-"`
	CreatedAt    time.Time          `json:"-"`
}

func (Kategori) TableName() string {
	return "kategoris"
}

type UpdateKategori struct {
	NamaKategori string `json:"nama_category"`
}
