package http

import (
	"github.com/gin-gonic/gin"
	"github.com/joshmgreen/Car-Log-API/internal/vehicles/handlers"
)

func NewRouter() *gin.Engine {
    router := gin.Default()

    // Vehicle routes
    vehicleGroup := router.Group("/vehicles")
    {
        vehicleGroup.GET("", handlers.GetVehiclesHandler)
        vehicleGroup.GET("/model/:model", handlers.GetVehicleByModelHandler)
        vehicleGroup.POST("", handlers.AddVehicleHandler)
        vehicleGroup.PUT("/:id", handlers.UpdateVehicleHandler)
        vehicleGroup.DELETE("/:id", handlers.DeleteVehicleHandler)
    }

    return router
}
