package obj

import (
	"parking-back/models"
	"time"
)

type LoginDto struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SignupDto struct {
	FullName string `json:"full_name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type ParkingDto struct {
	ID          uint               `json:"id"`
	ParkingName string             `json:"parking_name" validate:"required"`
	Owner       string             `json:"owner" validate:"required"`
	Address     string             `json:"address" validate:"required"`
	Capacity    int                `json:"capacity" validate:"required,gt=0"`
	Coordinates models.Coordinates `json:"coordinates"`
	CreatedBy   string             `json:"created_by"`
	CreatedAt   time.Time          `json:"created_at"`
}

type CarDto struct {
	ID         uint      `json:"id"`
	Vrp        string    `json:"vrp" validate:"required"`
	Arrived    time.Time `json:"arrived" validate:"required"`
	Expiration time.Time `json:"expiration" validate:"required"`
	ParkingId  uint      `json:"parking_id" validate:"gte=0"`
}
