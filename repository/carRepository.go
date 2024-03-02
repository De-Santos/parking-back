package repository

import (
	"fmt"
	"gorm.io/gorm/clause"
	"parking-back/initializers"
	"parking-back/models"
	"parking-back/obj"
	"parking-back/utils"
)

func SaveCar(car models.Car) {
	initializers.DB.Create(&car)
}

func GetCarPage(pageable obj.Pageable, parkingId uint) []models.Car {
	var carList []models.Car
	initializers.DB.
		Where(&models.Car{ParkingId: parkingId}).
		Limit(pageable.GetLimit()).
		Offset(utils.GetOffset(pageable.GetPage(), pageable.GetLimit())).
		Find(&carList)
	return carList
}

func UpdateCar(car models.Car) (models.Car, error) {
	tx := initializers.DB.Begin()

	if err := tx.Model(&car).Clauses(clause.Returning{}).Updates(car.GetUpdatedColumns()).Error; err != nil {
		tx.Rollback()
		return models.Car{}, err
	}

	if err := tx.Commit().Error; err != nil {
		return models.Car{}, err
	}

	return car, nil
}

func DeleteCarById(id int) bool {
	result := initializers.DB.Delete(&models.Car{}, id)
	if result.Error != nil {
		fmt.Println("Error:", result.Error)
		return false
	}
	if result.RowsAffected == 0 {
		return false
	}
	return true
}
