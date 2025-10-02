package domain

import (
	"encoding/json"

	"gorm.io/gorm"
)

type HotelEvenModel struct {
	gorm.Model
	HotelID      uint   `json:"hotel_id" binding:"required"`
	EventName    string `json:"event_name" binding:"required"`
	EventContent string `json:"event_content" binding:"required"`
}

type HotelEvent struct {
	gorm.Model
	HotelID      uint
	EventName    string
	EventContent interface{}
}

type HotelEventCreateRequest struct {
	HotelID      uint        `json:"hotel_id" binding:"required"`
	EventName    string      `json:"event_name" binding:"required"`
	EventContent interface{} `json:"event_content" binding:"required"`
}

func (h *HotelEvent) MapToModel() (HotelEvenModel, error) {
	contentBytes, err := json.Marshal(h.EventContent)
	if err != nil {
		return HotelEvenModel{}, err
	}
	return HotelEvenModel{
		HotelID:      h.HotelID,
		EventName:    h.EventName,
		EventContent: string(contentBytes),
	}, nil
}
