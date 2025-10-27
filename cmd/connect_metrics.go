package main

import (
	"ecostars-fake-api/internal/config"
	"fmt"
)

func BootstrapMetricsCollector(config *config.Config) {
	fmt.Println("Metrics collector started with config:", config)
}
