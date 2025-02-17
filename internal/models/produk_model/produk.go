package produkmodel

import (
	"time"
)

type Produk struct {
	ID            int       `json:"id" gorm:"primaryKey"`
	NamaProduk    string    `json:"nama_produk" form:"nama_produk" gorm:"type:varchar(255)"`
	Slug          string    `json:"slug" gorm:"type:varchar(255)"`
	HargaReseller string    `json:"harga_reseller" form:"harga_reseller" gorm:"type:varchar(255)"`
	HargaKonsumen string    `json:"harga_konsumen" form:"harga_konsumen" gorm:"type:varchar(255)"`
	Stok          int       `json:"stok" form:"stok"`
	Deskripsi     string    `json:"deskripsi" form:"stok" gorm:"type:text"`
	IdToko        int       `json:"-"`
	IdKategori    int       `json:"-" form:"category_id"`
	UpdatedAt     time.Time `json:"-"`
	CreatedAt     time.Time `json:"-"`
}

func (Produk) TableName() string {
	return "produks"
}

type GetProdukQueries struct {
	NamaProduk string `query:"nama_produk"`
	Limit      int    `query:"limit"`
	Page       int    `query:"page"`
	CategoryId int    `query:"category_id"`
	TokoId     int    `query:"toko_id"`
	MaxHarga   string `query:"max_harga"`
	MinHarga   string `query:"min_harga"`
}

type GetProdukQueriesDB struct {
	NamaProduk string
	Limit      int
	Offset     int
	CategoryId int
	TokoId     int
	MaxHarga   string
	MinHarga   string
}

type GetProdukResp struct {
	Produk
	Toko     Toko     `json:"toko"`
	Kategori Kategori `json:"category"`
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

type UpdateProduk struct {
	NamaProduk    string    `form:"nama_produk"`
	Slug          string    `form:"slug"`
	HargaReseller string    `form:"harga_reseller"`
	HargaKonsumen string    `form:"harga_konsumen"`
	Stok          int       `json:"stok" form:"stok"`
	Deskripsi     string    `json:"deskripsi" form:"stok" gorm:"type:text"`
	IdKategori    int       `form:"category_id"`
	UpdatedAt     time.Time `json:"-"`
}
