package vehicles

import (
	"strings"

	"github.com/joshmgreen/Car-Log-API/db"
	"github.com/joshmgreen/Car-Log-API/models"
)

func GetVehicles() ([]models.Vehicle, error) {
	var vehicles []models.Vehicle
	result := db.DB.Find(&vehicles)
	return vehicles, result.Error
}

func AddVehicle(v models.Vehicle) error {
	result := db.DB.Create(&v)
	return result.Error
}

func DeleteVehicleByID(id int) (bool, error) {
	result := db.DB.Delete(&models.Vehicle{}, id)
	if result.RowsAffected == 0 {
		return false, nil
	}
	return true, result.Error
}

// GetVehicleByModel returns a pointer to a vehicle matching the model (case-insensitive)
func GetVehicleByModel(model string) (*models.Vehicle, error) {
	var vehicle models.Vehicle
	result := db.DB.Where("LOWER(model) = ?", strings.ToLower(model)).First(&vehicle)
	if result.Error != nil {
		return nil, result.Error
	}
	return &vehicle, nil

}

func UpdateVehicle(updated models.Vehicle) (bool, error) {
	result := db.DB.Model(&models.Vehicle{}).
		Where("id = ?", updated.ID).
		Updates(updated)

	if result.RowsAffected == 0 {
		return false, nil
	}
	return true, result.Error
}
