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
	offset := utils.GetOffset(pageable.GetPage(), pageable.GetLimit())
	fmt.Println(offset)
	initializers.DB.
		Preload("CreatedBy").
		Limit(pageable.GetLimit()).
		Offset(offset).
		Find(&parkingList)
	return parkingList
}
