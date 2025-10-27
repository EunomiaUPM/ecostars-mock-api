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
	HotelID                            uint    `json:"hotel_id" binding:"required" faker:"boundary_start=1, boundary_end=50"`
	Year                               int     `json:"year" binding:"required" faker:"boundary_start=2023, boundary_end=2027"`
	Rooms                              int     `json:"rooms" binding:"required" faker:"boundary_start=10, boundary_end=2000"`
	Stays                              int     `json:"stays" binding:"required" faker:"boundary_start=98316, boundary_end=182588"`
	GasKwh                             float64 `json:"gasKwh" binding:"required" faker:"boundary_start=614693, boundary_end=1141573"`
	WaterM3                            float64 `json:"waterM3" binding:"required" faker:"boundary_start=12950, boundary_end=24050"`
	Occupied                           int     `json:"occupied" binding:"required" faker:"boundary_start=55457, boundary_end=102993"`
	Deviation                          float64 `json:"deviation" binding:"required" faker:"boundary_start=-0.7618, boundary_end=-0.4102"`
	EnergyKwh                          float64 `json:"energyKwh" binding:"required" faker:"boundary_start=740566, boundary_end=1375338"`
	RatingCo2                          float64 `json:"ratingCo2" binding:"required" faker:"boundary_start=-1.732, boundary_end=-0.933"`
	TotalArea                          float64 `json:"totalArea" binding:"required" faker:"boundary_start=9026, boundary_end=16763"`
	StaysPerRN                         float64 `json:"staysPerRN" binding:"required" faker:"boundary_start=1.24, boundary_end=2.30"`
	Electricity                        float64 `json:"electricity" binding:"required" faker:"boundary_start=830.5, boundary_end=1542.5"`
	RatingCo2RN                        float64 `json:"ratingCo2RN" binding:"required" faker:"boundary_start=-1.730, boundary_end=-0.931"`
	RatingWater                        float64 `json:"ratingWater" binding:"required" faker:"boundary_start=0.164, boundary_end=0.306"`
	DeviationCo2                       float64 `json:"deviationCo2" binding:"required" faker:"boundary_start=-1.440, boundary_end=-0.775"`
	RatingEnergy                       float64 `json:"ratingEnergy" binding:"required" faker:"boundary_start=9.36, boundary_end=17.38"`
	RatingCo2S1RN                      float64 `json:"ratingCo2S1RN" binding:"required" faker:"boundary_start=1.415, boundary_end=2.629"`
	RatingCo2S2RN                      float64 `json:"ratingCo2S2RN" binding:"required" faker:"boundary_start=0.010, boundary_end=0.019"`
	RatingCo2Stay                      float64 `json:"ratingCo2Stay" binding:"required" faker:"boundary_start=-0.976, boundary_end=-0.525"`
	RatingWaterRN                      float64 `json:"ratingWaterRN" binding:"required" faker:"boundary_start=0.163, boundary_end=0.303"`
	DeviationCo2RN                     float64 `json:"deviationCo2RN" binding:"required" faker:"boundary_start=-1.440, boundary_end=-0.775"`
	DeviationWaste                     float64 `json:"deviationWaste" binding:"required" faker:"boundary_start=-1.3, boundary_end=-0.7"`
	DeviationWater                     float64 `json:"deviationWater" binding:"required" faker:"boundary_start=-0.552, boundary_end=-0.297"`
	ElectricityKwh                     float64 `json:"electricityKwh" binding:"required" faker:"boundary_start=2934.9, boundary_end=5450.5"`
	RatingEnergyRN                     float64 `json:"ratingEnergyRN" binding:"required" faker:"boundary_start=9.35, boundary_end=17.36"`
	TotalEmissions                     float64 `json:"totalEmissions" binding:"required" faker:"boundary_start=-137069, boundary_end=-73806"`
	Co2Compensation                    float64 `json:"co2Compensation" binding:"required" faker:"boundary_start=186786, boundary_end=346889"`
	DeviationEnergy                    float64 `json:"deviationEnergy" binding:"required" faker:"boundary_start=-0.914, boundary_end=-0.492"`
	RatingCo2S1Stay                    float64 `json:"ratingCo2S1Stay" binding:"required" faker:"boundary_start=0.798, boundary_end=1.483"`
	RatingCo2S2Stay                    float64 `json:"ratingCo2S2Stay" binding:"required" faker:"boundary_start=0.005, boundary_end=0.011"`
	RatingWaterArea                    float64 `json:"ratingWaterArea" binding:"required" faker:"boundary_start=1.00, boundary_end=1.86"`
	RatingWaterRoom                    float64 `json:"ratingWaterRoom" binding:"required" faker:"boundary_start=52.8, boundary_end=98.2"`
	RatingWaterStay                    float64 `json:"ratingWaterStay" binding:"required" faker:"boundary_start=0.092, boundary_end=0.171"`
	Scope1Emissions                    float64 `json:"scope1Emissions" binding:"required" faker:"boundary_start=112149, boundary_end=208278"`
	Scope2Emissions                    float64 `json:"scope2Emissions" binding:"required" faker:"boundary_start=830.5, boundary_end=1542.5"`
	DeviationCo2S1RN                   float64 `json:"deviationCo2S1RN" binding:"required" faker:"boundary_start=-0.875, boundary_end=-0.471"`
	DeviationCo2S2RN                   float64 `json:"deviationCo2S2RN" binding:"required" faker:"boundary_start=-1.297, boundary_end=-0.698"`
	DeviationCo2S3RN                   float64 `json:"deviationCo2S3RN" binding:"required" faker:"boundary_start=-1.3, boundary_end=-0.7"`
	DeviationCo2Stay                   float64 `json:"deviationCo2Stay" binding:"required" faker:"boundary_start=-1.442, boundary_end=-0.776"`
	DeviationWasteRN                   float64 `json:"deviationWasteRN" binding:"required" faker:"boundary_start=-1.3, boundary_end=-0.7"`
	DeviationWaterRN                   float64 `json:"deviationWaterRN" binding:"required" faker:"boundary_start=-0.562, boundary_end=-0.302"`
	DirectStationary                   float64 `json:"directStationary" binding:"required" faker:"boundary_start=112149, boundary_end=208278"`
	RatingEnergyArea                   float64 `json:"ratingEnergyArea" binding:"required" faker:"boundary_start=57.4, boundary_end=106.6"`
	RatingEnergyRoom                   float64 `json:"ratingEnergyRoom" binding:"required" faker:"boundary_start=3022, boundary_end=5614"`
	RatingEnergyStay                   float64 `json:"ratingEnergyStay" binding:"required" faker:"boundary_start=5.27, boundary_end=9.79"`
	RecommendedStars                   float64 `json:"recommendedStars" binding:"required" faker:"boundary_start=3.0, boundary_end=5.0"`
	DeviationEnergyRN                  float64 `json:"deviationEnergyRN" binding:"required" faker:"boundary_start=-0.914, boundary_end=-0.492"`
	DeviationCo2S1Stay                 float64 `json:"deviationCo2S1Stay" binding:"required" faker:"boundary_start=-0.869, boundary_end=-0.468"`
	DeviationCo2S2Stay                 float64 `json:"deviationCo2S2Stay" binding:"required" faker:"boundary_start=-1.297, boundary_end=-0.698"`
	DeviationCo2S3Stay                 float64 `json:"deviationCo2S3Stay" binding:"required" faker:"boundary_start=-1.3, boundary_end=-0.7"`
	DeviationWasteStay                 float64 `json:"deviationWasteStay" binding:"required" faker:"boundary_start=-1.3, boundary_end=-0.7"`
	DeviationWaterStay                 float64 `json:"deviationWaterStay" binding:"required" faker:"boundary_start=-0.539, boundary_end=-0.290"`
	RecommendedStarsV1                 float64 `json:"recommendedStarsV1" binding:"required" faker:"boundary_start=3.0, boundary_end=5.0"`
	RecommendedStarsV2                 float64 `json:"recommendedStarsV2" binding:"required" faker:"boundary_start=3.0, boundary_end=5.0"`
	Scope2Compensation                 float64 `json:"scope2Compensation" binding:"required" faker:"boundary_start=186786, boundary_end=346889"`
	DeviationEnergyStay                float64 `json:"deviationEnergyStay" binding:"required" faker:"boundary_start=-0.908, boundary_end=-0.489"`
	EnergyKwhWithCo2Compensation       float64 `json:"energyKwhWithCo2Compensation" binding:"required" faker:"boundary_start=740126, boundary_end=1374520"`
	RatingEnergyWithCo2Compensation    float64 `json:"ratingEnergyWithCo2Compensation" binding:"required" faker:"boundary_start=9.35, boundary_end=17.37"`
	DeviationEnergyWithCo2Compensation float64 `json:"deviationEnergyWithCo2Compensation" binding:"required" faker:"boundary_start=-0.914, boundary_end=-0.492"`
}

type HotelMeasureUpdateRequest struct {
	Rooms                              *int     `json:"rooms" binding:"required" faker:"boundary_start=10, boundary_end=2000"`
	Stays                              *int     `json:"stays" binding:"required" faker:"boundary_start=98316, boundary_end=182588"`
	GasKwh                             *float64 `json:"gasKwh" binding:"required" faker:"boundary_start=614693, boundary_end=1141573"`
	WaterM3                            *float64 `json:"waterM3" binding:"required" faker:"boundary_start=12950, boundary_end=24050"`
	Occupied                           *int     `json:"occupied" binding:"required" faker:"boundary_start=55457, boundary_end=102993"`
	Deviation                          *float64 `json:"deviation" binding:"required" faker:"boundary_start=-0.7618, boundary_end=-0.4102"`
	EnergyKwh                          *float64 `json:"energyKwh" binding:"required" faker:"boundary_start=740566, boundary_end=1375338"`
	RatingCo2                          *float64 `json:"ratingCo2" binding:"required" faker:"boundary_start=-1.732, boundary_end=-0.933"`
	TotalArea                          *float64 `json:"totalArea" binding:"required" faker:"boundary_start=9026, boundary_end=16763"`
	StaysPerRN                         *float64 `json:"staysPerRN" binding:"required" faker:"boundary_start=1.24, boundary_end=2.30"`
	Electricity                        *float64 `json:"electricity" binding:"required" faker:"boundary_start=830.5, boundary_end=1542.5"`
	RatingCo2RN                        *float64 `json:"ratingCo2RN" binding:"required" faker:"boundary_start=-1.730, boundary_end=-0.931"`
	RatingWater                        *float64 `json:"ratingWater" binding:"required" faker:"boundary_start=0.164, boundary_end=0.306"`
	DeviationCo2                       *float64 `json:"deviationCo2" binding:"required" faker:"boundary_start=-1.440, boundary_end=-0.775"`
	RatingEnergy                       *float64 `json:"ratingEnergy" binding:"required" faker:"boundary_start=9.36, boundary_end=17.38"`
	RatingCo2S1RN                      *float64 `json:"ratingCo2S1RN" binding:"required" faker:"boundary_start=1.415, boundary_end=2.629"`
	RatingCo2S2RN                      *float64 `json:"ratingCo2S2RN" binding:"required" faker:"boundary_start=0.010, boundary_end=0.019"`
	RatingCo2Stay                      *float64 `json:"ratingCo2Stay" binding:"required" faker:"boundary_start=-0.976, boundary_end=-0.525"`
	RatingWaterRN                      *float64 `json:"ratingWaterRN" binding:"required" faker:"boundary_start=0.163, boundary_end=0.303"`
	DeviationCo2RN                     *float64 `json:"deviationCo2RN" binding:"required" faker:"boundary_start=-1.440, boundary_end=-0.775"`
	DeviationWaste                     *float64 `json:"deviationWaste" binding:"required" faker:"boundary_start=-1.3, boundary_end=-0.7"`
	DeviationWater                     *float64 `json:"deviationWater" binding:"required" faker:"boundary_start=-0.552, boundary_end=-0.297"`
	ElectricityKwh                     *float64 `json:"electricityKwh" binding:"required" faker:"boundary_start=2934.9, boundary_end=5450.5"`
	RatingEnergyRN                     *float64 `json:"ratingEnergyRN" binding:"required" faker:"boundary_start=9.35, boundary_end=17.36"`
	TotalEmissions                     *float64 `json:"totalEmissions" binding:"required" faker:"boundary_start=-137069, boundary_end=-73806"`
	Co2Compensation                    *float64 `json:"co2Compensation" binding:"required" faker:"boundary_start=186786, boundary_end=346889"`
	DeviationEnergy                    *float64 `json:"deviationEnergy" binding:"required" faker:"boundary_start=-0.914, boundary_end=-0.492"`
	RatingCo2S1Stay                    *float64 `json:"ratingCo2S1Stay" binding:"required" faker:"boundary_start=0.798, boundary_end=1.483"`
	RatingCo2S2Stay                    *float64 `json:"ratingCo2S2Stay" binding:"required" faker:"boundary_start=0.005, boundary_end=0.011"`
	RatingWaterArea                    *float64 `json:"ratingWaterArea" binding:"required" faker:"boundary_start=1.00, boundary_end=1.86"`
	RatingWaterRoom                    *float64 `json:"ratingWaterRoom" binding:"required" faker:"boundary_start=52.8, boundary_end=98.2"`
	RatingWaterStay                    *float64 `json:"ratingWaterStay" binding:"required" faker:"boundary_start=0.092, boundary_end=0.171"`
	Scope1Emissions                    *float64 `json:"scope1Emissions" binding:"required" faker:"boundary_start=112149, boundary_end=208278"`
	Scope2Emissions                    *float64 `json:"scope2Emissions" binding:"required" faker:"boundary_start=830.5, boundary_end=1542.5"`
	DeviationCo2S1RN                   *float64 `json:"deviationCo2S1RN" binding:"required" faker:"boundary_start=-0.875, boundary_end=-0.471"`
	DeviationCo2S2RN                   *float64 `json:"deviationCo2S2RN" binding:"required" faker:"boundary_start=-1.297, boundary_end=-0.698"`
	DeviationCo2S3RN                   *float64 `json:"deviationCo2S3RN" binding:"required" faker:"boundary_start=-1.3, boundary_end=-0.7"`
	DeviationCo2Stay                   *float64 `json:"deviationCo2Stay" binding:"required" faker:"boundary_start=-1.442, boundary_end=-0.776"`
	DeviationWasteRN                   *float64 `json:"deviationWasteRN" binding:"required" faker:"boundary_start=-1.3, boundary_end=-0.7"`
	DeviationWaterRN                   *float64 `json:"deviationWaterRN" binding:"required" faker:"boundary_start=-0.562, boundary_end=-0.302"`
	DirectStationary                   *float64 `json:"directStationary" binding:"required" faker:"boundary_start=112149, boundary_end=208278"`
	RatingEnergyArea                   *float64 `json:"ratingEnergyArea" binding:"required" faker:"boundary_start=57.4, boundary_end=106.6"`
	RatingEnergyRoom                   *float64 `json:"ratingEnergyRoom" binding:"required" faker:"boundary_start=3022, boundary_end=5614"`
	RatingEnergyStay                   *float64 `json:"ratingEnergyStay" binding:"required" faker:"boundary_start=5.27, boundary_end=9.79"`
	RecommendedStars                   *float64 `json:"recommendedStars" binding:"required" faker:"boundary_start=3.0, boundary_end=5.0"`
	DeviationEnergyRN                  *float64 `json:"deviationEnergyRN" binding:"required" faker:"boundary_start=-0.914, boundary_end=-0.492"`
	DeviationCo2S1Stay                 *float64 `json:"deviationCo2S1Stay" binding:"required" faker:"boundary_start=-0.869, boundary_end=-0.468"`
	DeviationCo2S2Stay                 *float64 `json:"deviationCo2S2Stay" binding:"required" faker:"boundary_start=-1.297, boundary_end=-0.698"`
	DeviationCo2S3Stay                 *float64 `json:"deviationCo2S3Stay" binding:"required" faker:"boundary_start=-1.3, boundary_end=-0.7"`
	DeviationWasteStay                 *float64 `json:"deviationWasteStay" binding:"required" faker:"boundary_start=-1.3, boundary_end=-0.7"`
	DeviationWaterStay                 *float64 `json:"deviationWaterStay" binding:"required" faker:"boundary_start=-0.539, boundary_end=-0.290"`
	RecommendedStarsV1                 *float64 `json:"recommendedStarsV1" binding:"required" faker:"boundary_start=3.0, boundary_end=5.0"`
	RecommendedStarsV2                 *float64 `json:"recommendedStarsV2" binding:"required" faker:"boundary_start=3.0, boundary_end=5.0"`
	Scope2Compensation                 *float64 `json:"scope2Compensation" binding:"required" faker:"boundary_start=186786, boundary_end=346889"`
	DeviationEnergyStay                *float64 `json:"deviationEnergyStay" binding:"required" faker:"boundary_start=-0.908, boundary_end=-0.489"`
	EnergyKwhWithCo2Compensation       *float64 `json:"energyKwhWithCo2Compensation" binding:"required" faker:"boundary_start=740126, boundary_end=1374520"`
	RatingEnergyWithCo2Compensation    *float64 `json:"ratingEnergyWithCo2Compensation" binding:"required" faker:"boundary_start=9.35, boundary_end=17.37"`
	DeviationEnergyWithCo2Compensation *float64 `json:"deviationEnergyWithCo2Compensation" binding:"required" faker:"boundary_start=-0.914, boundary_end=-0.492"`
}
