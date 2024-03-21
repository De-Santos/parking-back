package repository

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"parking-back/gorm_scope"
	"parking-back/initializers"
	"parking-back/models"
	"parking-back/obj"
)

func SaveCar(car models.Car) {
	initializers.DB.Create(&car)
}

func GetCarPage(pagination obj.Pagination, query obj.Search, parkingId uint) []models.Car {
	var carList []models.Car
	whereFunc := gorm_scope.FlexWhere(carList, query)
	initializers.DB.
		Scopes(gorm_scope.Paginate(countFunction(parkingId, whereFunc), pagination, initializers.DB)).
		Scopes(whereFunc).
		Where(&models.Car{ParkingId: parkingId}).
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

func countFunction(parkingId uint, where func(db *gorm.DB) *gorm.DB) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.
			Model(models.Car{}).
			Where(&models.Car{ParkingId: parkingId}).
			Scopes(where)
	}
}
