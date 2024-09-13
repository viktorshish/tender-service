package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"tender-service/internal/config"
	"tender-service/internal/routes"
)

func main() {
	config.ConnectDB()

	router := gin.Default()

	routes.SetupRoutes(router)

	log.Fatal(router.Run(":8080"))
}
