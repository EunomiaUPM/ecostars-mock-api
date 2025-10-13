package http

import (
	"ecostars-fake-api/internal/domain"
	"ecostars-fake-api/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HotelRouter struct {
	HotelService services.HotelService
}

func (r *HotelRouter) Create(router *gin.Engine) {
	hotelGroup := router.Group("/hotels")
	hotelGroup.GET("/", r.getAllHotels)
	hotelGroup.POST("/", r.createHotel)
	hotelGroup.GET("/:id", r.getHotelById)
	hotelGroup.GET("/:id/measures", r.getMeasuresByHotelId)
	measuresGroup := router.Group("/measures")
	measuresGroup.POST("/", r.createMeasure)
	measuresGroup.GET("/:id", r.getMeasureById)
}

func (r *HotelRouter) getAllHotels(c *gin.Context) {
	hotels, err := r.HotelService.GetAllHotels(c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, hotels)
}

func (r *HotelRouter) createHotel(c *gin.Context) {
	var newHotel domain.HotelCreateRequest
	if err := c.ShouldBindJSON(&newHotel); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	hotel, err := r.HotelService.CreateHotel(c.Request.Context(), &newHotel)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, hotel)
}

func (r *HotelRouter) getHotelById(c *gin.Context) {
	id := c.Param("id")
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid hotel ID"})
		return
	}
	hotel, err := r.HotelService.GetHotelByID(c.Request.Context(), uint(uid))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, hotel)
}

func (r *HotelRouter) createMeasure(c *gin.Context) {
	var newMeasure domain.HotelMeasureCreateRequest
	if err := c.ShouldBindJSON(&newMeasure); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	measure, err := r.HotelService.CreateHotelMeasure(c.Request.Context(), &newMeasure)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, measure)
}

func (r *HotelRouter) getMeasuresByHotelId(c *gin.Context) {
	id := c.Param("id")
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid hotel ID"})
		return
	}
	measures, err := r.HotelService.GetHotelMeasuresByHotelID(c.Request.Context(), uint(uid))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, measures)
}

func (r *HotelRouter) getMeasureById(c *gin.Context) {
	id := c.Param("id")
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid measure ID"})
		return
	}
	measure, err := r.HotelService.GetHotelMeasureByID(c.Request.Context(), uint(uid))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, measure)
}
