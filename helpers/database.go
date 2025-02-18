package helpers

import (
	"fmt"
	"log"

	alamatmodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/alamat_model"
	fotoprodukmodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/foto_produk_model"
	kategorimodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/kategori_model"
	produkmodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/produk_model"
	tokomodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/toko_model"
	trxmodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/trx_model"
	usermodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/user_model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadDatabase() {
	var err error
	host := GetEnv("DB_HOST", "")
	port := GetEnv("DB_PORT", "3306")
	user := GetEnv("DB_USER", "")
	password := GetEnv("DB_PASSWORD", "")
	dbname := GetEnv("DB_NAME", "")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", user, password, host, port, dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DB = db

	_, err = db.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}

	log.Println("Connected to database")

	DB.AutoMigrate(
		&usermodel.User{},
		&alamatmodel.Alamat{},
		&tokomodel.Toko{},
		&kategorimodel.Kategori{},
		&produkmodel.Produk{},
		&fotoprodukmodel.FotoProduk{},
		&trxmodel.Trx{},
	)
}
