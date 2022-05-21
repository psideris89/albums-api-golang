package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	apiRouting "psideris/albums-api/routes"
)

func init() {
	fmt.Println("Initialising API")
}

func main() {
	router := gin.Default()

	apiRouting.ConfigureRoutes(router)

	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatalln("Failed to start the API on port 8080")
	}
}
