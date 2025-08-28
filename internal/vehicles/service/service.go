package service

import (
	"strings"

	"github.com/joshmgreen/Car-Log-API/internal/db"
	"github.com/joshmgreen/Car-Log-API/internal/vehicles/model"
)

// GetVehicles returns all vehicles
func GetVehicles() ([]model.Vehicle, error) {
	var vehicles []model.Vehicle
	result := db.DB.Find(&vehicles)
	return vehicles, result.Error
}

// AddVehicle adds a new vehicle
func AddVehicle(v model.Vehicle) error {
	result := db.DB.Create(&v)
	return result.Error
}

// DeleteVehicleByID deletes a vehicle by ID
func DeleteVehicleByID(id uint) (bool, error) {
	result := db.DB.Delete(&model.Vehicle{}, id)
	return result.RowsAffected > 0, result.Error
}

// GetVehicleByModel returns a vehicle matching the model (case-insensitive)
func GetVehicleByModel(modelName string) ([]model.Vehicle, error) {
	var vehicles []model.Vehicle
	result := db.DB.
		Where("LOWER(model) LIKE ?", "%" + strings.ToLower(modelName)+"%").
		Find(&vehicles)
	if result.Error != nil {
		return nil, result.Error
	}
	return vehicles, nil
}

// UpdateVehicle updates an existing vehicle
func UpdateVehicle(updated model.Vehicle) (bool, error) {
	result := db.DB.Model(&model.Vehicle{}).
		Where("id = ?", updated.ID).
		Updates(updated)
	return result.RowsAffected > 0, result.Error
}
