package trxmodel

import (
	"time"

	detailtrxmodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/detail_trx_model"
)

type Trx struct {
	ID               int                      `json:"id" gorm:"primaryKey"`
	IdUser           int                      `json:"-"`
	AlamatPengiriman int                      `json:"alamat_kirim"`
	HargaTotal       int                      `json:"harga_total"`
	KodeInvoice      string                   `json:"kode_invoice" gorm:"type:varchar(255)"`
	MethodBayar      string                   `json:"method_bayar" gorm:"type:varchar(255)"`
	DetailTrx        detailtrxmodel.DetailTrx `json:"-" gorm:"foreignKey:IdTrx"`
	UpdatedAt        time.Time                `json:"-"`
	CreatedAt        time.Time                `json:"-"`
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

type GetTrxResp struct {
	ID          int             `json:"id"`
	IdUser      int             `json:"-"`
	HargaTotal  int             `json:"harga_total"`
	KodeInvoice string          `json:"kode_invoice"`
	MethodBayar string          `json:"method_bayar"`
	AlamatKirim AlamatKirim     `json:"alamat_kirim"`
	DetailTrx   []DetailTrxResp `json:"detail_trx"`
}

type AlamatKirim struct {
	ID           int    `json:"id"`
	JudulAlamat  string `json:"judul_alamat"`
	NamaPenerima string `json:"nama_penerima"`
	NoTelp       string `json:"no_telp"`
	DetailAlamat string `json:"detail_alamat"`
}

type DetailTrxResp struct {
	Product   Produk `json:"product"`
	Toko      Toko   `json:"toko"`
	Kuantitas int    `json:"kuantitas"`
	Harga     int    `json:"harga"`
}

type Produk struct {
	ID            int          `json:"id"`
	NamaProduk    string       `json:"nama_produk"`
	Slug          string       `json:"slug"`
	HargaReseller string       `json:"harga_reseller"`
	HargaKonsumen string       `json:"harga_konsumen"`
	Deskripsi     string       `json:"deskripsi"`
	Toko          Toko         `json:"toko"`
	Kategori      Kategori     `json:"category"`
	Photos        []FotoProduk `json:"photos"`
}

type Toko struct {
	ID       int    `json:"id"`
	NamaToko string `json:"nama_toko"`
	URLFoto  string `json:"url_foto"`
}

type Kategori struct {
	ID           int    `json:"id"`
	NamaKategori string `json:"nama_category"`
}

type FotoProduk struct {
	ID       int    `json:"id"`
	IdProduk int    `json:"product_id"`
	URL      string `json:"url"`
}
