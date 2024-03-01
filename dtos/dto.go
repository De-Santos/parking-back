package dtos

import "parking-back/models"

type LoginDto struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
}

type SignupDto struct {
	FullName string `validate:"required"`
	Username string `validate:"required"`
	Password string `validate:"required"`
}

type ParkingDto struct {
	ID          uint
	Owner       string `validate:"required"`
	Address     string `validate:"required"`
	Capacity    int    `validate:"required,gt=0"`
	Coordinates models.Coordinates
	CreatedBy   string
}
