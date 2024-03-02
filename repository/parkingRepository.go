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
