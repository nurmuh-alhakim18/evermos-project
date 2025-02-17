package fotoprodukmodel

import "time"

type FotoProduk struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	IdProduk  int       `json:"-"`
	URL       string    `json:"url" gorm:"type:varchar(255)"`
	UpdatedAt time.Time `json:"-"`
	CreatedAt time.Time `json:"-"`
}
