package domain

import (
	"gorm.io/gorm"
)

type HotelMeasure struct {
	gorm.Model
	HotelID                            uint    `json:"hotel_id"`
	Year                               int     `json:"year"`
	Rooms                              int     `json:"rooms"`
	Stays                              int     `json:"stays"`
	GasKwh                             float64 `json:"gasKwh"`
	WaterM3                            float64 `json:"waterM3"`
	Occupied                           int     `json:"occupied"`
	Deviation                          float64 `json:"deviation"`
	EnergyKwh                          float64 `json:"energyKwh"`
	RatingCo2                          float64 `json:"ratingCo2"`
	TotalArea                          float64 `json:"totalArea"`
	StaysPerRN                         float64 `json:"staysPerRN"`
	Electricity                        float64 `json:"electricity"`
	RatingCo2RN                        float64 `json:"ratingCo2RN"`
	RatingWater                        float64 `json:"ratingWater"`
	DeviationCo2                       float64 `json:"deviationCo2"`
	RatingEnergy                       float64 `json:"ratingEnergy"`
	RatingCo2S1RN                      float64 `json:"ratingCo2S1RN"`
	RatingCo2S2RN                      float64 `json:"ratingCo2S2RN"`
	RatingCo2Stay                      float64 `json:"ratingCo2Stay"`
	RatingWaterRN                      float64 `json:"ratingWaterRN"`
	DeviationCo2RN                     float64 `json:"deviationCo2RN"`
	DeviationWaste                     float64 `json:"deviationWaste"`
	DeviationWater                     float64 `json:"deviationWater"`
	ElectricityKwh                     float64 `json:"electricityKwh"`
	RatingEnergyRN                     float64 `json:"ratingEnergyRN"`
	TotalEmissions                     float64 `json:"totalEmissions"`
	Co2Compensation                    float64 `json:"co2Compensation"`
	DeviationEnergy                    float64 `json:"deviationEnergy"`
	RatingCo2S1Stay                    float64 `json:"ratingCo2S1Stay"`
	RatingCo2S2Stay                    float64 `json:"ratingCo2S2Stay"`
	RatingWaterArea                    float64 `json:"ratingWaterArea"`
	RatingWaterRoom                    float64 `json:"ratingWaterRoom"`
	RatingWaterStay                    float64 `json:"ratingWaterStay"`
	Scope1Emissions                    float64 `json:"scope1Emissions"`
	Scope2Emissions                    float64 `json:"scope2Emissions"`
	DeviationCo2S1RN                   float64 `json:"deviationCo2S1RN"`
	DeviationCo2S2RN                   float64 `json:"deviationCo2S2RN"`
	DeviationCo2S3RN                   float64 `json:"deviationCo2S3RN"`
	DeviationCo2Stay                   float64 `json:"deviationCo2Stay"`
	DeviationWasteRN                   float64 `json:"deviationWasteRN"`
	DeviationWaterRN                   float64 `json:"deviationWaterRN"`
	DirectStationary                   float64 `json:"directStationary"`
	RatingEnergyArea                   float64 `json:"ratingEnergyArea"`
	RatingEnergyRoom                   float64 `json:"ratingEnergyRoom"`
	RatingEnergyStay                   float64 `json:"ratingEnergyStay"`
	RecommendedStars                   float64 `json:"recommendedStars"`
	DeviationEnergyRN                  float64 `json:"deviationEnergyRN"`
	DeviationCo2S1Stay                 float64 `json:"deviationCo2S1Stay"`
	DeviationCo2S2Stay                 float64 `json:"deviationCo2S2Stay"`
	DeviationCo2S3Stay                 float64 `json:"deviationCo2S3Stay"`
	DeviationWasteStay                 float64 `json:"deviationWasteStay"`
	DeviationWaterStay                 float64 `json:"deviationWaterStay"`
	RecommendedStarsV1                 float64 `json:"recommendedStarsV1"`
	RecommendedStarsV2                 float64 `json:"recommendedStarsV2"`
	Scope2Compensation                 float64 `json:"scope2Compensation"`
	DeviationEnergyStay                float64 `json:"deviationEnergyStay"`
	EnergyKwhWithCo2Compensation       float64 `json:"energyKwhWithCo2Compensation"`
	RatingEnergyWithCo2Compensation    float64 `json:"ratingEnergyWithCo2Compensation"`
	DeviationEnergyWithCo2Compensation float64 `json:"deviationEnergyWithCo2Compensation"`
}

type HotelMeasureCreateRequest struct {
	HotelID                            uint    `json:"hotel_id" binding:"required"`
	Year                               int     `json:"year" binding:"required"`
	Rooms                              int     `json:"rooms" binding:"required"`
	Stays                              int     `json:"stays" binding:"required"`
	GasKwh                             float64 `json:"gasKwh" binding:"required"`
	WaterM3                            float64 `json:"waterM3" binding:"required"`
	Occupied                           int     `json:"occupied" binding:"required"`
	Deviation                          float64 `json:"deviation" binding:"required"`
	EnergyKwh                          float64 `json:"energyKwh" binding:"required"`
	RatingCo2                          float64 `json:"ratingCo2" binding:"required"`
	TotalArea                          float64 `json:"totalArea" binding:"required"`
	StaysPerRN                         float64 `json:"staysPerRN" binding:"required"`
	Electricity                        float64 `json:"electricity" binding:"required"`
	RatingCo2RN                        float64 `json:"ratingCo2RN" binding:"required"`
	RatingWater                        float64 `json:"ratingWater" binding:"required"`
	DeviationCo2                       float64 `json:"deviationCo2" binding:"required"`
	RatingEnergy                       float64 `json:"ratingEnergy" binding:"required"`
	RatingCo2S1RN                      float64 `json:"ratingCo2S1RN" binding:"required"`
	RatingCo2S2RN                      float64 `json:"ratingCo2S2RN" binding:"required"`
	RatingCo2Stay                      float64 `json:"ratingCo2Stay" binding:"required"`
	RatingWaterRN                      float64 `json:"ratingWaterRN" binding:"required"`
	DeviationCo2RN                     float64 `json:"deviationCo2RN" binding:"required"`
	DeviationWaste                     float64 `json:"deviationWaste" binding:"required"`
	DeviationWater                     float64 `json:"deviationWater" binding:"required"`
	ElectricityKwh                     float64 `json:"electricityKwh" binding:"required"`
	RatingEnergyRN                     float64 `json:"ratingEnergyRN" binding:"required"`
	TotalEmissions                     float64 `json:"totalEmissions" binding:"required"`
	Co2Compensation                    float64 `json:"co2Compensation" binding:"required"`
	DeviationEnergy                    float64 `json:"deviationEnergy" binding:"required"`
	RatingCo2S1Stay                    float64 `json:"ratingCo2S1Stay" binding:"required"`
	RatingCo2S2Stay                    float64 `json:"ratingCo2S2Stay" binding:"required"`
	RatingWaterArea                    float64 `json:"ratingWaterArea" binding:"required"`
	RatingWaterRoom                    float64 `json:"ratingWaterRoom" binding:"required"`
	RatingWaterStay                    float64 `json:"ratingWaterStay" binding:"required"`
	Scope1Emissions                    float64 `json:"scope1Emissions" binding:"required"`
	Scope2Emissions                    float64 `json:"scope2Emissions" binding:"required"`
	DeviationCo2S1RN                   float64 `json:"deviationCo2S1RN" binding:"required"`
	DeviationCo2S2RN                   float64 `json:"deviationCo2S2RN" binding:"required"`
	DeviationCo2S3RN                   float64 `json:"deviationCo2S3RN" binding:"required"`
	DeviationCo2Stay                   float64 `json:"deviationCo2Stay" binding:"required"`
	DeviationWasteRN                   float64 `json:"deviationWasteRN" binding:"required"`
	DeviationWaterRN                   float64 `json:"deviationWaterRN" binding:"required"`
	DirectStationary                   float64 `json:"directStationary" binding:"required"`
	RatingEnergyArea                   float64 `json:"ratingEnergyArea" binding:"required"`
	RatingEnergyRoom                   float64 `json:"ratingEnergyRoom" binding:"required"`
	RatingEnergyStay                   float64 `json:"ratingEnergyStay" binding:"required"`
	RecommendedStars                   float64 `json:"recommendedStars" binding:"required"`
	DeviationEnergyRN                  float64 `json:"deviationEnergyRN" binding:"required"`
	DeviationCo2S1Stay                 float64 `json:"deviationCo2S1Stay" binding:"required"`
	DeviationCo2S2Stay                 float64 `json:"deviationCo2S2Stay" binding:"required"`
	DeviationCo2S3Stay                 float64 `json:"deviationCo2S3Stay" binding:"required"`
	DeviationWasteStay                 float64 `json:"deviationWasteStay" binding:"required"`
	DeviationWaterStay                 float64 `json:"deviationWaterStay" binding:"required"`
	RecommendedStarsV1                 float64 `json:"recommendedStarsV1" binding:"required"`
	RecommendedStarsV2                 float64 `json:"recommendedStarsV2" binding:"required"`
	Scope2Compensation                 float64 `json:"scope2Compensation" binding:"required"`
	DeviationEnergyStay                float64 `json:"deviationEnergyStay" binding:"required"`
	EnergyKwhWithCo2Compensation       float64 `json:"energyKwhWithCo2Compensation" binding:"required"`
	RatingEnergyWithCo2Compensation    float64 `json:"ratingEnergyWithCo2Compensation" binding:"required"`
	DeviationEnergyWithCo2Compensation float64 `json:"deviationEnergyWithCo2Compensation" binding:"required"`
}

type HotelMeasureUpdateRequest struct {
	Rooms                              *int     `json:"rooms"`
	Stays                              *int     `json:"stays"`
	GasKwh                             *float64 `json:"gasKwh"`
	WaterM3                            *float64 `json:"waterM3"`
	Occupied                           *int     `json:"occupied"`
	Deviation                          *float64 `json:"deviation"`
	EnergyKwh                          *float64 `json:"energyKwh"`
	RatingCo2                          *float64 `json:"ratingCo2"`
	TotalArea                          *float64 `json:"totalArea"`
	StaysPerRN                         *float64 `json:"staysPerRN"`
	Electricity                        *float64 `json:"electricity"`
	RatingCo2RN                        *float64 `json:"ratingCo2RN"`
	RatingWater                        *float64 `json:"ratingWater"`
	DeviationCo2                       *float64 `json:"deviationCo2"`
	RatingEnergy                       *float64 `json:"ratingEnergy"`
	RatingCo2S1RN                      *float64 `json:"ratingCo2S1RN"`
	RatingCo2S2RN                      *float64 `json:"ratingCo2S2RN"`
	RatingCo2Stay                      *float64 `json:"ratingCo2Stay"`
	RatingWaterRN                      *float64 `json:"ratingWaterRN"`
	DeviationCo2RN                     *float64 `json:"deviationCo2RN"`
	DeviationWaste                     *float64 `json:"deviationWaste"`
	DeviationWater                     *float64 `json:"deviationWater"`
	ElectricityKwh                     *float64 `json:"electricityKwh"`
	RatingEnergyRN                     *float64 `json:"ratingEnergyRN"`
	TotalEmissions                     *float64 `json:"totalEmissions"`
	Co2Compensation                    *float64 `json:"co2Compensation"`
	DeviationEnergy                    *float64 `json:"deviationEnergy"`
	RatingCo2S1Stay                    *float64 `json:"ratingCo2S1Stay"`
	RatingCo2S2Stay                    *float64 `json:"ratingCo2S2Stay"`
	RatingWaterArea                    *float64 `json:"ratingWaterArea"`
	RatingWaterRoom                    *float64 `json:"ratingWaterRoom"`
	RatingWaterStay                    *float64 `json:"ratingWaterStay"`
	Scope1Emissions                    *float64 `json:"scope1Emissions"`
	Scope2Emissions                    *float64 `json:"scope2Emissions"`
	DeviationCo2S1RN                   *float64 `json:"deviationCo2S1RN"`
	DeviationCo2S2RN                   *float64 `json:"deviationCo2S2RN"`
	DeviationCo2S3RN                   *float64 `json:"deviationCo2S3RN"`
	DeviationCo2Stay                   *float64 `json:"deviationCo2Stay"`
	DeviationWasteRN                   *float64 `json:"deviationWasteRN"`
	DeviationWaterRN                   *float64 `json:"deviationWaterRN"`
	DirectStationary                   *float64 `json:"directStationary"`
	RatingEnergyArea                   *float64 `json:"ratingEnergyArea"`
	RatingEnergyRoom                   *float64 `json:"ratingEnergyRoom"`
	RatingEnergyStay                   *float64 `json:"ratingEnergyStay"`
	RecommendedStars                   *float64 `json:"recommendedStars"`
	DeviationEnergyRN                  *float64 `json:"deviationEnergyRN"`
	DeviationCo2S1Stay                 *float64 `json:"deviationCo2S1Stay"`
	DeviationCo2S2Stay                 *float64 `json:"deviationCo2S2Stay"`
	DeviationCo2S3Stay                 *float64 `json:"deviationCo2S3Stay"`
	DeviationWasteStay                 *float64 `json:"deviationWasteStay"`
	DeviationWaterStay                 *float64 `json:"deviationWaterStay"`
	RecommendedStarsV1                 *float64 `json:"recommendedStarsV1"`
	RecommendedStarsV2                 *float64 `json:"recommendedStarsV2"`
	Scope2Compensation                 *float64 `json:"scope2Compensation"`
	DeviationEnergyStay                *float64 `json:"deviationEnergyStay"`
	EnergyKwhWithCo2Compensation       *float64 `json:"energyKwhWithCo2Compensation"`
	RatingEnergyWithCo2Compensation    *float64 `json:"ratingEnergyWithCo2Compensation"`
	DeviationEnergyWithCo2Compensation *float64 `json:"deviationEnergyWithCo2Compensation"`
}
