package main

import (
	"context"
	"ecostars-fake-api/internal/config"
	"ecostars-fake-api/internal/controllers/http"
	"ecostars-fake-api/internal/database"
	"ecostars-fake-api/internal/domain"
	"ecostars-fake-api/internal/repo"
	"ecostars-fake-api/internal/services"
	"sync"

	"github.com/gin-gonic/gin"
)

func BootstrapMetricsCollector(config *config.Config) {
	db, err := database.ConnectDB(config)
	if err != nil {
		panic("failed to connect database")
	}
	_ = db

	// Initialize Metric Simulator Service
	var wg sync.WaitGroup
	ctx, _ := context.WithCancel(context.Background())
	notificationService := services.NewNotificationService(
		&repo.SubscriptionRepository{DB: db},
		10,
	)
	go notificationService.Start(ctx, &wg)
	metricSimulatorService := services.NewMetricSimulatorService(
		&repo.MetricItemRepository{DB: db},
		notificationService,
	)

	// Register some metric items
	metricSimulatorService.RegisterMetricItem(&domain.MetricItemCreateRequest{
		ItemType: domain.MetricKeyHotelEnergyUsage,
	})
	metricSimulatorService.RegisterMetricItem(&domain.MetricItemCreateRequest{
		ItemType: domain.MetricKeyHotelWasteGenerated,
	})
	metricSimulatorService.RegisterMetricItem(&domain.MetricItemCreateRequest{
		ItemType: domain.MetricKeyHotelWaterUsage,
	})
	// Start the metric simulator
	go metricSimulatorService.Start(ctx)

	// Initialize Notification Service
	router := gin.Default()
	subscriptionsRouter := http.SubscriptionsRouter{
		SubscriptionService: services.NewSubscriptionService(
			&repo.SubscriptionRepository{DB: db},
		),
	}
	subscriptionsRouter.Create(router)

	router.Run(config.GetDynamicServerAddress()) // listens on
}
