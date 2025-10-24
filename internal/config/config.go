package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseHost string
	DatabasePort string
	DatabaseUser string
	DatabasePass string
	DatabaseName string
	ServerPort   string
	ServerHost   string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	config := &Config{
		DatabaseHost: os.Getenv("DATABASE_HOST"),
		DatabasePort: os.Getenv("DATABASE_PORT"),
		DatabaseUser: os.Getenv("DATABASE_USER"),
		DatabasePass: os.Getenv("DATABASE_PASS"),
		DatabaseName: os.Getenv("DATABASE_NAME"),
		ServerPort:   os.Getenv("SERVER_PORT"),
		ServerHost:   os.Getenv("SERVER_HOST"),
	}

	return config
}

func (c *Config) GetConfig() *Config {
	return c
}

func (c *Config) GetDatabaseDSN() string {
	return "host=" + c.DatabaseHost +
		" user=" + c.DatabaseUser +
		" password=" + c.DatabasePass +
		" dbname=" + c.DatabaseName +
		" port=" + c.DatabasePort +
		" sslmode=disable"
}

func (c *Config) GetServerAddress() string {
	return c.ServerHost + ":" + c.ServerPort
}
