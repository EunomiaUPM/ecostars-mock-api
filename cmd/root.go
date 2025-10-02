package main

import (
	"ecostars-fake-api/internal/database"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		panic("failed to connect database")
	}
	_ = db

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run("0.0.0.0:8081") // listens on 0.0.0.0:8080 by default
}
