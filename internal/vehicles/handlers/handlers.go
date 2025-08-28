package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joshmgreen/Car-Log-API/internal/vehicles/model"
	"github.com/joshmgreen/Car-Log-API/internal/vehicles/service"
)

// APIResponse standardizes responses
type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func parseID(c *gin.Context) (uint, bool) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{Success: false, Error: "Invalid ID"})
		return 0, false
	}
	return uint(id), true
}

func respondError(c *gin.Context, err error, status int) {
	c.JSON(status, APIResponse{Success: false, Error: err.Error()})
}

func GetVehiclesHandler(c *gin.Context) {
	v, err := service.GetVehicles()
	if err != nil {
		respondError(c, err, http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, APIResponse{Success: true, Data: v})
}

func GetVehicleByModelHandler(c *gin.Context) {
	modelName := c.Param("model")
	v, err := service.GetVehicleByModel(modelName)
	if err != nil {
		respondError(c, err, http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, APIResponse{Success: true, Data: v})
}

func AddVehicleHandler(c *gin.Context) {
	var v model.Vehicle
	if err := c.ShouldBindJSON(&v); err != nil {
		respondError(c, err, http.StatusBadRequest)
		return
	}
	if err := service.AddVehicle(v); err != nil {
		respondError(c, err, http.StatusInternalServerError)
		return
	}
	log.Printf("Vehicle added: %+v", v)
	c.JSON(http.StatusCreated, APIResponse{Success: true, Data: v})
}

func UpdateVehicleHandler(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	var v model.Vehicle
	if err := c.ShouldBindJSON(&v); err != nil {
		respondError(c, err, http.StatusBadRequest)
		return
	}
	v.ID = id
	updated, err := service.UpdateVehicle(v)
	if err != nil {
		respondError(c, err, http.StatusInternalServerError)
		return
	}
	if updated {
		c.JSON(http.StatusOK, APIResponse{Success: true, Data: v})
	} else {
		c.JSON(http.StatusNotFound, APIResponse{Success: false, Error: "Vehicle not found"})
	}
}

func DeleteVehicleHandler(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	deleted, err := service.DeleteVehicleByID(id)
	if err != nil {
		respondError(c, err, http.StatusInternalServerError)
		return
	}
	if deleted {
		c.JSON(http.StatusOK, APIResponse{Success: true, Data: "Vehicle deleted"})
	} else {
		c.JSON(http.StatusNotFound, APIResponse{Success: false, Error: "Vehicle not found"})
	}
}
