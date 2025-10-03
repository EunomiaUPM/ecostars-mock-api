package http

import (
	"ecostars-fake-api/internal/services"

	"github.com/gin-gonic/gin"
)

type HotelRouter struct {
	HotelService services.HotelService
}

func (r *HotelRouter) Create(router *gin.Engine) {
	hotelGroup := router.Group("/hotels")
	hotelGroup.GET("/", r.getAllHotels)
}

func (r *HotelRouter) getAllHotels(c *gin.Context) {
	hotels, err := r.HotelService.GetAllHotels(c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, hotels)
}
