package detailtrxmodel

import "time"

type DetailTrx struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	IdTrx       int       `json:"-"`
	IdLogProduk int       `json:"-"`
	IdToko      int       `json:"-"`
	Kuantitas   int       `json:"kuantitas"`
	HargaTotal  int       `json:"harga_total"`
	UpdatedAt   time.Time `json:"-"`
	CreatedAt   time.Time `json:"-"`
}

func (DetailTrx) TableName() string {
	return "detail_trxs"
}
