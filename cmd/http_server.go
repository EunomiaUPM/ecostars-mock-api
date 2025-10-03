package main

import (
	"ecostars-fake-api/internal/controllers/http"
	"ecostars-fake-api/internal/repo"
	"ecostars-fake-api/internal/services"

	"ecostars-fake-api/internal/database"

	"github.com/gin-gonic/gin"
)

func BootstrapHTTPServer() {
	db, err := database.ConnectDB()
	if err != nil {
		panic("failed to connect database")
	}
	_ = db

	router := gin.Default()
	hotelRouter := http.HotelRouter{
		HotelService: services.HotelService{
			HotelRepository:        &repo.HotelRepository{DB: db},
			HotelMeasureRepository: &repo.HotelMeasureRepository{DB: db},
		},
	}
	hotelRouter.Create(router)

	router.Run("0.0.0.0:8081") // listens on 0.0.0.0:8080 by default
}
