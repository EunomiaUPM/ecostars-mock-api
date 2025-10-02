package repo

import (
	"context"
	"ecostars-fake-api/internal/domain"

	"gorm.io/gorm"
)

type HotelMeasureRepository struct {
	DB *gorm.DB
}

func (repo *HotelMeasureRepository) GetHotelMeasuresByHotelID(ctx context.Context, d uint) ([]domain.HotelMeasure, error) {
	var measures []domain.HotelMeasure
	err := repo.DB.WithContext(ctx).Where("hotel_id = ?", d).Find(&measures).Error
	return measures, err
}

func (repo *HotelMeasureRepository) TableName() string {
	return "hotel_measures"
}

func (repo *HotelMeasureRepository) GetAllHotelMeasures(ctx context.Context) ([]domain.HotelMeasure, error) {
	var measures []domain.HotelMeasure
	err := repo.DB.WithContext(ctx).Find(&measures).Error
	return measures, err
}

func (repo *HotelMeasureRepository) CreateHotelMeasure(ctx context.Context, newMeasure *domain.HotelMeasureCreateRequest) (domain.HotelMeasure, error) {
	measure := domain.HotelMeasure{
		HotelID:                            newMeasure.HotelID,
		Year:                               newMeasure.Year,
		Rooms:                              newMeasure.Rooms,
		Stays:                              newMeasure.Stays,
		GasKwh:                             newMeasure.GasKwh,
		WaterM3:                            newMeasure.WaterM3,
		Occupied:                           newMeasure.Occupied,
		Deviation:                          newMeasure.Deviation,
		EnergyKwh:                          newMeasure.EnergyKwh,
		RatingCo2:                          newMeasure.RatingCo2,
		TotalArea:                          newMeasure.TotalArea,
		StaysPerRN:                         newMeasure.StaysPerRN,
		Electricity:                        newMeasure.Electricity,
		RatingCo2RN:                        newMeasure.RatingCo2RN,
		RatingWater:                        newMeasure.RatingWater,
		DeviationCo2:                       newMeasure.DeviationCo2,
		RatingEnergy:                       newMeasure.RatingEnergy,
		RatingCo2S1RN:                      newMeasure.RatingCo2S1RN,
		RatingCo2S2RN:                      newMeasure.RatingCo2S2RN,
		RatingCo2Stay:                      newMeasure.RatingCo2Stay,
		RatingWaterRN:                      newMeasure.RatingWaterRN,
		DeviationCo2RN:                     newMeasure.DeviationCo2RN,
		DeviationWaste:                     newMeasure.DeviationWaste,
		DeviationWater:                     newMeasure.DeviationWater,
		ElectricityKwh:                     newMeasure.ElectricityKwh,
		RatingEnergyRN:                     newMeasure.RatingEnergyRN,
		TotalEmissions:                     newMeasure.TotalEmissions,
		Co2Compensation:                    newMeasure.Co2Compensation,
		DeviationEnergy:                    newMeasure.DeviationEnergy,
		RatingCo2S1Stay:                    newMeasure.RatingCo2S1Stay,
		RatingCo2S2Stay:                    newMeasure.RatingCo2S2Stay,
		RatingWaterArea:                    newMeasure.RatingWaterArea,
		RatingWaterRoom:                    newMeasure.RatingWaterRoom,
		RatingWaterStay:                    newMeasure.RatingWaterStay,
		Scope1Emissions:                    newMeasure.Scope1Emissions,
		Scope2Emissions:                    newMeasure.Scope2Emissions,
		DeviationCo2S1RN:                   newMeasure.DeviationCo2S1RN,
		DeviationCo2S2RN:                   newMeasure.DeviationCo2S2RN,
		DeviationCo2S3RN:                   newMeasure.DeviationCo2S3RN,
		DeviationCo2Stay:                   newMeasure.DeviationCo2Stay,
		DeviationWasteRN:                   newMeasure.DeviationWasteRN,
		DeviationWaterRN:                   newMeasure.DeviationWaterRN,
		DirectStationary:                   newMeasure.DirectStationary,
		RatingEnergyArea:                   newMeasure.RatingEnergyArea,
		RatingEnergyRoom:                   newMeasure.RatingEnergyRoom,
		RatingEnergyStay:                   newMeasure.RatingEnergyStay,
		RecommendedStars:                   newMeasure.RecommendedStars,
		DeviationEnergyRN:                  newMeasure.DeviationEnergyRN,
		DeviationCo2S1Stay:                 newMeasure.DeviationCo2S1Stay,
		DeviationCo2S2Stay:                 newMeasure.DeviationCo2S2Stay,
		DeviationCo2S3Stay:                 newMeasure.DeviationCo2S3Stay,
		DeviationWasteStay:                 newMeasure.DeviationWasteStay,
		DeviationWaterStay:                 newMeasure.DeviationWaterStay,
		RecommendedStarsV1:                 newMeasure.RecommendedStarsV1,
		RecommendedStarsV2:                 newMeasure.RecommendedStarsV2,
		Scope2Compensation:                 newMeasure.Scope2Compensation,
		DeviationEnergyStay:                newMeasure.DeviationEnergyStay,
		EnergyKwhWithCo2Compensation:       newMeasure.EnergyKwhWithCo2Compensation,
		RatingEnergyWithCo2Compensation:    newMeasure.RatingEnergyWithCo2Compensation,
		DeviationEnergyWithCo2Compensation: newMeasure.DeviationEnergyWithCo2Compensation,
	}
	err := repo.DB.WithContext(ctx).Create(&measure).Error
	return measure, err
}

func (repo *HotelMeasureRepository) GetHotelMeasureByID(ctx context.Context, id uint) (domain.HotelMeasure, error) {
	var measure domain.HotelMeasure
	err := repo.DB.WithContext(ctx).First(&measure, id).Error
	return measure, err
}

func (repo *HotelMeasureRepository) UpdateHotelMeasure(ctx context.Context, id uint, updatedMeasure *domain.HotelMeasureUpdateRequest) (domain.HotelMeasure, error) {
	var measure domain.HotelMeasure
	if err := repo.DB.WithContext(ctx).First(&measure, id).Error; err != nil {
		return measure, err
	}
	if updatedMeasure.Rooms != nil {
		measure.Rooms = *updatedMeasure.Rooms
	}
	if updatedMeasure.Stays != nil {
		measure.Stays = *updatedMeasure.Stays
	}
	if updatedMeasure.GasKwh != nil {
		measure.GasKwh = *updatedMeasure.GasKwh
	}
	if updatedMeasure.WaterM3 != nil {
		measure.WaterM3 = *updatedMeasure.WaterM3
	}
	if updatedMeasure.Occupied != nil {
		measure.Occupied = *updatedMeasure.Occupied
	}
	if updatedMeasure.Deviation != nil {
		measure.Deviation = *updatedMeasure.Deviation
	}
	if updatedMeasure.EnergyKwh != nil {
		measure.EnergyKwh = *updatedMeasure.EnergyKwh
	}
	if updatedMeasure.RatingCo2 != nil {
		measure.RatingCo2 = *updatedMeasure.RatingCo2
	}
	if updatedMeasure.TotalArea != nil {
		measure.TotalArea = *updatedMeasure.TotalArea
	}
	if updatedMeasure.StaysPerRN != nil {
		measure.StaysPerRN = *updatedMeasure.StaysPerRN
	}
	if updatedMeasure.Electricity != nil {
		measure.Electricity = *updatedMeasure.Electricity
	}
	if updatedMeasure.RatingCo2RN != nil {
		measure.RatingCo2RN = *updatedMeasure.RatingCo2RN
	}
	if updatedMeasure.RatingWater != nil {
		measure.RatingWater = *updatedMeasure.RatingWater
	}
	if updatedMeasure.DeviationCo2 != nil {
		measure.DeviationCo2 = *updatedMeasure.DeviationCo2
	}
	if updatedMeasure.RatingEnergy != nil {
		measure.RatingEnergy = *updatedMeasure.RatingEnergy
	}
	if updatedMeasure.RatingCo2S1RN != nil {
		measure.RatingCo2S1RN = *updatedMeasure.RatingCo2S1RN
	}
	if updatedMeasure.RatingCo2S2RN != nil {
		measure.RatingCo2S2RN = *updatedMeasure.RatingCo2S2RN
	}
	if updatedMeasure.RatingCo2Stay != nil {
		measure.RatingCo2Stay = *updatedMeasure.RatingCo2Stay
	}
	if updatedMeasure.RatingWaterRN != nil {
		measure.RatingWaterRN = *updatedMeasure.RatingWaterRN
	}
	if updatedMeasure.DeviationCo2RN != nil {
		measure.DeviationCo2RN = *updatedMeasure.DeviationCo2RN
	}
	if updatedMeasure.DeviationWaste != nil {
		measure.DeviationWaste = *updatedMeasure.DeviationWaste
	}
	if updatedMeasure.DeviationWater != nil {
		measure.DeviationWater = *updatedMeasure.DeviationWater
	}
	if updatedMeasure.ElectricityKwh != nil {
		measure.ElectricityKwh = *updatedMeasure.ElectricityKwh
	}
	if updatedMeasure.RatingEnergyRN != nil {
		measure.RatingEnergyRN = *updatedMeasure.RatingEnergyRN
	}
	if updatedMeasure.TotalEmissions != nil {
		measure.TotalEmissions = *updatedMeasure.TotalEmissions
	}
	if updatedMeasure.Co2Compensation != nil {
		measure.Co2Compensation = *updatedMeasure.Co2Compensation
	}
	if updatedMeasure.DeviationEnergy != nil {
		measure.DeviationEnergy = *updatedMeasure.DeviationEnergy
	}
	if updatedMeasure.RatingCo2S1Stay != nil {
		measure.RatingCo2S1Stay = *updatedMeasure.RatingCo2S1Stay
	}
	if updatedMeasure.RatingCo2S2Stay != nil {
		measure.RatingCo2S2Stay = *updatedMeasure.RatingCo2S2Stay
	}
	if updatedMeasure.RatingWaterArea != nil {
		measure.RatingWaterArea = *updatedMeasure.RatingWaterArea
	}
	if updatedMeasure.RatingWaterRoom != nil {
		measure.RatingWaterRoom = *updatedMeasure.RatingWaterRoom
	}
	if updatedMeasure.RatingWaterStay != nil {
		measure.RatingWaterStay = *updatedMeasure.RatingWaterStay
	}
	if updatedMeasure.Scope1Emissions != nil {
		measure.Scope1Emissions = *updatedMeasure.Scope1Emissions
	}
	if updatedMeasure.Scope2Emissions != nil {
		measure.Scope2Emissions = *updatedMeasure.Scope2Emissions
	}
	if updatedMeasure.DeviationCo2S1RN != nil {
		measure.DeviationCo2S1RN = *updatedMeasure.DeviationCo2S1RN
	}
	if updatedMeasure.DeviationCo2S2RN != nil {
		measure.DeviationCo2S2RN = *updatedMeasure.DeviationCo2S2RN
	}
	if updatedMeasure.DeviationCo2S3RN != nil {
		measure.DeviationCo2S3RN = *updatedMeasure.DeviationCo2S3RN
	}
	if updatedMeasure.DeviationCo2Stay != nil {
		measure.DeviationCo2Stay = *updatedMeasure.DeviationCo2Stay
	}
	if updatedMeasure.DeviationWasteRN != nil {
		measure.DeviationWasteRN = *updatedMeasure.DeviationWasteRN
	}
	if updatedMeasure.DeviationWaterRN != nil {
		measure.DeviationWaterRN = *updatedMeasure.DeviationWaterRN
	}
	if updatedMeasure.DirectStationary != nil {
		measure.DirectStationary = *updatedMeasure.DirectStationary
	}
	if updatedMeasure.RatingEnergyArea != nil {
		measure.RatingEnergyArea = *updatedMeasure.RatingEnergyArea
	}
	if updatedMeasure.RatingEnergyRoom != nil {
		measure.RatingEnergyRoom = *updatedMeasure.RatingEnergyRoom
	}
	if updatedMeasure.RatingEnergyStay != nil {
		measure.RatingEnergyStay = *updatedMeasure.RatingEnergyStay
	}
	if updatedMeasure.RecommendedStars != nil {
		measure.RecommendedStars = *updatedMeasure.RecommendedStars
	}
	if updatedMeasure.DeviationEnergyRN != nil {
		measure.DeviationEnergyRN = *updatedMeasure.DeviationEnergyRN
	}
	if updatedMeasure.DeviationCo2S1Stay != nil {
		measure.DeviationCo2S1Stay = *updatedMeasure.DeviationCo2S1Stay
	}
	if updatedMeasure.DeviationCo2S2Stay != nil {
		measure.DeviationCo2S2Stay = *updatedMeasure.DeviationCo2S2Stay
	}
	if updatedMeasure.DeviationCo2S3Stay != nil {
		measure.DeviationCo2S3Stay = *updatedMeasure.DeviationCo2S3Stay
	}
	if updatedMeasure.DeviationWasteStay != nil {
		measure.DeviationWasteStay = *updatedMeasure.DeviationWasteStay
	}
	if updatedMeasure.DeviationWaterStay != nil {
		measure.DeviationWaterStay = *updatedMeasure.DeviationWaterStay
	}
	if updatedMeasure.RecommendedStarsV1 != nil {
		measure.RecommendedStarsV1 = *updatedMeasure.RecommendedStarsV1
	}
	if updatedMeasure.RecommendedStarsV2 != nil {
		measure.RecommendedStarsV2 = *updatedMeasure.RecommendedStarsV2
	}
	if updatedMeasure.Scope2Compensation != nil {
		measure.Scope2Compensation = *updatedMeasure.Scope2Compensation
	}
	if updatedMeasure.DeviationEnergyStay != nil {
		measure.DeviationEnergyStay = *updatedMeasure.DeviationEnergyStay
	}
	if updatedMeasure.EnergyKwhWithCo2Compensation != nil {
		measure.EnergyKwhWithCo2Compensation = *updatedMeasure.EnergyKwhWithCo2Compensation
	}
	if updatedMeasure.RatingEnergyWithCo2Compensation != nil {
		measure.RatingEnergyWithCo2Compensation = *updatedMeasure.RatingEnergyWithCo2Compensation
	}
	if updatedMeasure.DeviationEnergyWithCo2Compensation != nil {
		measure.DeviationEnergyWithCo2Compensation = *updatedMeasure.DeviationEnergyWithCo2Compensation
	}
	err := repo.DB.WithContext(ctx).Save(&measure).Error
	return measure, err
}

func (repo *HotelMeasureRepository) DeleteHotelMeasure(ctx context.Context, id uint) error {
	return repo.DB.WithContext(ctx).Delete(&domain.HotelMeasure{}, id).Error
}
