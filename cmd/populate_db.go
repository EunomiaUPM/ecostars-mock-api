package main

import (
	"context"
	"ecostars-fake-api/internal/config"
	"ecostars-fake-api/internal/database"
	"ecostars-fake-api/internal/domain"
	"ecostars-fake-api/internal/repo"

	"github.com/go-faker/faker/v4"
)

func PopulateDB(config *config.Config) error {
	db, err := database.ConnectDB(config)
	if err != nil {
		panic("failed to connect database")
	}
	_ = db
	hotelRepository := repo.HotelRepository{
		DB: db,
	}
	hotelMeasureRepository := repo.HotelMeasureRepository{
		DB: db,
	}

	// Create some hotels and measures
	// hotel 1
	createHotelAndMeasures(&hotelRepository, &hotelMeasureRepository)
	// hotel 2
	createHotelAndMeasures(&hotelRepository, &hotelMeasureRepository)
	// hotel 3
	createHotelAndMeasures(&hotelRepository, &hotelMeasureRepository)
	// hotel 4
	createHotelAndMeasures(&hotelRepository, &hotelMeasureRepository)
	// hotel 5
	createHotelAndMeasures(&hotelRepository, &hotelMeasureRepository)

	return nil
}

func createHotelAndMeasures(hotelRepo *repo.HotelRepository, measureRepo *repo.HotelMeasureRepository) error {
	var hotelRequest domain.HotelCreateRequest
	faker.FakeData(&hotelRequest)
	hotel, _ := hotelRepo.CreateHotel(context.Background(), &hotelRequest)

	var measureRequest domain.HotelMeasureCreateRequest
	faker.FakeData(&measureRequest)
	measureRequest.HotelID = hotel.ID
	measureRepo.CreateHotelMeasure(context.Background(), &measureRequest)

	faker.FakeData(&measureRequest)
	measureRequest.HotelID = hotel.ID
	measureRepo.CreateHotelMeasure(context.Background(), &measureRequest)

	faker.FakeData(&measureRequest)
	measureRequest.HotelID = hotel.ID
	measureRepo.CreateHotelMeasure(context.Background(), &measureRequest)

	faker.FakeData(&measureRequest)
	measureRequest.HotelID = hotel.ID
	measureRepo.CreateHotelMeasure(context.Background(), &measureRequest)

	return nil
}
