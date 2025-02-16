package kategorimodel

import "time"

type Kategori struct {
	ID           int       `json:"id" gorm:"primaryKey"`
	NamaKategori string    `json:"nama_category" gorm:"type:varchar(255)"`
	UpdatedAt    time.Time `json:"-"`
	CreatedAt    time.Time `json:"-"`
}

func (Kategori) TableName() string {
	return "kategoris"
}

type UpdateKategori struct {
	NamaKategori string `json:"nama_category"`
}
