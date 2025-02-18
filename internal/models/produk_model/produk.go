package produkmodel

import (
	"mime/multipart"
	"time"

	fotoprodukmodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/foto_produk_model"
	logprodukmodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/log_produk_model"
)

type Produk struct {
	ID            int                          `json:"id" gorm:"primaryKey"`
	NamaProduk    string                       `json:"nama_produk" form:"nama_produk" gorm:"type:varchar(255)"`
	Slug          string                       `json:"slug" gorm:"type:varchar(255)"`
	HargaReseller string                       `json:"harga_reseller" form:"harga_reseller" gorm:"type:varchar(255)"`
	HargaKonsumen string                       `json:"harga_konsumen" form:"harga_konsumen" gorm:"type:varchar(255)"`
	Stok          int                          `json:"stok" form:"stok"`
	Deskripsi     string                       `json:"deskripsi" form:"stok" gorm:"type:text"`
	IdToko        int                          `json:"-"`
	IdKategori    int                          `json:"-" form:"category_id"`
	FotoProduk    []fotoprodukmodel.FotoProduk `json:"-" gorm:"foreignKey:IdProduk"`
	LogProduk     logprodukmodel.LogProduk     `json:"-" gorm:"foreignKey:IdProduk"`
	UpdatedAt     time.Time                    `json:"-"`
	CreatedAt     time.Time                    `json:"-"`
}

type ProdukReq struct {
	ID            int                     `form:"-"`
	NamaProduk    string                  `form:"nama_produk"`
	HargaReseller string                  `form:"harga_reseller"`
	HargaKonsumen string                  `form:"harga_konsumen"`
	Stok          int                     `form:"stok"`
	Deskripsi     string                  `form:"stok"`
	IdKategori    int                     `form:"category_id"`
	Photos        []*multipart.FileHeader `form:"photos"`
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
	Toko     Toko         `json:"toko"`
	Kategori Kategori     `json:"category"`
	Photos   []FotoProduk `json:"photos"`
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
