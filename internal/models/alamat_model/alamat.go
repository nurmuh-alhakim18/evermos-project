package alamatmodel

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Alamat struct {
	ID           int       `json:"id" gorm:"primaryKey"`
	IdUser       int       `json:"-"`
	JudulAlamat  string    `json:"judul_alamat" gorm:"type:varchar(255)"`
	NamaPenerima string    `json:"nama_penerima" gorm:"type:varchar(255)"`
	NoTelp       string    `json:"no_telp" gorm:"type:varchar(255)"`
	DetailAlamat string    `json:"detail_alamat" gorm:"type:varchar(255)"`
	UpdatedAt    time.Time `json:"-"`
	CreatedAt    time.Time `json:"-"`
}

func (Alamat) TableName() string {
	return "alamats"
}

func (s Alamat) Validate() error {
	v := validator.New()
	return v.Struct(s)
}

type UpdateAlamat struct {
	JudulAlamat  string `json:"judul_alamat"`
	NamaPenerima string `json:"nama_penerima"`
	NoTelp       string `json:"no_telp"`
	DetailAlamat string `json:"detail_alamat"`
}
