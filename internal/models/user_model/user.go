package usermodel

import (
	"time"

	"github.com/go-playground/validator/v10"
	alamatmodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/alamat_model"
	tokomodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/toko_model"
)

type User struct {
	ID           int                  `json:"id" gorm:"primaryKey"`
	Nama         string               `json:"nama" gorm:"type:varchar(255)"`
	KataSandi    string               `json:"kata_sandi,omitempty" gorm:"type:varchar(255)"`
	NoTelp       string               `json:"no_telp" gorm:"type:varchar(255);unique"`
	TanggalLahir string               `json:"tanggal_lahir" gorm:"type:date"`
	JenisKelamin string               `json:"jenis_kelamin" gorm:"type:varchar(255)"`
	Tentang      string               `json:"tentang" gorm:"type:text"`
	Pekerjaan    string               `json:"pekerjaan" gorm:"type:varchar(255)"`
	Email        string               `json:"email" gorm:"type:varchar(255);unique"`
	IdProvinsi   string               `json:"id_provinsi" gorm:"type:varchar(255)"`
	IdKota       string               `json:"id_kota" gorm:"type:varchar(255)"`
	IsAdmin      bool                 `json:"-"`
	Alamats      []alamatmodel.Alamat `json:"-" gorm:"foreignKey:IdUser"`
	Toko         tokomodel.Toko       `json:"-" gorm:"foreignKey:IdUser"`
	UpdatedAt    time.Time            `json:"-"`
	CreatedAt    time.Time            `json:"-"`
}

func (User) TableName() string {
	return "users"
}

func (s User) Validate() error {
	v := validator.New()
	return v.Struct(s)
}

type LoginRequest struct {
	NoTelp    string `json:"no_telp"`
	KataSandi string `json:"kata_sandi"`
}

type LoginResponse struct {
	Nama         string   `json:"nama"`
	NoTelp       string   `json:"no_telp"`
	TanggalLahir string   `json:"tanggal_lahir"`
	JenisKelamin string   `json:"jenis_kelamin"`
	Tentang      string   `json:"tentang"`
	Pekerjaan    string   `json:"pekerjaan"`
	Email        string   `json:"email"`
	Provinsi     Provinsi `json:"provinsi"`
	Kota         Kota     `json:"kota"`
	Token        string   `json:"token"`
}

type Provinsi struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Kota struct {
	ID         string `json:"id"`
	ProvinceId string `json:"province_id"`
	Name       string `json:"name"`
}

type UpdateUser struct {
	Nama         string `json:"nama"`
	KataSandi    string `json:"kata_sandi,omitempty"`
	NoTelp       string `json:"no_telp"`
	TanggalLahir string `json:"tanggal_lahir"`
	JenisKelamin string `json:"jenis_kelamin"`
	Tentang      string `json:"tentang"`
	Pekerjaan    string `json:"pekerjaan"`
	Email        string `json:"email"`
	IdProvinsi   string `json:"id_provinsi"`
	IdKota       string `json:"id_kota"`
}
