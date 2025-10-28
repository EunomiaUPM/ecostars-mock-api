package main

import (
	"ecostars-fake-api/internal/config"
	"ecostars-fake-api/internal/controllers/http"
	"ecostars-fake-api/internal/repo"
	"ecostars-fake-api/internal/services"

	"ecostars-fake-api/internal/database"

	"github.com/gin-gonic/gin"
)

func BootstrapHTTPServer(config *config.Config) {
	db, err := database.ConnectDB(config)
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

	router.Run(config.GetStaticServerAddress()) // listens on 0.0.0.0:8080 by default
}
