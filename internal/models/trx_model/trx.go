package trxmodel

import "time"

type Trx struct {
	ID               int       `json:"id" gorm:"primaryKey"`
	IdUser           int       `json:"-"`
	AlamatPengiriman int       `json:"alamat_kirim"`
	HargaTotal       int       `json:"harga_total"`
	KodeInvoice      string    `json:"kode_invoice" gorm:"type:varchar(255)"`
	MethodBayar      string    `json:"method_bayar" gorm:"type:varchar(255)"`
	UpdatedAt        time.Time `json:"-"`
	CreatedAt        time.Time `json:"-"`
}

func (Trx) TableName() string {
	return "trxs"
}

type TrxReq struct {
	MethodBayar      string      `json:"method_bayar"`
	AlamatPengiriman int         `json:"alamat_kirim"`
	DetailTrx        []DetailTrx `json:"detail_trx"`
}

type DetailTrx struct {
	ProductID int `json:"product_id"`
	Kuantitas int `json:"kuantitas"`
}
