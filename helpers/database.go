package helpers

import (
	"fmt"
	"log"

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
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)

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

	// DB.AutoMigrate()
}
