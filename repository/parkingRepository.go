package repository

import (
	"fmt"
	"parking-back/initializers"
	"parking-back/models"
	"parking-back/obj"
	"parking-back/utils"
)

func GetParkingPage(pageable obj.Pageable) []models.Parking {
	var parkingList []models.Parking
	initializers.DB.
		Preload("CreatedBy").
		Limit(pageable.GetLimit()).
		Offset(utils.GetOffset(pageable.GetPage(), pageable.GetLimit())).
		Find(&parkingList)
	return parkingList
}

func UpdateParking(parking models.Parking) (models.Parking, error) {
	tx := initializers.DB.Begin()

	if err := tx.Model(&parking).Updates(parking.GetUpdatedColumns()).Error; err != nil {
		tx.Rollback()
		return models.Parking{}, err
	}

	var updatedParking models.Parking
	if err := tx.Preload("CreatedBy").First(&updatedParking, parking.ID).Error; err != nil {
		tx.Rollback()
		return models.Parking{}, err
	}

	if err := tx.Commit().Error; err != nil {
		return models.Parking{}, err
	}

	return updatedParking, nil
}

func DeleteParkingById(id int) bool {
	result := initializers.DB.Delete(&models.Parking{}, id)
	if result.Error != nil {
		fmt.Println("Error:", result.Error)
		return false
	}
	if result.RowsAffected == 0 {
		return false
	}
	return true
}

func CreateParking(parking models.Parking) (models.Parking, error) {
	tx := initializers.DB.Begin()

	tx.Create(&parking)

	if err := tx.Preload("CreatedBy").First(&parking, parking.ID).Error; err != nil {
		tx.Rollback()
		return models.Parking{}, err
	}

	if err := tx.Commit().Error; err != nil {
		return models.Parking{}, err
	}

	return parking, nil
}
