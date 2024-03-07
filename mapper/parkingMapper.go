package mapper

import (
	"gorm.io/gorm"
	"parking-back/models"
	"parking-back/obj"
)

func MapToParkingModelWithUser(dto obj.ParkingDto, userId uint) models.Parking {
	return models.Parking{
		ParkingName: dto.ParkingName,
		Owner:       dto.Owner,
		Address:     dto.Address,
		Capacity:    dto.Capacity,
		Coordinates: dto.Coordinates,
		CreatedByID: userId,
	}
}

func MapToParkingModel(dto obj.ParkingDto) models.Parking {
	return models.Parking{
		Model:       gorm.Model{ID: dto.ID},
		ParkingName: dto.ParkingName,
		Owner:       dto.Owner,
		Address:     dto.Address,
		Capacity:    dto.Capacity,
		Coordinates: dto.Coordinates,
	}
}

func MapToParkingDtoList(parkingList []models.Parking) []obj.ParkingDto {
	parkingDtoList := make([]obj.ParkingDto, len(parkingList))
	for i, parking := range parkingList {
		parkingDtoList[i] = MapToParkingDto(parking)
	}
	return parkingDtoList
}

func MapToParkingDto(parking models.Parking) obj.ParkingDto {
	return obj.ParkingDto{
		ID:          parking.ID,
		ParkingName: parking.ParkingName,
		Owner:       parking.Owner,
		Address:     parking.Address,
		Capacity:    parking.Capacity,
		Coordinates: parking.Coordinates,
		CreatedBy:   parking.CreatedBy.FullName,
		CreatedAt:   parking.CreatedAt,
	}
}
