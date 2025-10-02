package services

import (
	"context"

	"ecostars-fake-api/internal/domain"
	"ecostars-fake-api/internal/repo"
)

type HotelService struct {
	HotelRepository        *repo.HotelRepository
	HotelMeasureRepository *repo.HotelMeasureRepository
}

func (service *HotelService) CreateHotel(ctx context.Context, newHotel *domain.HotelCreateRequest) (domain.Hotel, error) {
	hotel, err := service.HotelRepository.CreateHotel(ctx, newHotel)
	measures := []domain.HotelMeasure{}
	hotelWithMeasures := domain.Hotel{
		Name:     hotel.Name,
		Address:  hotel.Address,
		City:     hotel.City,
		Measures: measures,
	}
	return hotelWithMeasures, err
}

func (service *HotelService) GetAllHotels(ctx context.Context) ([]domain.Hotel, error) {
	hotels, err := service.HotelRepository.GetAllHotels(ctx)
	if err != nil {
		return nil, err
	}
	var hotelsWithMeasures []domain.Hotel
	for _, hotel := range hotels {
		measures, err := service.HotelMeasureRepository.GetHotelMeasuresByHotelID(ctx, hotel.ID)
		if err != nil {
			return nil, err
		}
		hotelWithMeasures := domain.Hotel{
			Name:     hotel.Name,
			Address:  hotel.Address,
			City:     hotel.City,
			Measures: measures,
		}
		hotelsWithMeasures = append(hotelsWithMeasures, hotelWithMeasures)
	}
	return hotelsWithMeasures, nil
}

func (service *HotelService) GetHotelByID(ctx context.Context, id uint) (domain.Hotel, error) {
	hotel, err := service.HotelRepository.GetHotelByID(ctx, id)
	if err != nil {
		return domain.Hotel{}, err
	}
	measures, err := service.HotelMeasureRepository.GetHotelMeasuresByHotelID(ctx, hotel.ID)
	if err != nil {
		return domain.Hotel{}, err
	}
	hotelWithMeasures := domain.Hotel{
		Name:     hotel.Name,
		Address:  hotel.Address,
		City:     hotel.City,
		Measures: measures,
	}
	return hotelWithMeasures, nil
}

func (service *HotelService) CreateHotelMeasure(ctx context.Context, newMeasure *domain.HotelMeasureCreateRequest) (domain.HotelMeasure, error) {
	return service.HotelMeasureRepository.CreateHotelMeasure(ctx, newMeasure)
}

func (service *HotelService) GetAllHotelMeasures(ctx context.Context) ([]domain.HotelMeasure, error) {
	return service.HotelMeasureRepository.GetAllHotelMeasures(ctx)
}

func (service *HotelService) GetHotelMeasureByID(ctx context.Context, id uint) (domain.HotelMeasure, error) {
	return service.HotelMeasureRepository.GetHotelMeasureByID(ctx, id)
}

func (service *HotelService) GetHotelMeasuresByHotelID(ctx context.Context, hotelID uint) ([]domain.HotelMeasure, error) {
	return service.HotelMeasureRepository.GetHotelMeasuresByHotelID(ctx, hotelID)
}

func (service *HotelService) DeleteHotelMeasure(ctx context.Context, id uint) error {
	return service.HotelMeasureRepository.DeleteHotelMeasure(ctx, id)
}
