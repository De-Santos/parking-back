package mapper

import (
	"gorm.io/gorm"
	"parking-back/models"
	"parking-back/obj"
)

func MapToCarModel(dto obj.CarDto) models.Car {
	return models.Car{
		Model:      gorm.Model{ID: dto.ID},
		Vpr:        dto.Vpr,
		Arrived:    dto.Arrived,
		Expiration: dto.Expiration,
		ParkingId:  dto.ParkingId,
	}
}

func MapToCarDtoList(parkingList []models.Car) []obj.CarDto {
	parkingDtoList := make([]obj.CarDto, len(parkingList))
	for i, parking := range parkingList {
		parkingDtoList[i] = MapToCarDto(parking)
	}
	return parkingDtoList
}

func MapToCarDto(car models.Car) obj.CarDto {
	return obj.CarDto{
		ID:         car.ID,
		Vpr:        car.Vpr,
		Arrived:    car.Arrived,
		Expiration: car.Expiration,
		ParkingId:  car.ParkingId,
	}
}
