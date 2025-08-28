package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joshmgreen/Car-Log-API/internal/db"
	"github.com/joshmgreen/Car-Log-API/internal/vehicles/handlers"
)

func main() {
	db.Init()

	router := gin.Default()

	router.GET("/vehicles", handlers.GetVehiclesHandler)
	router.GET("/vehicles/model/:model", handlers.GetVehicleByModelHandler)
	router.POST("/vehicles", handlers.AddVehicleHandler)
	router.PUT("/vehicles/:id", handlers.UpdateVehicleHandler)
	router.DELETE("/vehicles/:id", handlers.DeleteVehicleHandler)

	router.Run("0.0.0.0:8080")
}
