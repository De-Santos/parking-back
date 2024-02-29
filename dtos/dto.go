package dtos

import "parking-back/models"

type LoginDto struct {
	Username string
	Password string
}

type SignupDto struct {
	FullName string
	Username string
	Password string
}

type ParkingDto struct {
	ID          uint
	Owner       string
	Address     string
	Capacity    int
	Password    string
	Coordinates models.Coordinates
	CreatedBy   string
}
