package mapper

import (
	"parking-back/dtos"
	"parking-back/models"
)

func MapToParkingModel(dto dtos.ParkingDto, userId uint) models.Parking {
	return models.Parking{
		Owner:       dto.Owner,
		Address:     dto.Address,
		Capacity:    dto.Capacity,
		Coordinates: dto.Coordinates,
		CreatedByID: userId,
	}
}
