package tokomodel

import (
	"mime/multipart"
	"time"

	produkmodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/produk_model"
)

type Toko struct {
	ID        int                  `json:"id" gorm:"primaryKey"`
	IdUser    int                  `json:"-"`
	NamaToko  string               `json:"nama_toko" gorm:"type:varchar(255)"`
	URLFoto   string               `json:"url_foto" gorm:"type:varchar(255)"`
	Produks   []produkmodel.Produk `json:"-" gorm:"foreignKey:IdToko"`
	UpdatedAt time.Time            `json:"-"`
	CreatedAt time.Time            `json:"-"`
}

func (Toko) TableName() string {
	return "tokos"
}

type UpdateTokoReq struct {
	NamaToko string                `form:"nama_toko"`
	Photo    *multipart.FileHeader `form:"photo"`
}

type UpdateToko struct {
	NamaToko string
	URLFoto  string
}
