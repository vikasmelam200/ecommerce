package config

import (
	"log"
	"onvif310824/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=postgres password=Vikas@2000 dbname=vms port=5432 sslmode=disable"

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("error while connecting to database %v", err)
	}
	DB.AutoMigrate(&models.Camera{}, &models.StreamLog{})
}
func GetDB() *gorm.DB {
	if DB == nil {
		log.Fatal("Database not initialized")
	}
	return DB
}
