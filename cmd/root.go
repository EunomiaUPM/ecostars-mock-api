package main

import (
	"ecostars-fake-api/internal/config"
)

func main() {
	config := config.LoadConfig()
	BootstrapHTTPServer(config)
}
