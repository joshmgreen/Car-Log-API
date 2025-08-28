package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joshmgreen/Car-Log-API/db"
	"github.com/joshmgreen/Car-Log-API/handlers"
)

func main() {
	// Initialize database and run migrations
	db.Init()

	//Router setup
	router := gin.Default()

	// Register handlers
	router.GET("/vehicles", handlers.HandleGetVehicles)
	router.GET("/vehicles/model/:model", handlers.HandleGetVehicleByModel)
	router.POST("/vehicles", handlers.HandleAddVehicle)
	router.PUT("/vehicles/:id", handlers.HandleUpdateVehicle)
	router.DELETE("/vehicles/:id", handlers.HandleDeleteVehicleById)

	router.Run("0.0.0.0:8080")
}
