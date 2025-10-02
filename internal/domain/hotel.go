package domain

import (
	"gorm.io/gorm"
)

type HotelModel struct {
	gorm.Model
	Name    string `json:"name"`
	Address string `json:"address"`
	City    string `json:"city"`
}

type Hotel struct {
	Name     string         `json:"name"`
	Address  string         `json:"address"`
	City     string         `json:"city"`
	Measures []HotelMeasure `json:"measures,omitempty"`
}

type HotelCreateRequest struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
	City    string `json:"city" binding:"required"`
}

type HotelUpdateRequest struct {
	Name    *string `json:"name"`
	Address *string `json:"address"`
	City    *string `json:"city"`
}
