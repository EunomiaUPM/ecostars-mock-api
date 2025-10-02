package database

import (
	"ecostars-fake-api/internal/domain"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&domain.Hotel{})
	db.AutoMigrate(&domain.HotelEvenModel{})
	db.AutoMigrate(&domain.Subscription{})
	db.AutoMigrate(&domain.HotelMeasure{})
	return db, nil
}
