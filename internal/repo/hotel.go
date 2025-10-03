package repo

import (
	"context"
	"ecostars-fake-api/internal/domain"

	"gorm.io/gorm"
)

type HotelRepository struct {
	DB *gorm.DB
}

func (repo *HotelRepository) GetAllHotels(ctx context.Context) ([]domain.HotelModel, error) {
	var hotels []domain.HotelModel
	err := repo.DB.WithContext(ctx).Find(&hotels).Error
	return hotels, err
}

func (repo *HotelRepository) CreateHotel(ctx context.Context, newHotel *domain.HotelCreateRequest) (domain.HotelModel, error) {
	hotel := domain.HotelModel{
		Name:    newHotel.Name,
		Address: newHotel.Address,
		City:    newHotel.City,
	}
	err := repo.DB.WithContext(ctx).Create(&hotel).Error
	return hotel, err
}

func (repo *HotelRepository) GetHotelByID(ctx context.Context, id uint) (domain.HotelModel, error) {
	return domain.HotelModel{}, nil
}

func (repo *HotelRepository) UpdateHotel(ctx context.Context, id uint, updatedHotel *domain.HotelUpdateRequest) (domain.HotelModel, error) {
	return domain.HotelModel{}, nil
}

func (repo *HotelRepository) DeleteHotel(ctx context.Context, id uint) error {
	return nil
}
