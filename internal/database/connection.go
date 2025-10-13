package database

import (
	"ecostars-fake-api/internal/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	// dsn := "host=postgres user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=10000 sslmode=disable"
	postres_conn := postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	})
	db, err := gorm.Open(postres_conn, &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&domain.HotelModel{}, &domain.HotelEvenModel{}, &domain.Subscription{}, &domain.HotelMeasure{})
	return db, nil
}
