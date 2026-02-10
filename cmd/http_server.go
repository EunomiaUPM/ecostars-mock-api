package main

import (
	"ecostars-fake-api/internal/config"
	"ecostars-fake-api/internal/controllers/http"
	"ecostars-fake-api/internal/middleware"
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

	// Public routes (Documentation)
	router.StaticFile("/openapi.yaml", "./specs/openapi/hotels.yaml")
	router.GET("/docs", func(c *gin.Context) {
		c.Header("Content-Type", "text/html")
		c.String(200, `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="utf-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1" />
  <title>Ecostars API Documentation</title>
  <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@5/swagger-ui.css" />
</head>
<body>
  <div id="swagger-ui"></div>
  <script src="https://unpkg.com/swagger-ui-dist@5/swagger-ui-bundle.js" charset="UTF-8"></script>
  <script>
    window.onload = () => {
      window.ui = SwaggerUIBundle({
        url: "/openapi.yaml",
        dom_id: '#swagger-ui',
      });
    };
  </script>
</body>
</html>`)
	})

	// Apply Auth Middleware
	router.Use(middleware.AuthMiddleware(config))

	hotelRouter := http.HotelRouter{
		HotelService: services.HotelService{
			HotelRepository:        &repo.HotelRepository{DB: db},
			HotelMeasureRepository: &repo.HotelMeasureRepository{DB: db},
		},
	}
	hotelRouter.Create(router)

	router.Run(config.GetStaticServerAddress()) // listens on 0.0.0.0:8080 by default
}
