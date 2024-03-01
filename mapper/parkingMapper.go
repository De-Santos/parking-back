package mapper

import (
	"parking-back/models"
	"parking-back/obj"
)

func MapToParkingModel(dto obj.ParkingDto, userId uint) models.Parking {
	return models.Parking{
		Owner:       dto.Owner,
		Address:     dto.Address,
		Capacity:    dto.Capacity,
		Coordinates: dto.Coordinates,
		CreatedByID: userId,
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
		Owner:       parking.Owner,
		Address:     parking.Address,
		Capacity:    parking.Capacity,
		Coordinates: parking.Coordinates,
		CreatedBy:   parking.CreatedBy.FullName,
		CreatedAt:   parking.CreatedAt,
	}
}
