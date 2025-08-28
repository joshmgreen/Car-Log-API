package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joshmgreen/Car-Log-API/models"
	"github.com/joshmgreen/Car-Log-API/vehicles"
)

// Standardized API response
type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// -------- Helper functions --------
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

// -------- Handlers --------

// GET /vehicles
func HandleGetVehicles(c *gin.Context) {
	v, err := vehicles.GetVehicles()
	if err != nil {
		respondError(c, err, http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, APIResponse{Success: true, Data: v})
}

// GET /vehicles/model/:model
func HandleGetVehicleByModel(c *gin.Context) {
	modelParam := c.Param("model")
	v, err := vehicles.GetVehicleByModel(modelParam)
	if err != nil {
		respondError(c, err, http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, APIResponse{Success: true, Data: v})
}

// POST /vehicles
func HandleAddVehicle(c *gin.Context) {
	var newVehicle models.Vehicle
	if err := c.ShouldBindJSON(&newVehicle); err != nil {
		respondError(c, err, http.StatusBadRequest)
		return
	}

	if err := vehicles.AddVehicle(newVehicle); err != nil {
		respondError(c, err, http.StatusInternalServerError)
		return
	}

	log.Printf("Vehicle added: %+v", newVehicle)
	c.JSON(http.StatusCreated, APIResponse{Success: true, Data: newVehicle})
}

// PUT /vehicles/:id
func HandleUpdateVehicle(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}

	var vehicle models.Vehicle
	if err := c.ShouldBindJSON(&vehicle); err != nil {
		respondError(c, err, http.StatusBadRequest)
		return
	}
	vehicle.ID = id

	updated, err := vehicles.UpdateVehicle(vehicle)
	if err != nil {
		respondError(c, err, http.StatusInternalServerError)
		return
	}

	if updated {
		c.JSON(http.StatusOK, APIResponse{Success: true, Data: vehicle})
	} else {
		c.JSON(http.StatusNotFound, APIResponse{Success: false, Error: "Vehicle not found"})
	}
}

// DELETE /vehicles/:id
func HandleDeleteVehicleById(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}

	deleted, err := vehicles.DeleteVehicleByID(int(id))
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
