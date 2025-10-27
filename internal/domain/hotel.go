package domain

import (
	"gorm.io/gorm"
)

type HotelModel struct {
	gorm.Model `gorm:"tableName:hotels"`
	Name       string         `json:"name"`
	Address    string         `json:"address"`
	City       string         `json:"city"`
	Measures   []HotelMeasure `json:"-" gorm:"foreignKey:HotelID"`
}

func (HotelModel) TableName() string {
	return "hotels"
}

type Hotel struct {
	Name     string         `json:"name"`
	Address  string         `json:"address"`
	City     string         `json:"city"`
	Measures []HotelMeasure `json:"measures,omitempty"`
}

type HotelCreateRequest struct {
	Name    string `json:"name" binding:"required" faker:"word"`
	Address string `json:"address" binding:"required" faker:"sentence"`
	City    string `json:"city" binding:"required" faker:"oneof:New York,Los Angeles,Chicago,Houston,Phoenix"`
}

type HotelUpdateRequest struct {
	Name    *string `json:"name"`
	Address *string `json:"address"`
	City    *string `json:"city"`
}
