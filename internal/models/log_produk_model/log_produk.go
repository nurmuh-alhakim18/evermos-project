package logprodukmodel

import (
	"time"

	detailtrxmodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/detail_trx_model"
)

type LogProduk struct {
	ID            int                        `json:"id" gorm:"primaryKey"`
	IdProduk      int                        `json:"-"`
	NamaProduk    string                     `json:"nama_produk" gorm:"type:varchar(255)"`
	Slug          string                     `json:"slug" gorm:"type:varchar(255)"`
	HargaReseller string                     `json:"harga_reseller" gorm:"type:varchar(255)"`
	HargaKonsumen string                     `json:"harga_konsumen" gorm:"type:varchar(255)"`
	Stok          int                        `json:"stok"`
	Deskripsi     string                     `json:"deskripsi" gorm:"type:text"`
	IdToko        int                        `json:"-"`
	IdKategori    int                        `json:"-"`
	DetailTrx     []detailtrxmodel.DetailTrx `json:"-" gorm:"foreignKey:IdLogProduk"`
	UpdatedAt     time.Time                  `json:"-"`
	CreatedAt     time.Time                  `json:"-"`
}

func (LogProduk) TableName() string {
	return "log_produks"
}
