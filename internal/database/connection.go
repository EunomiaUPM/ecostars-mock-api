package database

import (
	"ecostars-fake-api/internal/config"
	"ecostars-fake-api/internal/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(config *config.Config) (*gorm.DB, error) {
	dsn := config.GetDatabaseDSN()
	postres_conn := postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	})
	db, err := gorm.Open(postres_conn, &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&domain.HotelModel{}, &domain.HotelEvenModel{}, &domain.Subscription{}, &domain.HotelMeasure{}, &domain.MetricItemModel{})
	return db, nil
}
