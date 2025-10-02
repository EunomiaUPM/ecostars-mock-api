package repo

import (
	"ecostars-fake-api/internal/domain"

	"gorm.io/gorm"
)

type HotelEventRepository struct {
	DB *gorm.DB
}

func (repo *HotelEventRepository) TableName() string {
	return "hotel_events"
}

func (repo *HotelEventRepository) CreateHotelEvent(newEvent *domain.HotelEventCreateRequest) (domain.HotelEvent, error) {
	event := domain.HotelEvent{
		HotelID:      newEvent.HotelID,
		EventName:    newEvent.EventName,
		EventContent: newEvent.EventContent,
	}
	err := repo.DB.Create(&event).Error
	return event, err
}

func (repo *HotelEventRepository) GetAllHotelEvents() ([]domain.HotelEvent, error) {
	var events []domain.HotelEvent
	err := repo.DB.Find(&events).Error
	return events, err
}

func (repo *HotelEventRepository) GetHotelEventByID(id uint) (domain.HotelEvent, error) {
	var event domain.HotelEvent
	err := repo.DB.First(&event, id).Error
	return event, err
}

func (repo *HotelEventRepository) DeleteHotelEvent(id uint) error {
	return repo.DB.Delete(&domain.HotelEvent{}, id).Error
}
